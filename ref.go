package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func PrettyPrint(src map[string]interface{}) {
	empJSON, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent function output %s\n", string(empJSON))
}

func combineHeaders(src map[string]interface{}, pathIndex int, path []string) (map[string]interface{}, int) {

	header := make(map[string]interface{})

	for k, v := range src {
		if reflect.ValueOf(v).Type().Kind() != reflect.Slice {
			if pathIndex == 0 {
				header[k] = reflect.ValueOf(v).Interface()
			} else {
				header[path[pathIndex-1]+"."+k] = reflect.ValueOf(v).Interface()
			}
		}
	}

	return header, pathIndex
}

func Flatten(src map[string]interface{}, path []string, pathIndex int, header map[string]interface{}) {

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
			Flatten(v.(map[string]interface{}), path, pathIndex+1, newHeader)
		}
	}
}

func main() {
	var Raw = `
	{
		"a1": "value-a1",
		"b1": "value-b1",
		"c1": "value-c1",
		"d1": [
		  {
			"a2": "value-a2-01",
			"b2": "value-b2-02",
			"c2": [
			  {
				"a3": "value-a3-01",
				"b3": "value-b3-02",
				"c3": "value-c3-03"
			  },
			  {
				"a3": "value-a3-04",
				"b3": "value-b3-05",
				"c3": "value-c3-06"
			  }
			]
		  },
		  {
			"a2": "value-a2-03",
			"b2": "value-b2-04",
			"c2": []
		  }
		],
		"f1": [
		  {
			"af2": "value-af2-01",
			"bf2": "value-bf2-02",
			"cf2": "value-cf2-03"
		  },
		  {
			"af2": "value-af2-01",
			"bf2": "value-bf2-02",
			"cf2": "value-cf2-03"
		  }
		]
	  }
`

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(Raw), &jsonMap)
	if err != nil {
		panic(err)
	}

	path := []string{"d1", "c2"}

	var pathIndex int

	header := make(map[string]interface{})

	Flatten(jsonMap, path, pathIndex, header)

}
