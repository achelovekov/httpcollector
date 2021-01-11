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
				  "autoRemapReplicationServers": false,
				  "cfgSrc": "unknown",
				  "childAction": "",
				  "controllerId": 0,
				  "delayRestoreMsBrdrGwExpiryTime": "NA",
				  "descr": "",
				  "encapType": "unknown",
				  "epId": 1,
				  "holdDownTime": 500,
				  "holdDownTimerExpiryTime": "NA",
				  "holdUpTime": 30,
				  "holdUpTimerExpiryTime": "NA",
				  "hostReach": "bgp",
				  "ingressReplProtoBGP": false,
				  "learningMode": "CP",
				  "mac": "00:3A:9C:5B:6A:87",
				  "mcastGroupL2": "0.0.0.0",
				  "mcastGroupL3": "0.0.0.0",
				  "modTs": "2021-01-11T12:12:29.611+00:00",
				  "multisiteBordergwInterface": "unspecified",
				  "multisiteBrdrGwIntfIp": "0.0.0.0",
				  "operEncapType": "vxlan",
				  "operStAnycastSrcIntf": "down",
				  "operStMultisiteBrdrGwLoopbackIntf": "down",
				  "operStSrcLoopbackIntf": "up",
				  "operState": "up",
				  "primaryIp": "172.17.1.2",
				  "primaryIpv6": "",
				  "rn": "epId-1",
				  "secondaryIp": "172.17.1.101",
				  "secondaryIpv6": "",
				  "sourceInterface": "lo1",
				  "status": "",
				  "suppressARP": false,
				  "virtualMac": "00:00:00:00:00:00",
				  "virtualRtrMac": "02:00:AC:11:01:65",
				  "virtualRtrMacReorig": "00:00:00:00:00:00",
				  "vpcVIPNotified": true
				},
				"children": [
				  {
					"nvoPeers": {
					  "attributes": {
						"childAction": "",
						"modTs": "2021-01-11T11:52:27.839+00:00",
						"rn": "peers",
						"status": ""
					  },
					  "children": [
						{
						  "nvoDyPeer": {
							"attributes": {
							  "createTs": "2021-01-11T19:09:01.979+00:00",
							  "firstVNI": 2012007,
							  "ip": "172.17.1.203",
							  "mac": "BC:4A:56:CD:EC:3F",
							  "rn": "dy_peer-[172.17.1.203]",
							  "state": "Up",
							  "upStateTransitionTs": "2021-01-11T19:09:02.012+00:00"
							}
						  }
						},
						{
						  "nvoDyPeer": {
							"attributes": {
							  "createTs": "2021-01-11T19:09:01.978+00:00",
							  "firstVNI": 2012007,
							  "ip": "172.17.1.202",
							  "mac": "D4:C9:3C:85:62:7F",
							  "rn": "dy_peer-[172.17.1.202]",
							  "state": "Up",
							  "upStateTransitionTs": "2021-01-11T19:09:02.009+00:00"
							}
						  }
						},
						{
						  "nvoDyPeer": {
							"attributes": {
							  "createTs": "2021-01-11T19:09:01.976+00:00",
							  "firstVNI": 2012007,
							  "ip": "172.17.1.201",
							  "mac": "D4:C9:3C:85:62:3F",
							  "rn": "dy_peer-[172.17.1.201]",
							  "state": "Up",
							  "upStateTransitionTs": "2021-01-11T19:09:02.008+00:00"
							}
						  }
						}
					  ]
					}
				  }
				]
			  }
			}
		  ]
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
