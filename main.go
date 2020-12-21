package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func FlattenStruct(src interface{}, dst interface{}, baseIndex int) {
	// typeOf wrapped value inside the Interface{}
	// r is an object with proper methods to each of supertype

	tSrc := reflect.TypeOf(src)
	tDst := reflect.TypeOf(dst)

	fmt.Printf("tSrc kind is: %v\n", tSrc.Kind())
	fmt.Printf("tDst kind is: %v\n", tDst.Kind())

	vSrc := reflect.ValueOf(src).Elem()
	vDst := reflect.ValueOf(dst).Elem()

	//fmt.Printf("vSrc value is: %v\n", vSrc)
	//fmt.Printf("vDst value is: %v\n", vDst)

	if tSrc.Kind() != reflect.Slice {

		nSrc := vSrc.NumField()
		nDst := vDst.NumField()

		fmt.Printf("tSrc has %v fields\n", nSrc)
		fmt.Printf("tDst has %v fields\n", nDst)

		for i := 0; i < nSrc-1; i++ {
			v := vSrc.Field(i + baseIndex).Interface()
			switch v.(type) {
			case string:
				vDst.Field(i + baseIndex).SetString(v.(string))
			case int64:
				vDst.Field(i + baseIndex).SetInt(v.(int64))
			}
		}

		res := vSrc.Field(nSrc - 1).Interface()
		fmt.Printf("type: %v", reflect.TypeOf(&res))
		fmt.Printf("type: %v", reflect.TypeOf(res))

	} else if tSrc.Kind() == reflect.Slice {
		tSrc = reflect.TypeOf(&src)
		fmt.Printf("sdf: %v", tSrc)
		//FlattenStruct(&src[0], &dst, nSrc-1)
	}
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
	VersionStr                string        `json:"version_str"`
	NodeIDStr                 string        `json:"node_id_str"`
	EncodingPath              string        `json:"encoding_path"`
	CollectionID              int           `json:"collection_id"`
	CollectionStartTime       string        `json:"collection_start_time"`
	CollectionEndTime         string        `json:"collection_end_time"`
	MsgTimestamp              string        `json:"msg_timestamp"`
	SubscriptionID            string        `json:"subscription_id"`
	SensorGroupID             []interface{} `json:"sensor_group_id"`
	DataSource                string        `json:"data_source"`
	DataVrfName               string        `json:"data.vrfName"`
	DataAddress               string        `json:"data.address"`
	DataMaskLen               int           `json:"data.maskLen"`
	DataL3NextHopCount        int           `json:"data.l3NextHopCount"`
	DataNexhHopEventType      string        `json:"data.eventType"`
	DataNexhHopNextHopAddress string        `json:"data.nexthop.address"`
	DataNexhHopOutInterface   string        `json:"data.nexthop.outInterface"`
	DataNexhHopNextHopVrfName string        `json:"data.nexthop.vrfName"`
	DataNexhHopOwner          string        `json:"data.nexthop.owner"`
	DataNexhHopPreference     int           `json:"data.nexthop.preference"`
	DataNexhHopMetric         int           `json:"data.nexthop.metric"`
	DataNexhHopTag            int           `json:"data.nexthop.tag"`
	DataNexhHopSegmentID      int           `json:"data.nexthop.segmentId"`
	DataNexhHopTunnelID       int           `json:"data.nexthop.tunnelId"`
	DataNexhHopEncapType      string        `json:"data.nexthop.encapType"`
	DataNexhHopNhTypeFlags    int           `json:"data.nexthop.nhTypeFlags"`
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

		var newRib RibGeneric

		var index int = 0

		FlattenStruct(&response, &newRib, index)
	}
}

func main() {

	http.HandleFunc("/network/rib", ribhandler)
	http.ListenAndServe(":10000", nil)
}
