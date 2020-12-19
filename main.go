package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

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

type RibDataNextHop struct {
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
}

type RibData struct {
	VrfName        string           `json:"vrfName"`
	Address        string           `json:"address"`
	MaskLen        int              `json:"maskLen"`
	L3NextHopCount int              `json:"l3NextHopCount"`
	EventType      string           `json:"eventType"`
	NextHop        []RibDataNextHop `json:"nextHop"`
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
	Data                []RibData     `json:"data"`
}

func ribhandler(w http.ResponseWriter, r *http.Request) {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://10.62.186.54:9200",
		},
	}

	es, _ := elasticsearch.NewClient(cfg)

	switch r.Method {
	case "POST":
		data, _ := ioutil.ReadAll(r.Body)

		var response Rib
		err := json.Unmarshal([]byte(data), &response)

		if err != nil {
			log.Print(err)
		}

		for _, data := range response.Data {

			ribData := make([]RibData, 1)
			ribDataNextHop := make([]RibDataNextHop, 1)

			if len(data.NextHop) > 0 {
				for _, nexthop := range data.NextHop {
					ribDataNextHop[0] = RibDataNextHop{
						Address:      nexthop.Address,
						OutInterface: nexthop.OutInterface,
						VrfName:      nexthop.VrfName,
						Owner:        nexthop.Owner,
						Preference:   nexthop.Preference,
						Metric:       nexthop.Metric,
						Tag:          nexthop.Tag,
						SegmentID:    nexthop.SegmentID,
						TunnelID:     nexthop.TunnelID,
						EncapType:    nexthop.EncapType,
						NhTypeFlags:  nexthop.NhTypeFlags,
					}
				}
			}
			ribData[0] = RibData{
				VrfName:        data.VrfName,
				Address:        data.Address,
				MaskLen:        data.MaskLen,
				L3NextHopCount: data.L3NextHopCount,
				EventType:      data.EventType,
				NextHop:        ribDataNextHop,
			}

			var newRib = RibGeneric{
				VersionStr:          response.VersionStr,
				NodeIDStr:           response.NodeIDStr,
				EncodingPath:        response.EncodingPath,
				CollectionID:        response.CollectionID,
				CollectionStartTime: response.CollectionStartTime,
				CollectionEndTime:   response.CollectionEndTime,
				MsgTimestamp:        response.MsgTimestamp,
				SubscriptionID:      response.SubscriptionID,
				SensorGroupID:       response.SensorGroupID,
				DataSource:          response.DataSource,
				Data:                ribData,
			}
			empJSON, err := json.MarshalIndent(newRib, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Println(string(empJSON))
			res, err := es.Index("golang-index", strings.NewReader(string(empJSON)))
			if err != nil {
				log.Fatalf("ERROR: %s", err)
				log.Println(res)
			}
		}
	}

	/*
		empJSON, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println(string(empJSON))
	*/
	/*
		cfg := elasticsearch.Config{
			Addresses: []string{
				"http://10.62.186.54:9200",
			},
		}

		es, _ := elasticsearch.NewClient(cfg)

		res, err := es.Index("golang-index", strings.NewReader(string(empJSON)))
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
		defer res.Body.Close()

		log.Println(res)
	*/
}

func main() {

	http.HandleFunc("/network/rib", ribhandler)
	http.ListenAndServe(":10000", nil)
}
