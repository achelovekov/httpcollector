package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	for k1, v := range src {
		if reflect.ValueOf(v).Type().Kind() != reflect.Slice && reflect.ValueOf(v).Type().Kind() != reflect.Map {
			if pathIndex == 0 {
				header[k1] = reflect.ValueOf(v).Interface()
			} else {
				header[path[pathIndex-1]+"."+k1] = reflect.ValueOf(v).Interface()
			}
		} else if reflect.ValueOf(v).Type().Kind() == reflect.Map {
			for k2, v := range src[k1].(map[string]interface{}) {
				if pathIndex == 0 {
					header[k1+"."+k2] = reflect.ValueOf(v).Interface()
				} else {
					header[path[pathIndex-1]+"."+k1+"."+k2] = reflect.ValueOf(v).Interface()
				}
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

func ribhandler(w http.ResponseWriter, r *http.Request) {
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

		path := []string{"data", "nextHop"}

		Flatten(src, path, pathIndex, header)
	}
}

func macAllHandler(w http.ResponseWriter, r *http.Request) {
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

		var pathIndex int

		header := make(map[string]interface{})

		path := []string{"data", "list"}

		Flatten(src, path, pathIndex, header)
	}
}

func adjacencyHandler(w http.ResponseWriter, r *http.Request) {
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

		var pathIndex int

		header := make(map[string]interface{})

		path := []string{"data"}

		Flatten(src, path, pathIndex, header)

	}
}

func main() {
	http.HandleFunc("/network/rib", ribhandler)
	http.HandleFunc("/network/mac-all", macAllHandler)
	http.HandleFunc("/network/adjacency", adjacencyHandler)
	http.ListenAndServe(":10000", nil)
}
