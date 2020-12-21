package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func printStructInfo(x interface{}) {
	// typeOf wrapped value inside the Interface{}
	t := reflect.TypeOf(x)
	fmt.Printf("-----Kind - %v - ----\n", t.Kind())
	if t.Kind() != reflect.Struct {
		fmt.Printf("ERR: Not a struct value, expected struct value of 'kind' struct\n")
		return
	}

	n := t.NumField()
	fmt.Printf("struct of type %v has %v fields\n", t, n)

	for i := 0; i < n; i++ {
		tt := t.Field(i)
		fmt.Printf("Field index: %v,\nname: %v,\ntype: %v\n", i, tt.Name, tt.Type)
	}
	fmt.Println("-------")
}

type Rib struct {
	VersionStr          string        `json:"version_str"`
	NodeIDStr           string        `json:"node_id_str"`
	EncodingPath        string        `json:"encoding_path"`
	CollectionID        int           `json:"collection_id"`
	CollectionStartTime string        `json:"collection_start_time"`
	CollectionEndTime   string        `json:"collection_end_time"`
	MsgTimestamp        string        `json:"msg_timestamp"`
	SubscriptionID      string        `json:"subscription_id"`
	SensorGroupID       []interface{} `json:"sensor_group_id"`
	DataSource          string        `json:"data_source"`
	Data                []struct {
		VrfName        string `json:"vrfName"`
		Address        string `json:"address"`
		MaskLen        int    `json:"maskLen"`
		L3NextHopCount int    `json:"l3NextHopCount"`
		EventType      string `json:"eventType"`
		NextHop        []struct {
			Address      string `json:"address"`
			OutInterface string `json:"outInterface"`
			VrfName      string `json:"vrfName"`
			Owner        string `json:"owner"`
			Preference   int    `json:"preference"`
			Metric       int    `json:"metric"`
			Tag          int    `json:"tag"`
			SegmentID    int    `json:"segmentId"`
			TunnelID     int    `json:"tunnelId"`
			EncapType    string `json:"encapType"`
			NhTypeFlags  int    `json:"nhTypeFlags"`
		} `json:"nextHop"`
	} `json:"data"`
}

type RibGeneric struct {
	VersionStr          string        `json:"version_str"`
	NodeIDStr           string        `json:"node_id_str"`
	EncodingPath        string        `json:"encoding_path"`
	CollectionID        int           `json:"collection_id"`
	CollectionStartTime string        `json:"collection_start_time"`
	CollectionEndTime   string        `json:"collection_end_time"`
	MsgTimestamp        string        `json:"msg_timestamp"`
	SubscriptionID      string        `json:"subscription_id"`
	SensorGroupID       []interface{} `json:"sensor_group_id"`
	DataSource          string        `json:"data_source"`
	VrfName             string        `json:"data.vrfName"`
	Address             string        `json:"data.address"`
	MaskLen             int           `json:"data.maskLen"`
	L3NextHopCount      int           `json:"data.l3NextHopCount"`
	EventType           string        `json:"data.eventType"`
	NextHopAddress      string        `json:"data.nexthop.address"`
	OutInterface        string        `json:"data.nexthop.outInterface"`
	NextHopVrfName      string        `json:"data.nexthop.vrfName"`
	Owner               string        `json:"data.nexthop.owner"`
	Preference          int           `json:"data.nexthop.preference"`
	Metric              int           `json:"data.nexthop.metric"`
	Tag                 int           `json:"data.nexthop.tag"`
	SegmentID           int           `json:"data.nexthop.segmentId"`
	TunnelID            int           `json:"data.nexthop.tunnelId"`
	EncapType           string        `json:"data.nexthop.encapType"`
	NhTypeFlags         int           `json:"data.nexthop.nhTypeFlags"`
}

func ribhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Is not POST method")
		return
	} else {
		data, _ := ioutil.ReadAll(r.Body)

		var response Rib
		err := json.Unmarshal(data, &response)

		if err != nil {
			log.Print(err)
		}
		printStructInfo(response)
		for _, v := range response.Data {
			printStructInfo(v)
			for _, v := range v.NextHop {
				printStructInfo(v)
			}
		}
	}
}

func main() {

	http.HandleFunc("/network/rib", ribhandler)
	http.ListenAndServe(":10000", nil)
}
