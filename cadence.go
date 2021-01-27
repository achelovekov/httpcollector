package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch"
	es "github.com/elastic/go-elasticsearch"
	esapi "github.com/elastic/go-elasticsearch/esapi"
)

func enrich(src map[string]interface{}, enrichmentMap map[string]map[string]int, enrichKeys []string) {
	for _, key := range enrichKeys {
		if v, ok := src[key]; ok {
			if reflect.ValueOf(v).Type().Kind() == reflect.Int {
				v = strconv.Itoa(v.(int))
			}
			if reflect.ValueOf(v).Type().Kind() == reflect.Float64 {
				v = strconv.FormatFloat(v.(float64), 'E', -1, 64)
			}
			src[key+"/code"] = enrichmentMap[key][v.(string)]
		}
	}
}

func filter(src map[string]interface{}, filterList []string) {
	for _, key := range filterList {
		if _, ok := src[key]; ok {
			//fmt.Printf("SUCCESSFULY DELETED - %v\n", key)
			delete(src, key)
		}
	}
}

func toInt(v interface{}) interface{} {
	if reflect.ValueOf(v).Type().Kind() == reflect.String {
		vInt, err := strconv.Atoi(v.(string))
		if err != nil {
			return v
		}
		return vInt
	}
	return v
}

//only direct paths supported. Full path to end must be specified (no children).
func flattenMap(esClient *es.Client, src map[string]interface{}, path [][]string, pathIndex int, header map[string]interface{}, enrichmentMap map[string]map[string]int, enrichKeys []string, filterList []string) {
	newHeader := make(map[string]interface{})
	for k, v := range header {
		newHeader[k] = toInt(v)
	}

	for index := range path[pathIndex] {
		if _, ok := src[path[pathIndex][index]]; ok {
			v := reflect.ValueOf(src[path[pathIndex][index]]).Interface().(map[string]interface{})
			if v, ok := v["attributes"]; ok {
				for k, v := range v.(map[string]interface{}) {
					newHeader[path[pathIndex][index]+"."+k] = toInt(v)
				}
			}

			if v, ok := v["children"]; ok && pathIndex != len(path)-1 {
				for i := 0; i < reflect.ValueOf(v).Len(); i++ {
					v := reflect.ValueOf(v).Index(i).Interface().(map[string]interface{})
					for index := range path[pathIndex+1] {
						if _, ok := v[path[pathIndex+1][index]]; ok {
							flattenMap(esClient, v, path, pathIndex+1, newHeader, enrichmentMap, enrichKeys, filterList)
						}
					}
				}
			} else {
				enrich(newHeader, enrichmentMap, enrichKeys)
				filter(newHeader, filterList)
				//PrettyPrint(newHeader)
				esPush(esClient, "golang-index", newHeader)
			}
		}
	}
}

func flattenList(esClient *es.Client, src map[string]interface{}, path [][]string, pathIndex int, header map[string]interface{}, enrichmentMap map[string]map[string]int, enrichKeys []string, filterList []string) {
	newHeader := make(map[string]interface{})
	for k, v := range header {
		newHeader[k] = v
	}

	for i := 0; i < reflect.ValueOf(src["data"]).Len(); i++ {
		v := reflect.ValueOf(src["data"]).Index(i).Interface()
		flattenMap(esClient, v.(map[string]interface{}), path, pathIndex, newHeader, enrichmentMap, enrichKeys, filterList)
	}

}

func esConnect(ipaddr string, port string) (*es.Client, error) {

	var fulladdress string = "http://" + ipaddr + ":" + port

	cfg := elasticsearch.Config{
		Addresses: []string{
			fulladdress,
		},
	}

	es, _ := elasticsearch.NewClient(cfg)

	return es, nil
}

func esPush(esClient *es.Client, indexName string, body map[string]interface{}) {
	empJSON, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	req := esapi.IndexRequest{
		Index: indexName,                          // Index name
		Body:  strings.NewReader(string(empJSON)), // Document body
	}

	res, err := req.Do(context.Background(), esClient)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)
}

func PrettyPrint(src map[string]interface{}) {
	empJSON, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Pretty processed output %s\n", string(empJSON))
}

func worker(esClient *es.Client, r *http.Request, path [][]string, enrichmentMap map[string]map[string]int, enrichKeys []string, filterList []string) {
	if r.Method != "POST" {
		fmt.Println("Is not POST method")
		return
	} else {
		data, _ := ioutil.ReadAll(r.Body)

		src := make(map[string]interface{})
		err := json.Unmarshal(data, &src)
		if err != nil {
			panic(err)
		}

		/* 		srcJSON, err := json.MarshalIndent(src, "", "  ")
		   		if err != nil {
		   			log.Fatalf(err.Error())
		   		}
		   		fmt.Printf("MarshalIndent function output %s\n", string(srcJSON))

		   		if err != nil {
		   			log.Fatalf(err.Error())
		   		} */

		var pathIndex int

		newHeader := make(map[string]interface{})

		for k, v := range src {
			if k != "data" {
				newHeader[k] = v
			}
		}

		if reflect.ValueOf(src["data"]).Type().Kind() == reflect.Map {
			flattenMap(esClient, src["data"].(map[string]interface{}), path, pathIndex, newHeader, enrichmentMap, enrichKeys, filterList)
		} else {
			flattenList(esClient, src, path, pathIndex, newHeader, enrichmentMap, enrichKeys, filterList)
		}
	}
}

type postReqHandler struct {
	esClient      *es.Client
	enrichmentMap map[string]map[string]int
	filterList    []string
}

func (prh *postReqHandler) sysEpsHandler(w http.ResponseWriter, r *http.Request) {
	var path = [][]string{{"nvoEps"}, {"nvoEp"}, {"nvoPeers", "nvoNws"}, {"nvoDyPeer", "nvoNw"}}
	var enrichKeys = []string{"nvoEp.operState", "nvoDyPeer.state"}
	var filterList = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys, filterList)
}

func (prh *postReqHandler) sysBdHandler(w http.ResponseWriter, r *http.Request) {
	var path = [][]string{{"l2VlanStats"}}
	var enrichKeys = []string{}
	var filterList = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys, filterList)
}

func (prh *postReqHandler) sysIntfHandler(w http.ResponseWriter, r *http.Request) {
	var path = [][]string{{"l1PhysIf"}, {"rmonIfIn"}}
	var enrichKeys = []string{"l1PhysIf.adminSt"}
	var filterList = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys, filterList)
}

func (prh *postReqHandler) sysChHandler(w http.ResponseWriter, r *http.Request) {
	var path = [][]string{{"eqptLCSlot", "eqptSupCSlot"}, {"eqptLC", "eqptSupC"}, {"eqptSpromLc", "eqptCPU"}, {"eqptSpCmnBlk"}}
	var enrichKeys = []string{"eqptSupC.operSt", "eqptLC.operSt"}
	var filterList = []string{"eqptSpromLc.modTs"}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys, filterList)
}

func (prh *postReqHandler) sysProcHandler(w http.ResponseWriter, r *http.Request) {
	var path = [][]string{{"procEntity"}, {"procEntry"}}
	var enrichKeys = []string{"procEntry.operState"}
	var filterList = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys, filterList)
}

func (prh *postReqHandler) sysProcSysHandler(w http.ResponseWriter, r *http.Request) {
	var path = [][]string{{"procSysLoad", "procSysMemUsed", "procSysCpuSummary"}}
	var enrichKeys = []string{}
	var filterList = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys, filterList)
}

func (prh *postReqHandler) customSysBgp(w http.ResponseWriter, r *http.Request) {
	var path = [][]string{{"bgpEntity"}, {"bgpInst"}, {"bgpDom"}, {"bgpPeer"}, {"bgpPeerEntry"}, {"bgpPeerEntryStats", "bgpPeerAfEntry"}}
	var enrichKeys = []string{"bgpPeerEntry.operSt"}
	var filterList = []string{"bgpPeerEntry.modTs"}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys, filterList)
}

func (prh *postReqHandler) customSysOspf(w http.ResponseWriter, r *http.Request) {
	var path = [][]string{{"ospfEntity"}, {"ospfInst"}, {"ospfDom"}, {"ospfArea"}, {"ospfAreaStats"}}
	var enrichKeys = []string{}
	var filterList = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys, filterList)
}

func enrichmentMapCreate() map[string]map[string]int {
	var EnrichmentMap = map[string]map[string]int{}

	EnrichmentMap["bgpPeerEntry.operSt"] = map[string]int{}
	EnrichmentMap["bgpPeerEntry.operSt"]["unspecified"] = 0
	EnrichmentMap["bgpPeerEntry.operSt"]["illegal"] = 1
	EnrichmentMap["bgpPeerEntry.operSt"]["shut"] = 2
	EnrichmentMap["bgpPeerEntry.operSt"]["idle"] = 3
	EnrichmentMap["bgpPeerEntry.operSt"]["connect"] = 4
	EnrichmentMap["bgpPeerEntry.operSt"]["active"] = 5
	EnrichmentMap["bgpPeerEntry.operSt"]["open-sent"] = 6
	EnrichmentMap["bgpPeerEntry.operSt"]["open-confirm"] = 7
	EnrichmentMap["bgpPeerEntry.operSt"]["established"] = 8
	EnrichmentMap["bgpPeerEntry.operSt"]["closing"] = 9
	EnrichmentMap["bgpPeerEntry.operSt"]["error"] = 10
	EnrichmentMap["bgpPeerEntry.operSt"]["unknown"] = 11

	EnrichmentMap["nvoEp.operState"] = map[string]int{}
	EnrichmentMap["nvoEp.operState"]["up"] = 1
	EnrichmentMap["nvoEp.operState"]["down"] = 0

	EnrichmentMap["nvoDyPeer.state"] = map[string]int{}
	EnrichmentMap["nvoDyPeer.state"]["Up"] = 1
	EnrichmentMap["nvoDyPeer.state"]["Down"] = 0

	EnrichmentMap["l1PhysIf.adminSt"] = map[string]int{}
	EnrichmentMap["l1PhysIf.adminSt"]["up"] = 1
	EnrichmentMap["l1PhysIf.adminSt"]["down"] = 0

	EnrichmentMap["eqptSupC.operSt"] = map[string]int{}
	EnrichmentMap["eqptSupC.operSt"]["online"] = 1
	EnrichmentMap["eqptSupC.operSt"]["offline"] = 0

	EnrichmentMap["eqptLC.operSt"] = map[string]int{}
	EnrichmentMap["eqptLC.operSt"]["online"] = 1

	EnrichmentMap["procEntry.operState"] = map[string]int{}
	EnrichmentMap["procEntry.operState"]["up"] = 1
	EnrichmentMap["procEntry.operState"]["down"] = 0

	return EnrichmentMap
}

func main() {
	esClient, error := esConnect("10.62.186.54", "9200")

	if error != nil {
		log.Fatalf("error: %s", error)
	}

	var enrichmentMap map[string]map[string]int = enrichmentMapCreate()

	postReqHandler := &postReqHandler{esClient: esClient, enrichmentMap: enrichmentMap}
	http.HandleFunc("/network/vxlan:sys/eps", postReqHandler.sysEpsHandler)
	http.HandleFunc("/network/vxlan:sys/bd", postReqHandler.sysBdHandler)
	http.HandleFunc("/network/interface:sys/intf", postReqHandler.sysIntfHandler)
	http.HandleFunc("/network/environment:sys/ch", postReqHandler.sysChHandler)
	http.HandleFunc("/network/resources:sys/proc", postReqHandler.sysProcHandler)
	http.HandleFunc("/network/resources:sys/procsys", postReqHandler.sysProcSysHandler)
	http.HandleFunc("/network/sys/bgp", postReqHandler.customSysBgp)
	http.HandleFunc("/network/sys/ospf", postReqHandler.customSysOspf)

	http.ListenAndServe(":11000", nil)

}
