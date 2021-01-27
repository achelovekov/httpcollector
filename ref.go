package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	fmt.Printf("MarshalIndent function output %s\n", string(empJSON))
}

func flattenMap(src map[string]interface{}, preHeader map[string]interface{}, path []string, pathIndex int, prefix string) map[string]interface{} {

	for k, v := range src {
		fmt.Printf("path: %v, prefix in flatten: %v\n", strings.Join(path[:pathIndex], "."), prefix)

		if reflect.ValueOf(v).Type().Kind() != reflect.Map {
			preHeader[strings.Join(path[:pathIndex], ".")+"."+prefix+"."+k] = v
			fmt.Println(preHeader)
		} else if reflect.ValueOf(v).Type().Kind() == reflect.Map {
			prefix = prefix + "." + k
			flattenMap(src[k].(map[string]interface{}), preHeader, path, pathIndex, prefix)
		}

	}
	fmt.Printf("Preheader: %v\n", preHeader)
	return preHeader
}

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
			//fmt.Printf("prefix before flatten: %v\n", prefix)
			header = flattenMap(reflect.ValueOf(src[v]).Interface().(map[string]interface{}), preHeader, path, pathIndex, prefix)
		}

	}

	return header, pathIndex
}

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
	} else if reflect.ValueOf(src[path[pathIndex]]).Len() == 0 {
		newHeader[path[pathIndex]] = make([]interface{}, 0)
		PrettyPrint(newHeader)
	} else {
		for i := 0; i < reflect.ValueOf(src[path[pathIndex]]).Len(); i++ {
			v := reflect.ValueOf(src[path[pathIndex]]).Index(i).Interface()
			Flatten(esClient, v.(map[string]interface{}), path, pathIndex+1, newHeader)
		}
	}
}

func main() {
	es, error := esConnect("10.62.186.54", "9200")

	if error != nil {
		log.Fatalf("error: %s", error)
	}

	var Raw = `
	{
		"nvoEps": {
		  "attributes": {
			"childAction": "",
			"dn": "sys/eps",
			"modTs": "2021-01-11T11:52:27.087+00:00",
			"status": ""
		  },
		  "children": [
			{
			  "nvoEp": {
				"attributes": {
				  "adminSt": "enabled",
				  "adminStMultisiteBrdrGwLoopackIntf": "disabled",
				  "advertiseVmac": false,
				  "anycastIntf": "unspecified",
				}
			  }
			}
		}
	}
	`

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(Raw), &jsonMap)
	if err != nil {
		panic(err)
	}

	path := []string{"children", "children", "children"}

	var pathIndex int

	header := make(map[string]interface{})

	Flatten(es, jsonMap, path, pathIndex, header)

}
