package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func PrettyPrint(src map[string]interface{}) {
	empJSON, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Pretty processed output %s\n", string(empJSON))
}

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

func worker(r *http.Request, path []string) {
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

		Flatten(src, path, pathIndex, header)
	}
}

func ribhandler(w http.ResponseWriter, r *http.Request) {
	path := []string{"data", "nextHop"}
	worker(r, path)
}

func macAllHandler(w http.ResponseWriter, r *http.Request) {
	path := []string{"data", "list"}
	worker(r, path)
}

func adjacencyHandler(w http.ResponseWriter, r *http.Request) {
	path := []string{"data"}
	worker(r, path)
}

func vxlanHandler(w http.ResponseWriter, r *http.Request) {
	path := []string{"data", "imdata"}
	worker(r, path)
}

func main() {
	http.HandleFunc("/network/rib", ribhandler)
	http.HandleFunc("/network/mac-all", macAllHandler)
	http.HandleFunc("/network/adjacency", adjacencyHandler)
	http.HandleFunc("/network/EVENT-LIST", vxlanHandler)
	http.ListenAndServe(":10000", nil)
}
