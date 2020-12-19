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

/*
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
	CollectionStartTime string        `json:"blablabla.collection_start_time"`
	CollectionEndTime   string        `json:"collection_end_time"`
	MsgTimestamp        string        `json:"msg_timestamp"`
	SubscriptionID      string        `json:"subscription_id"`
	SensorGroupID       []interface{} `json:"sensor_group_id"`
	DataSource          string        `json:"data_source"`
	Data                []RibData     `json:"data"`
}
*/

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

			var newRib RibGeneric

			newRib.VersionStr = response.VersionStr
			newRib.NodeIDStr = response.NodeIDStr
			newRib.EncodingPath = response.EncodingPath
			newRib.CollectionID = response.CollectionID
			newRib.CollectionStartTime = response.CollectionStartTime
			newRib.CollectionEndTime = response.CollectionEndTime
			newRib.MsgTimestamp = response.MsgTimestamp
			newRib.SubscriptionID = response.SubscriptionID
			newRib.SensorGroupID = response.SensorGroupID
			newRib.DataSource = response.DataSource
			newRib.VrfName = data.VrfName
			newRib.Address = data.Address
			newRib.MaskLen = data.MaskLen
			newRib.L3NextHopCount = data.L3NextHopCount
			newRib.EventType = data.EventType

			if len(data.NextHop) > 0 {
				for _, nexthop := range data.NextHop {
					newRib.NextHopAddress = nexthop.Address
					newRib.OutInterface = nexthop.OutInterface
					newRib.NextHopVrfName = nexthop.VrfName
					newRib.Owner = nexthop.Owner
					newRib.Preference = nexthop.Preference
					newRib.Metric = nexthop.Metric
					newRib.Tag = nexthop.Tag
					newRib.SegmentID = nexthop.SegmentID
					newRib.TunnelID = nexthop.TunnelID
					newRib.EncapType = nexthop.EncapType
					newRib.NhTypeFlags = nexthop.NhTypeFlags

					empJSON, err := json.MarshalIndent(newRib, "", "  ")
					if err != nil {
						log.Fatalf(err.Error())
					}
					fmt.Println(string(empJSON))
					res, err := es.Index("golang-index", strings.NewReader(string(empJSON)))
					if err != nil {
						log.Fatalf("ERROR: %s", err)
					}
					log.Println(res)
				}
			}
			empJSON, err := json.MarshalIndent(newRib, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Println(string(empJSON))
			res, err := es.Index("golang-index", strings.NewReader(string(empJSON)))
			if err != nil {
				log.Fatalf("ERROR: %s", err)
			}
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

func main() {

	http.HandleFunc("/network/rib", ribhandler)
	http.ListenAndServe(":10000", nil)
}
