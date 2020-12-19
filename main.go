package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func ribhandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		data, _ := ioutil.ReadAll(r.Body)

		var response Rib
		err := json.Unmarshal([]byte(data), &response)

		if err != nil {
			log.Print(err)
		}

		for i, v := range response.Data {
			fmt.Println(response.VersionStr)
			fmt.Println(response.NodeIDStr)
			fmt.Println(response.EncodingPath)
			fmt.Println(response.CollectionID)
			fmt.Println(response.CollectionStartTime)
			fmt.Println(response.CollectionEndTime)
			fmt.Println(response.MsgTimestamp)
			fmt.Println(response.SubscriptionID)
			fmt.Println(response.SensorGroupID)
			fmt.Println(response.DataSource)
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
}

func main() {

	http.HandleFunc("/network/rib", ribhandler)
	http.ListenAndServe(":10000", nil)
}
