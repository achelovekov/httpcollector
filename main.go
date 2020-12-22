package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func FlattenStruct(src map[string]interface{}, dst map[string]interface{}, prefix string) {
	if reflect.TypeOf(src).Kind() == reflect.Map {
		for _, v := range src {
			fmt.Printf("type kind: %v", reflect.TypeOf(v).Kind())
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
		dst := make(map[string]interface{})
		prefix := ""
		json.Unmarshal(data, &src)

		FlattenStruct(src, dst, prefix)
	}
}

func main() {
	http.HandleFunc("/network/rib", ribhandler)
	http.ListenAndServe(":10000", nil)
}
