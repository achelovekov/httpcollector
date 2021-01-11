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

//flattens the field if that field is map
func flattenMap(src map[string]interface{}, preHeader map[string]interface{}, path []string, pathIndex int, prefix string) map[string]interface{} {

	for k, v := range src {

		if reflect.ValueOf(v).Type().Kind() != reflect.Map {
			preHeader[strings.Join(path[:pathIndex], ".")+"."+prefix+"."+k] = v
		} else if reflect.ValueOf(v).Type().Kind() == reflect.Map {
			prefix = prefix + "." + k
			flattenMap(src[k].(map[string]interface{}), preHeader, path, pathIndex, prefix)
		}

	}
	return preHeader
}

//combines headers from previous recursive calls
func combineHeaders(src map[string]interface{}, pathIndex int, path []string) (map[string]interface{}, int) {

	header := make(map[string]interface{})
	mapsKeys := make(map[int]string)

	var i int

	for k, v := range src {
		if reflect.ValueOf(v).Type().Kind() != reflect.Slice && reflect.ValueOf(v).Type().Kind() != reflect.Map {
			if pathIndex == 0 {
				header[k] = reflect.ValueOf(v).Interface()
			} else {
				header[strings.Join(path[:pathIndex], ".")+"."+k] = reflect.ValueOf(v).Interface()
			}
		} else if reflect.ValueOf(v).Type().Kind() == reflect.Map {
			mapsKeys[i] = k
			i = i + 1
		}
	}

	if len(mapsKeys) > 0 {
		for _, v := range mapsKeys {
			preHeader := make(map[string]interface{})
			for k, v := range header {
				preHeader[k] = v
			}
			prefix := v
			header = flattenMap(reflect.ValueOf(src[v]).Interface().(map[string]interface{}), preHeader, path, pathIndex, prefix)
		}

	}

	return header, pathIndex
}

//general flatten func
func Flatten(esClient *es.Client, src map[string]interface{}, path []string, pathIndex int, header map[string]interface{}) {

	addHeader, pathIndex := combineHeaders(src, pathIndex, path)
	newHeader := make(map[string]interface{})

	for k, v := range addHeader {
		newHeader[k] = v
	}

	for k, v := range header {
		newHeader[k] = v
	}

	if pathIndex == len(path) {
		PrettyPrint(newHeader)
		esPush(esClient, "golang-index", newHeader)
	} else if reflect.ValueOf(src[path[pathIndex]]).Len() == 0 {
		newHeader[path[pathIndex]] = make([]interface{}, 0)
		PrettyPrint(newHeader)
		esPush(esClient, "golang-index", newHeader)
	} else {
		for i := 0; i < reflect.ValueOf(src[path[pathIndex]]).Len(); i++ {
			v := reflect.ValueOf(src[path[pathIndex]]).Index(i).Interface()
			Flatten(esClient, v.(map[string]interface{}), path, pathIndex+1, newHeader)
		}
	}
}

//general post call serving func
func worker(esClient *es.Client, r *http.Request, path []string) {
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

		if err != nil {
			log.Fatalf(err.Error())
		}

		var pathIndex int

		header := make(map[string]interface{})

		Flatten(esClient, src, path, pathIndex, header)
	}
}

type postReqHandler struct {
	esClient *es.Client
}

func (prh *postReqHandler) ribhandler(w http.ResponseWriter, r *http.Request) {
	path := []string{"data", "nextHop"}
	worker(prh.esClient, r, path)
}

func (prh *postReqHandler) macAllHandler(w http.ResponseWriter, r *http.Request) {
	path := []string{"data", "list"}
	worker(prh.esClient, r, path)
}

func (prh *postReqHandler) adjacencyHandler(w http.ResponseWriter, r *http.Request) {
	path := []string{"data"}
	worker(prh.esClient, r, path)
}

func (prh *postReqHandler) vxlanHandler(w http.ResponseWriter, r *http.Request) {
	path := []string{"data", "imdata"}
	worker(prh.esClient, r, path)
}

func (prh *postReqHandler) vxlanSysEpsHandler(w http.ResponseWriter, r *http.Request) {
	path := []string{"children", "children", "children"}
	worker(prh.esClient, r, path)
}

func (prh *postReqHandler) vxlanSysBdHandler(w http.ResponseWriter, r *http.Request) {
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
}

func main() {
	esClient, error := esConnect("10.62.186.54", "9200")

	if error != nil {
		log.Fatalf("error: %s", error)
	}

	postReqHandler := &postReqHandler{esClient: esClient}

	http.HandleFunc("/network/rib", postReqHandler.ribhandler)
	http.HandleFunc("/network/mac-all", postReqHandler.macAllHandler)
	http.HandleFunc("/network/adjacency", postReqHandler.adjacencyHandler)
	http.HandleFunc("/network/EVENT-LIST", postReqHandler.vxlanHandler)
	http.HandleFunc("/network/vxlan:sys/eps", postReqHandler.vxlanSysEpsHandler)
	http.HandleFunc("/network/vxlan:sys/bd", postReqHandler.vxlanSysBdHandler)

	http.ListenAndServe(":10000", nil)
}
