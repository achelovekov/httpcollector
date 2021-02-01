package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func worker(src map[string]interface{}) {

	var path = []string{"nvoEps", "nvoEp", "nvoNws", "nvoNw"}
	var pathIndex int

	newHeader := make(map[string]interface{})

	for k, v := range src {
		if k != "data" {
			newHeader[k] = v
		}
	}

	flatten(src["data"].(map[string]interface{}), path, pathIndex, newHeader)

}

func flatten(src map[string]interface{}, path []string, pathIndex int, header map[string]interface{}) {
	newHeader := make(map[string]interface{})
	for k, v := range header {
		newHeader[k] = v
	}

	src = src[path[pathIndex]].(map[string]interface{})
	if v, ok := src["attributes"]; ok {
		for k, v := range v.(map[string]interface{}) {
			newHeader[path[pathIndex]+"."+k] = v
		}
	}
	if v, ok := src["children"]; ok && pathIndex != len(path)-1 {
		for i := 0; i < reflect.ValueOf(v).Len(); i++ {
			v := reflect.ValueOf(v).Index(i).Interface().(map[string]interface{})
			if _, ok := v[path[pathIndex+1]]; ok {
				flatten(v, path, pathIndex+1, newHeader)
			}
		}
	} else {
		fmt.Println(newHeader)
	}

}

func main() {

	var Raw = `{
		"collection_end_time": "0",
		"collection_id": 8753,
		"collection_start_time": "0",
		"data": {
		  "nvoEps": {
			"attributes": {
			  "childAction": "",
			  "dn": "sys/eps",
			  "modTs": "2021-01-11T11:52:27.087+00:00",
			  "status": ""
			},
			"children": [
			  {
				"nvoEp": {
				  "attributes": {
					"adminSt": "enabled",
					"adminStMultisiteBrdrGwLoopackIntf": "disabled",
					"advertiseVmac": false,
					"anycastIntf": "unspecified",
					"autoRemapReplicationServers": false,
					"cfgSrc": "unknown",
					"childAction": "",
					"controllerId": 0,
					"delayRestoreMsBrdrGwExpiryTime": "NA",
					"descr": "",
					"encapType": "unknown",
					"epId": 1,
					"holdDownTime": 500,
					"holdDownTimerExpiryTime": "NA",
					"holdUpTime": 30,
					"holdUpTimerExpiryTime": "NA",
					"hostReach": "bgp",
					"ingressReplProtoBGP": false,
					"learningMode": "CP",
					"mac": "00:3A:9C:5B:6A:87",
					"mcastGroupL2": "0.0.0.0",
					"mcastGroupL3": "0.0.0.0",
					"modTs": "2021-01-11T12:12:29.611+00:00",
					"multisiteBordergwInterface": "unspecified",
					"multisiteBrdrGwIntfIp": "0.0.0.0",
					"operEncapType": "vxlan",
					"operStAnycastSrcIntf": "down",
					"operStMultisiteBrdrGwLoopbackIntf": "down",
					"operStSrcLoopbackIntf": "up",
					"operState": "up",
					"primaryIp": "172.17.1.2",
					"primaryIpv6": "",
					"rn": "epId-1",
					"secondaryIp": "172.17.1.101",
					"secondaryIpv6": "",
					"sourceInterface": "lo1",
					"status": "",
					"suppressARP": false,
					"virtualMac": "00:00:00:00:00:00",
					"virtualRtrMac": "02:00:AC:11:01:65",
					"virtualRtrMacReorig": "00:00:00:00:00:00",
					"vpcVIPNotified": true
				  },
				  "children": [
					{
					  "nvoPeers": {
						"attributes": {
						  "childAction": "",
						  "modTs": "2021-01-11T11:52:27.839+00:00",
						  "rn": "peers",
						  "status": ""
						},
						"children": [
						  {
							"nvoDyPeer": {
							  "attributes": {
								"createTs": "2021-01-11T19:09:01.979+00:00",
								"firstVNI": 2012007,
								"ip": "172.17.1.203",
								"mac": "BC:4A:56:CD:EC:3F",
								"rn": "dy_peer-[172.17.1.203]",
								"state": "Up",
								"upStateTransitionTs": "2021-01-11T19:09:02.012+00:00"
							  }
							}
						  },
						  {
							"nvoDyPeer": {
							  "attributes": {
								"createTs": "2021-01-11T19:09:01.978+00:00",
								"firstVNI": 2012007,
								"ip": "172.17.1.202",
								"mac": "D4:C9:3C:85:62:7F",
								"rn": "dy_peer-[172.17.1.202]",
								"state": "Up",
								"upStateTransitionTs": "2021-01-11T19:09:02.009+00:00"
							  }
							}
						  },
						  {
							"nvoDyPeer": {
							  "attributes": {
								"createTs": "2021-01-11T19:09:01.976+00:00",
								"firstVNI": 2012007,
								"ip": "172.17.1.201",
								"mac": "D4:C9:3C:85:62:3F",
								"rn": "dy_peer-[172.17.1.201]",
								"state": "Up",
								"upStateTransitionTs": "2021-01-11T19:09:02.008+00:00"
							  }
							}
						  }
						]
					  }
					},
					{
					  "nvoNws": {
						"attributes": {
						  "childAction": "",
						  "modTs": "2021-01-11T11:52:27.839+00:00",
						  "rn": "nws",
						  "status": ""
						},
						"children": [
						  {
							"nvoOperNw": {
							  "attributes": {
								"epId": 1,
								"mode": "CP",
								"operMcastGroup": "225.1.0.1/32",
								"operSupprARP": false,
								"rn": "opervni-2012008",
								"state": "Up",
								"type": "L2",
								"vlanBD": "vlan-2008",
								"vni": 2012008
							  },
							  "children": [
								{
								  "nvoCounters": {
									"attributes": {
									  "rn": "cntrs",
									  "rxMcastbytes": 1624,
									  "rxMcastpkts": 11,
									  "rxUcastbytes": 0,
									  "rxUcastpkts": 0,
									  "txMcastbytes": 504,
									  "txMcastpkts": 4,
									  "txUcastbytes": 0,
									  "txUcastpkts": 0
									}
								  }
								}
							  ]
							}
						  },
						  {
							"nvoOperNw": {
							  "attributes": {
								"epId": 1,
								"mode": "CP",
								"operMcastGroup": "0.0.0.0/32",
								"operSupprARP": false,
								"rn": "opervni-10500",
								"state": "Down",
								"type": "L3",
								"vlanBD": "vlan-1500",
								"vni": 10500
							  }
							}
						  },
						  {
							"nvoOperNw": {
							  "attributes": {
								"epId": 1,
								"mode": "CP",
								"operMcastGroup": "0.0.0.0/32",
								"operSupprARP": false,
								"rn": "opervni-10501",
								"state": "Down",
								"type": "L2",
								"vlanBD": "vlan-1501",
								"vni": 10501
							  }
							}
						  },
						  {
							"nvoOperNw": {
							  "attributes": {
								"epId": 1,
								"mode": "CP",
								"operMcastGroup": "0.0.0.0/32",
								"operSupprARP": false,
								"rn": "opervni-10502",
								"state": "Down",
								"type": "L2",
								"vlanBD": "vlan-1502",
								"vni": 10502
							  }
							}
						  },
						  {
							"nvoOperNw": {
							  "attributes": {
								"epId": 1,
								"mode": "CP",
								"operMcastGroup": "225.1.0.1/32",
								"operSupprARP": false,
								"rn": "opervni-2012007",
								"state": "Up",
								"type": "L2",
								"vlanBD": "vlan-2009",
								"vni": 2012007
							  },
							  "children": [
								{
								  "nvoCounters": {
									"attributes": {
									  "rn": "cntrs",
									  "rxMcastbytes": 215704,
									  "rxMcastpkts": 1828,
									  "rxUcastbytes": 0,
									  "rxUcastpkts": 0,
									  "txMcastbytes": 5020274,
									  "txMcastpkts": 43257,
									  "txUcastbytes": 0,
									  "txUcastpkts": 0
									}
								  }
								}
							  ]
							}
						  },
						  {
							"nvoOperNw": {
							  "attributes": {
								"epId": 1,
								"mode": "CP",
								"operMcastGroup": "0.0.0.0/32",
								"operSupprARP": false,
								"rn": "opervni-2010000",
								"state": "Up",
								"type": "L3",
								"vlanBD": "vlan-3901",
								"vni": 2010000
							  },
							  "children": [
								{
								  "nvoCounters": {
									"attributes": {
									  "rn": "cntrs",
									  "rxMcastbytes": 0,
									  "rxMcastpkts": 0,
									  "rxUcastbytes": 0,
									  "rxUcastpkts": 0,
									  "txMcastbytes": 0,
									  "txMcastpkts": 0,
									  "txUcastbytes": 0,
									  "txUcastpkts": 0
									}
								  }
								}
							  ]
							}
						  },
						  {
							"nvoNw": {
							  "attributes": {
								"associateVrfFlag": false,
								"childAction": "",
								"isLegacyMode": false,
								"mcastGroup": "225.1.0.1",
								"modTs": "2021-01-11T12:12:29.672+00:00",
								"multisiteIngRepl": "disable",
								"rn": "vni-2012008",
								"status": "",
								"suppressARP": "off",
								"vni": 2012008
							  }
							}
						  },
						  {
							"nvoNw": {
							  "attributes": {
								"associateVrfFlag": false,
								"childAction": "",
								"isLegacyMode": false,
								"mcastGroup": "225.1.0.1",
								"modTs": "2021-01-11T12:12:29.648+00:00",
								"multisiteIngRepl": "disable",
								"rn": "vni-2012007",
								"status": "",
								"suppressARP": "off",
								"vni": 2012007
							  }
							}
						  },
						  {
							"nvoNw": {
							  "attributes": {
								"associateVrfFlag": true,
								"childAction": "",
								"isLegacyMode": false,
								"mcastGroup": "0.0.0.0",
								"modTs": "2021-01-11T12:12:29.623+00:00",
								"multisiteIngRepl": "disable",
								"rn": "vni-2010000",
								"status": "",
								"suppressARP": "off",
								"vni": 2010000
							  }
							}
						  }
						]
					  }
					},
					{
					  "nvoCounters": {
						"attributes": {
						  "rn": "cntrs",
						  "rxMcastbytes": 217328,
						  "rxMcastpkts": 1839,
						  "rxUcastbytes": 0,
						  "rxUcastpkts": 0,
						  "txMcastbytes": 5020778,
						  "txMcastpkts": 43261,
						  "txUcastbytes": 0,
						  "txUcastpkts": 0
						}
					  }
					}
				  ]
				}
			  }
			]
		  }
		},
		"data_source": "DME",
		"encoding_path": "vxlan:sys/eps",
		"msg_timestamp": "1610384112529",
		"node_id_str": "site-1-ac-2",
		"sensor_group_id": [],
		"subscription_id": "1000",
		"version_str": "1.0.0"
	  }
`

	src := make(map[string]interface{})
	err := json.Unmarshal([]byte(Raw), &src)
	if err != nil {
		panic(err)
	}

	//flatten(src, path, pathIndex, header)

	worker(src)

}
