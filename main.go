package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FlattenStruct(src map[string]interface{}, dst map[string]interface{}) {
	for k, v := range src {
		fmt.Printf("%v %v\n", k, v)
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
		json.Unmarshal(data, &src)

		FlattenStruct(src, dst)
	}
}

func main() {

	http.HandleFunc("/network/rib", ribhandler)
	http.ListenAndServe(":10000", nil)
}
