package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func flatten(src map[string]interface{}, path []string, pathIndex int, header map[string]interface{}) {
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
				flatten(v, path, pathIndex+1, newHeader)
			}
		}
	} else {
		fmt.Println(newHeader)
	}

}

func main() {

	var Raw = `{
		"k1": {
		  "attributes": {
			"a11": "v11",
			"a12": "v12"
		  },
		  "children": [
			{
				"k2": {
					"attributes": {
						"a21": "v21",
						"a22": "v22"
					},
					"children": [
						{
							"k3": {
								"attributes": {
									"a31": "v31",
									"a32": "v32"
								}, 
								"children": [
									{
										"k4": {
											"attributes": {
												"a41": "v41",
												"a42": "v42"
											}	
										}
									}
								]		
							}
						},
						{
							"k3": {
								"attributes": {
									"a31": "v33",
									"a32": "v34"
								}, 
								"children": [
									{
										"k4": {
											"attributes": {
												"a41": "v43",
												"a42": "v44"
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

	src := make(map[string]interface{})
	err := json.Unmarshal([]byte(Raw), &src)
	if err != nil {
		panic(err)
	}

	var path = []string{"k1", "k2", "k3", "k4"}
	var pathIndex int
	header := make(map[string]interface{})

	flatten(src, path, pathIndex, header)
}
