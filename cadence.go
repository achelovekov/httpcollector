package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/elastic/go-elasticsearch"
	es "github.com/elastic/go-elasticsearch"
	esapi "github.com/elastic/go-elasticsearch/esapi"
)

func enrich(src map[string]interface{}, enrichmentMap map[string]map[string]int, enrichKeys []string) {
	for _, key := range enrichKeys {
		if v, ok := src[key]; ok {
			src[key+".code"] = enrichmentMap[key][v.(string)]
		}
	}

	//fmt.Println(enrichmentMap)
	//fmt.Println(enrichKeys)
}

//only direct paths supported
func flattenMap(esClient *es.Client, src map[string]interface{}, path []string, pathIndex int, header map[string]interface{}, enrichmentMap map[string]map[string]int, enrichKeys []string) {
	newHeader := make(map[string]interface{})
	for k, v := range header {
		newHeader[k] = v
	}

	if _, ok := src[path[pathIndex]]; ok {
		v := reflect.ValueOf(src[path[pathIndex]]).Interface().(map[string]interface{})
		if v, ok := v["attributes"]; ok {
			for k, v := range v.(map[string]interface{}) {
				newHeader[path[pathIndex]+"/"+k] = v
			}
		}

		if v, ok := v["children"]; ok && pathIndex != len(path)-1 {
			for i := 0; i < reflect.ValueOf(v).Len(); i++ {
				v := reflect.ValueOf(v).Index(i).Interface().(map[string]interface{})
				if _, ok := v[path[pathIndex+1]]; ok {
					flattenMap(esClient, v, path, pathIndex+1, newHeader, enrichmentMap, enrichKeys)
				}
			}
		} else {
			enrich(newHeader, enrichmentMap, enrichKeys)
			PrettyPrint(newHeader)
			esPush(esClient, "golang-index", newHeader)
		}
	}
}

func flattenList(esClient *es.Client, src map[string]interface{}, path []string, pathIndex int, header map[string]interface{}, enrichmentMap map[string]map[string]int, enrichKeys []string) {
	newHeader := make(map[string]interface{})
	for k, v := range header {
		newHeader[k] = v
	}

	for i := 0; i < reflect.ValueOf(src["data"]).Len(); i++ {
		v := reflect.ValueOf(src["data"]).Index(i).Interface()
		flattenMap(esClient, v.(map[string]interface{}), path, pathIndex, newHeader, enrichmentMap, enrichKeys)
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

func worker(esClient *es.Client, r *http.Request, path []string, enrichmentMap map[string]map[string]int, enrichKeys []string) {
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
		/*
			srcJSON, err := json.MarshalIndent(src, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Printf("MarshalIndent function output %s\n", string(srcJSON))
		*/
		if err != nil {
			log.Fatalf(err.Error())
		}

		var pathIndex int

		newHeader := make(map[string]interface{})

		for k, v := range src {
			if k != "data" {
				newHeader[k] = v
			}
		}

		if reflect.ValueOf(src["data"]).Type().Kind() == reflect.Map {
			flattenMap(esClient, src["data"].(map[string]interface{}), path, pathIndex, newHeader, enrichmentMap, enrichKeys)
		} else {
			flattenList(esClient, src, path, pathIndex, newHeader, enrichmentMap, enrichKeys)
		}
	}
}

type postReqHandler struct {
	esClient      *es.Client
	enrichmentMap map[string]map[string]int
}

func (prh *postReqHandler) vxlanSysEpsHandler(w http.ResponseWriter, r *http.Request) {
	var path = []string{"nvoEps", "nvoEp", "nvoNws", "nvoNw"}
	var enrichKeys = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys)
}

func (prh *postReqHandler) vxlanSysBdHandler(w http.ResponseWriter, r *http.Request) {
	var path = []string{"l2VlanStats"}
	var enrichKeys = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys)
}

func (prh *postReqHandler) vxlanSysIntfHandler(w http.ResponseWriter, r *http.Request) {
	var path = []string{"l1PhysIf", "rmonIfIn"}
	var enrichKeys = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys)
}

func (prh *postReqHandler) vxlanSysChHandler(w http.ResponseWriter, r *http.Request) {
	var path = []string{"eqptSupCSlot", "eqptSupC", "eqptCPU"}
	var enrichKeys = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys)
}

func (prh *postReqHandler) vxlanSysProcHandler(w http.ResponseWriter, r *http.Request) {
	var path = []string{"procEntity", "procEntry"}
	var enrichKeys = []string{}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys)
}

func (prh *postReqHandler) customSysBgp(w http.ResponseWriter, r *http.Request) {
	var path = []string{"bgpEntity", "bgpInst", "bgpDom", "bgpPeer", "bgpPeerEntry", "bgpPeerEntryStats"}
	var enrichKeys = []string{"bgpPeerEntry.operSt"}
	worker(prh.esClient, r, path, prh.enrichmentMap, enrichKeys)
	/*
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

			srcJSON, err := json.MarshalIndent(src, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Printf("MarshalIndent function output %s\n", string(srcJSON))
		}
	*/

}

func enrichmentMapCreate() map[string]map[string]int {
	var EnrichmentMap = map[string]map[string]int{}

	EnrichmentMap["bgpPeerEntry.operSt"] = map[string]int{}
	EnrichmentMap["bgpPeerEntry.operSt"]["established"] = 1

	return EnrichmentMap
}

func main() {
	esClient, error := esConnect("10.62.186.54", "9200")

	if error != nil {
		log.Fatalf("error: %s", error)
	}

	var enrichmentMap map[string]map[string]int = enrichmentMapCreate()

	postReqHandler := &postReqHandler{esClient: esClient, enrichmentMap: enrichmentMap}
	http.HandleFunc("/network/vxlan:sys/eps", postReqHandler.vxlanSysEpsHandler)
	http.HandleFunc("/network/vxlan:sys/bd", postReqHandler.vxlanSysBdHandler)
	http.HandleFunc("/network/interface:sys/intf", postReqHandler.vxlanSysIntfHandler)
	http.HandleFunc("/network/environment:sys/ch", postReqHandler.vxlanSysChHandler)
	http.HandleFunc("/network/resources:sys/proc", postReqHandler.vxlanSysProcHandler)
	http.HandleFunc("/network/sys/bgp", postReqHandler.customSysBgp)

	http.ListenAndServe(":11000", nil)

}
