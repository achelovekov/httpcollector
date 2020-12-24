package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func mapToSlice(src map[string]interface{}, key string) []interface{} {
	sli := make([]interface{}, len(src))

	var i int

	for k, v := range src {
		if k != key {
			vType := reflect.ValueOf(v).Type().Kind()
			switch vType {
			case reflect.String:
				val := make([]interface{}, 2)
				val[0] = k
				val[1] = v.(string)
				sli[i] = val
				i = i + 1
			case reflect.Int64:
				val := make([]interface{}, 2)
				val[0] = k
				val[1] = v.(int)
				sli[i] = val
				i = i + 1
			case reflect.Float64:
				val := make([]interface{}, 2)
				val[0] = k
				val[1] = v.(float64)
				sli[i] = val
				i = i + 1
			}

		}
	}

	if reflect.ValueOf(src[key]).Len() == 0 {
		sli[len(src)-1] = []interface{}{key, "null"}
	} else {
		sli[len(src)-1] = []interface{}{key, src[key]}
	}
	nilIndexes := make([]int, 0)

	for i, v := range sli {
		if v == nil {
			nilIndexes = append(nilIndexes, i)
		}
	}

	var j int

	for _, v := range nilIndexes {
		sli = append(sli[:v-j], sli[v-j+1:]...)
		j = j + 1
	}

	return sli

}

func Flatten(src interface{}, dst map[string]interface{}, keysSlice []string, path string) {

	jsonMap := src.(map[string]interface{})

	if len(path) > 0 {
		path = path + "."
	}
	if len(keysSlice) > 0 {

		jsonSlice := mapToSlice(jsonMap, keysSlice[0])

		for _, v := range jsonSlice {
			key := reflect.ValueOf(v).Index(0).Interface().(string)
			value := reflect.ValueOf(v).Index(1).Interface()

			if reflect.ValueOf(value).Type().Kind() != reflect.Slice {
				dst[path+key] = value
			} else if reflect.ValueOf(key).Interface().(string) == keysSlice[0] {

				path = path + key

				keysSliceCopy := make([]string, len(keysSlice))
				copy(keysSliceCopy, keysSlice)

				keysSliceCopy = keysSliceCopy[1:]

				for i := 0; i < reflect.ValueOf(value).Len(); i++ {
					res := reflect.ValueOf(jsonMap[key]).Index(i).Interface()
					Flatten(res, dst, keysSliceCopy, path)
				}

			}
		}
	} else {
		for k, v := range jsonMap {
			if reflect.ValueOf(v).Type().Kind() != reflect.Slice {
				dst[path+k] = v
			}
		}
	}
	if len(keysSlice) == 0 {
		emp, err := json.MarshalIndent(dst, "", "  ")
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("Flattened output: %s\n", string(emp))
	}
}

func main() {
	var jsonStr = `
	{
		"name":"John",
		"age":30,
		"cars": [{"car1":"Ford","car2":"BMW","carts": [{"model1": "9.5hp","color":"green"},{"model1": "20hp", "color": "red"}]},
				{"car1":"Mazda","car2":"Porsche","carts": []}],
		"bikes": [
			{"bike1":"Honda","bike2":"Suza","bike3":"Pina"},
			{"bike1":"Bona","bike2":"Izha","bike3":"Turbo"}
			],
		"cities": []
	}
`

	jsonMap := make(map[string]interface{})
	dst := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}

	keysSlice := []string{"cars", "carts"}

	Flatten(jsonMap, dst, keysSlice, "")

}
