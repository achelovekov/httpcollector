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

func flatten(esClient *es.Client, src map[string]interface{}, path []string, pathIndex int, header map[string]interface{}) {
	newHeader := make(map[string]interface{})
	for k, v := range header {
		newHeader[k] = v
	}

	src = src[path[pathIndex]].(map[string]interface{})
	if v, ok := src["attributes"]; ok {
		for k, v := range v.(map[string]interface{}) {
			newHeader[path[pathIndex]+"."+k] = v
		}
	}
	if v, ok := src["children"]; ok && pathIndex != len(path)-1 {
		for i := 0; i < reflect.ValueOf(v).Len(); i++ {
			v := reflect.ValueOf(v).Index(i).Interface().(map[string]interface{})
			if _, ok := v[path[pathIndex+1]]; ok {
				flatten(esClient, v, path, pathIndex+1, newHeader)
			}
		}
	} else {
		PrettyPrint(newHeader)
		esPush(esClient, "golang-index", newHeader)
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

		newHeader := make(map[string]interface{})

		for k, v := range src {
			if k != "data" {
				newHeader[k] = v
			}
		}

		flatten(esClient, src["data"].(map[string]interface{}), path, pathIndex, newHeader)
	}
}

type postReqHandler struct {
	esClient *es.Client
}

func (prh *postReqHandler) vxlanSysEpsHandler(w http.ResponseWriter, r *http.Request) {
	var path = []string{"nvoEps", "nvoEp", "nvoNws", "nvoNw"}
	worker(prh.esClient, r, path)
}

func main() {
	esClient, error := esConnect("10.62.186.54", "9200")

	if error != nil {
		log.Fatalf("error: %s", error)
	}

	postReqHandler := &postReqHandler{esClient: esClient}
	http.HandleFunc("/network/vxlan:sys/eps", postReqHandler.vxlanSysEpsHandler)
	//http.HandleFunc("/network/vxlan:sys/bd", postReqHandler.vxlanSysBdHandler)

	http.ListenAndServe(":11000", nil)

}