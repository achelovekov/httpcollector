package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func FlattenMap(src map[string]interface{}, dst map[string]interface{}, sliceKey string, prefix string) ([]map[string]interface{}, map[string]interface{}, string) {
	if reflect.TypeOf(src).Kind() == reflect.Map {
		for k, v := range src {
			if reflect.TypeOf(v).Kind() != reflect.Slice {
				dst[prefix+k] = v
			} else if reflect.TypeOf(v).Kind() == reflect.Slice && k == sliceKey {
				prefix := k + "."
				s := reflect.ValueOf(v).Interface()
				return s.([]map[string]interface{}), dst, prefix
			}
		}
	}

	return nil, nil, ""
}

func FlattenSlice(src []map[string]interface{}, dst map[string]interface{}, sliceKey string, prefix string) {
	for _, v := range src {
		FlattenMap(v, dst, sliceKey, prefix)
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

		a, b, c := FlattenMap(src, dst, "data", prefix)
		FlattenSlice(a, b, "nextHop", c)
	}
}

func main() {
	http.HandleFunc("/network/rib", ribhandler)
	http.ListenAndServe(":10000", nil)
}
