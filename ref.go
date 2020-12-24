package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func combineHeaders(src map[string]interface{}, pathIndex int) (map[string]interface{}, int) {

	header := make(map[string]interface{})

	for k, v := range src {
		if reflect.ValueOf(v).Type().Kind() != reflect.Slice {
			header[k] = reflect.ValueOf(v).Interface()
		}
	}

	return header, pathIndex
}

func Flatten(src map[string]interface{}, path []string, pathIndex int, header map[string]interface{}) {

	addHeader, pathIndex := combineHeaders(src, pathIndex)
	newHeader := make(map[string]interface{})

	for k, v := range addHeader {
		newHeader[k] = v
	}

	for k, v := range header {
		newHeader[k] = v
	}

	if pathIndex == len(path) {
		fmt.Printf("%v\n", newHeader)
	} else if reflect.ValueOf(src[path[pathIndex]]).Len() == 0 {
		newHeader[path[pathIndex]] = make([]interface{}, 0)
		fmt.Printf("%v\n", newHeader)
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

	path := []string{"f1"}

	var pathIndex int

	header := make(map[string]interface{})

	Flatten(jsonMap, path, pathIndex, header)

}
