package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FlattenStruct(src interface{}, dst interface{}, prefix string) {
	for k, v := range src {
		fmt.Printf("%v %v", k, v)
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
		err := json.Unmarshal(data, &src)

		var newRib RibGeneric

		FlattenStruct(src, dst)
	}
}

func main() {

	http.HandleFunc("/network/rib", ribhandler)
	http.ListenAndServe(":10000", nil)
}
