package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func flatten(src map[string]interface{}, path []string, pathIndex int, header map[string]interface{}) {
	newHeader := make(map[string]interface{})
	for k, v := range header {
		newHeader[k] = v
	}

	for k, v := range src[path[pathIndex]].(map[string]interface{}) {
		if k == "attributes" {
			for k, v := range v.(map[string]interface{}) {
				newHeader[path[pathIndex]+"."+k] = v
			}
		} else if k == "children" && pathIndex != len(path)-1 {
			for i := 0; i < reflect.ValueOf(v).Len(); i++ {
				v := reflect.ValueOf(v).Index(i).Interface().(map[string]interface{})
				if _, ok := v[path[pathIndex+1]]; ok {
					flatten(v, path, pathIndex+1, newHeader)
				}
			}
		}
	}

	if pathIndex == len(path)-1 {
		fmt.Println(newHeader)
	}
}

func main() {

	var Raw = `
	{
		"bgpEntity": {
		  "attributes": {
			"adminSt": "enabled",
			"childAction": "",
			"dn": "sys/bgp",
			"modTs": "2020-10-01T08:53:57.611+00:00",
			"name": "bgp",
			"operSt": "enabled",
			"status": ""
		  },
		  "children": [
			{
			  "bgpInst": {
				"attributes": {
				  "activateTs": "2020-11-05T06:40:39.203+00:00",
				  "adminSt": "enabled",
				  "affGrpActv": "0",
				  "asPathDbSz": "30",
				  "asn": "1",
				  "attribDbSz": "933720",
				  "childAction": "",
				  "createTs": "2020-11-05T06:39:42.664+00:00",
				  "ctrl": "fastExtFallover",
				  "disPolBatch": "disabled",
				  "disPolBatchv4PfxLst": "",
				  "disPolBatchv6PfxLst": "",
				  "enhancedErr": "yes",
				  "epeActivePeers": "0",
				  "epeConfiguredPeers": "0",
				  "fabricSoo": "unknown:unknown:0:0",
				  "flushRoutes": "disabled",
				  "isolate": "disabled",
				  "lnkStClnt": "inactive",
				  "lnkStSrvr": "inactive",
				  "medDampIntvl": "0",
				  "memAlert": "normal",
				  "modTs": "2020-10-01T08:53:57.611+00:00",
				  "name": "bgp",
				  "nhSupprDefRes": "disabled",
				  "numAsPath": "1",
				  "numRtAttrib": "7530",
				  "operErr": "",
				  "rn": "inst",
				  "srgbMaxLbl": "none",
				  "srgbMinLbl": "none",
				  "status": "",
				  "waitDoneTs": "2020-11-05T06:45:47.716+00:00"
				},
				"children": [
				  {
					"bgpOperBgp": {
					  "attributes": {
						"rn": "oper"
					  },
					  "children": [
						{
						  "bgpOperRtctrlL3": {
							"attributes": {
							  "encap": "vxlan-2010000",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:3",
							  "rn": "l3-[vxlan-2010000]"
							},
							"children": [
							  {
								"bgpOperDomAf": {
								  "attributes": {
									"name": "",
									"rn": "af-ipv6-ucast",
									"type": "ipv6-ucast"
								  },
								  "children": [
									{
									  "bgpOperAfCtrl": {
										"attributes": {
										  "name": "",
										  "rn": "ctrl-ipv6-mvpn",
										  "type": "ipv6-mvpn"
										},
										"children": [
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-import",
												"type": "import"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:ipv4-nn2:172.17.1.1:3901",
													  "rtt": "route-target:ipv4-nn2:172.17.1.1:3901"
													}
												  }
												},
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  },
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-export",
												"type": "export"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  }
										]
									  }
									},
									{
									  "bgpOperAfCtrl": {
										"attributes": {
										  "name": "",
										  "rn": "ctrl-l2vpn-evpn",
										  "type": "l2vpn-evpn"
										},
										"children": [
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-export",
												"type": "export"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  },
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-import",
												"type": "import"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  }
										]
									  }
									},
									{
									  "bgpOperAfCtrl": {
										"attributes": {
										  "name": "",
										  "rn": "ctrl-ipv6-ucast",
										  "type": "ipv6-ucast"
										},
										"children": [
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-export",
												"type": "export"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  },
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-import",
												"type": "import"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  }
										]
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperDomAf": {
								  "attributes": {
									"name": "",
									"rn": "af-ipv4-ucast",
									"type": "ipv4-ucast"
								  },
								  "children": [
									{
									  "bgpOperAfCtrl": {
										"attributes": {
										  "name": "",
										  "rn": "ctrl-ipv4-mvpn",
										  "type": "ipv4-mvpn"
										},
										"children": [
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-import",
												"type": "import"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:ipv4-nn2:172.17.1.1:3901",
													  "rtt": "route-target:ipv4-nn2:172.17.1.1:3901"
													}
												  }
												},
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  },
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-export",
												"type": "export"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  }
										]
									  }
									},
									{
									  "bgpOperAfCtrl": {
										"attributes": {
										  "name": "",
										  "rn": "ctrl-l2vpn-evpn",
										  "type": "l2vpn-evpn"
										},
										"children": [
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-export",
												"type": "export"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  },
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-import",
												"type": "import"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  }
										]
									  }
									},
									{
									  "bgpOperAfCtrl": {
										"attributes": {
										  "name": "",
										  "rn": "ctrl-ipv4-ucast",
										  "type": "ipv4-ucast"
										},
										"children": [
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-export",
												"type": "export"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  },
										  {
											"bgpOperRttP": {
											  "attributes": {
												"descr": "",
												"name": "",
												"rn": "rttp-import",
												"type": "import"
											  },
											  "children": [
												{
												  "bgpOperRttEntry": {
													"attributes": {
													  "rn": "entry-route-target:as2-nn2:1:2010000",
													  "rtt": "route-target:as2-nn2:1:2010000"
													}
												  }
												}
											  ]
											}
										  }
										]
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331477",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34244",
							  "rn": "l2-[vxlan-3331477]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331477",
										  "rtt": "route-target:as2-nn2:65000:3331477"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331477",
										  "rtt": "route-target:as2-nn2:65000:3331477"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331476",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34243",
							  "rn": "l2-[vxlan-3331476]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331476",
										  "rtt": "route-target:as2-nn2:65000:3331476"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331476",
										  "rtt": "route-target:as2-nn2:65000:3331476"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331470",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34237",
							  "rn": "l2-[vxlan-3331470]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331470",
										  "rtt": "route-target:as2-nn2:65000:3331470"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331470",
										  "rtt": "route-target:as2-nn2:65000:3331470"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331469",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34236",
							  "rn": "l2-[vxlan-3331469]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331469",
										  "rtt": "route-target:as2-nn2:65000:3331469"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331469",
										  "rtt": "route-target:as2-nn2:65000:3331469"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331465",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34232",
							  "rn": "l2-[vxlan-3331465]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331465",
										  "rtt": "route-target:as2-nn2:65000:3331465"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331465",
										  "rtt": "route-target:as2-nn2:65000:3331465"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331463",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34230",
							  "rn": "l2-[vxlan-3331463]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331463",
										  "rtt": "route-target:as2-nn2:65000:3331463"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331463",
										  "rtt": "route-target:as2-nn2:65000:3331463"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331459",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34226",
							  "rn": "l2-[vxlan-3331459]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331459",
										  "rtt": "route-target:as2-nn2:65000:3331459"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331459",
										  "rtt": "route-target:as2-nn2:65000:3331459"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331457",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34224",
							  "rn": "l2-[vxlan-3331457]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331457",
										  "rtt": "route-target:as2-nn2:65000:3331457"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331457",
										  "rtt": "route-target:as2-nn2:65000:3331457"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331455",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34222",
							  "rn": "l2-[vxlan-3331455]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331455",
										  "rtt": "route-target:as2-nn2:65000:3331455"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331455",
										  "rtt": "route-target:as2-nn2:65000:3331455"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331454",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34221",
							  "rn": "l2-[vxlan-3331454]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331454",
										  "rtt": "route-target:as2-nn2:65000:3331454"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331454",
										  "rtt": "route-target:as2-nn2:65000:3331454"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331452",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34219",
							  "rn": "l2-[vxlan-3331452]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331452",
										  "rtt": "route-target:as2-nn2:65000:3331452"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331452",
										  "rtt": "route-target:as2-nn2:65000:3331452"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331447",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34214",
							  "rn": "l2-[vxlan-3331447]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331447",
										  "rtt": "route-target:as2-nn2:65000:3331447"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331447",
										  "rtt": "route-target:as2-nn2:65000:3331447"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331445",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34212",
							  "rn": "l2-[vxlan-3331445]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331445",
										  "rtt": "route-target:as2-nn2:65000:3331445"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331445",
										  "rtt": "route-target:as2-nn2:65000:3331445"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331443",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34210",
							  "rn": "l2-[vxlan-3331443]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331443",
										  "rtt": "route-target:as2-nn2:65000:3331443"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331443",
										  "rtt": "route-target:as2-nn2:65000:3331443"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331442",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34209",
							  "rn": "l2-[vxlan-3331442]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331442",
										  "rtt": "route-target:as2-nn2:65000:3331442"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331442",
										  "rtt": "route-target:as2-nn2:65000:3331442"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331440",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34207",
							  "rn": "l2-[vxlan-3331440]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331440",
										  "rtt": "route-target:as2-nn2:65000:3331440"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331440",
										  "rtt": "route-target:as2-nn2:65000:3331440"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331437",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34204",
							  "rn": "l2-[vxlan-3331437]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331437",
										  "rtt": "route-target:as2-nn2:65000:3331437"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331437",
										  "rtt": "route-target:as2-nn2:65000:3331437"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331435",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34202",
							  "rn": "l2-[vxlan-3331435]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331435",
										  "rtt": "route-target:as2-nn2:65000:3331435"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331435",
										  "rtt": "route-target:as2-nn2:65000:3331435"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331428",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34195",
							  "rn": "l2-[vxlan-3331428]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331428",
										  "rtt": "route-target:as2-nn2:65000:3331428"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331428",
										  "rtt": "route-target:as2-nn2:65000:3331428"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331426",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34193",
							  "rn": "l2-[vxlan-3331426]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331426",
										  "rtt": "route-target:as2-nn2:65000:3331426"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331426",
										  "rtt": "route-target:as2-nn2:65000:3331426"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331425",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34192",
							  "rn": "l2-[vxlan-3331425]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331425",
										  "rtt": "route-target:as2-nn2:65000:3331425"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331425",
										  "rtt": "route-target:as2-nn2:65000:3331425"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331423",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34190",
							  "rn": "l2-[vxlan-3331423]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331423",
										  "rtt": "route-target:as2-nn2:65000:3331423"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331423",
										  "rtt": "route-target:as2-nn2:65000:3331423"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331422",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34189",
							  "rn": "l2-[vxlan-3331422]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331422",
										  "rtt": "route-target:as2-nn2:65000:3331422"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331422",
										  "rtt": "route-target:as2-nn2:65000:3331422"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331419",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34186",
							  "rn": "l2-[vxlan-3331419]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331419",
										  "rtt": "route-target:as2-nn2:65000:3331419"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331419",
										  "rtt": "route-target:as2-nn2:65000:3331419"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331418",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34185",
							  "rn": "l2-[vxlan-3331418]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331418",
										  "rtt": "route-target:as2-nn2:65000:3331418"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331418",
										  "rtt": "route-target:as2-nn2:65000:3331418"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331417",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34184",
							  "rn": "l2-[vxlan-3331417]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331417",
										  "rtt": "route-target:as2-nn2:65000:3331417"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331417",
										  "rtt": "route-target:as2-nn2:65000:3331417"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331414",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34181",
							  "rn": "l2-[vxlan-3331414]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331414",
										  "rtt": "route-target:as2-nn2:65000:3331414"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331414",
										  "rtt": "route-target:as2-nn2:65000:3331414"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331412",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34179",
							  "rn": "l2-[vxlan-3331412]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331412",
										  "rtt": "route-target:as2-nn2:65000:3331412"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331412",
										  "rtt": "route-target:as2-nn2:65000:3331412"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331411",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34178",
							  "rn": "l2-[vxlan-3331411]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331411",
										  "rtt": "route-target:as2-nn2:65000:3331411"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331411",
										  "rtt": "route-target:as2-nn2:65000:3331411"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331436",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34203",
							  "rn": "l2-[vxlan-3331436]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331436",
										  "rtt": "route-target:as2-nn2:65000:3331436"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331436",
										  "rtt": "route-target:as2-nn2:65000:3331436"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331408",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34175",
							  "rn": "l2-[vxlan-3331408]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331408",
										  "rtt": "route-target:as2-nn2:65000:3331408"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331408",
										  "rtt": "route-target:as2-nn2:65000:3331408"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331407",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34174",
							  "rn": "l2-[vxlan-3331407]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331407",
										  "rtt": "route-target:as2-nn2:65000:3331407"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331407",
										  "rtt": "route-target:as2-nn2:65000:3331407"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331475",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34242",
							  "rn": "l2-[vxlan-3331475]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331475",
										  "rtt": "route-target:as2-nn2:65000:3331475"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331475",
										  "rtt": "route-target:as2-nn2:65000:3331475"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331420",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34187",
							  "rn": "l2-[vxlan-3331420]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331420",
										  "rtt": "route-target:as2-nn2:65000:3331420"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331420",
										  "rtt": "route-target:as2-nn2:65000:3331420"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331403",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34170",
							  "rn": "l2-[vxlan-3331403]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331403",
										  "rtt": "route-target:as2-nn2:65000:3331403"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331403",
										  "rtt": "route-target:as2-nn2:65000:3331403"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331402",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34169",
							  "rn": "l2-[vxlan-3331402]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331402",
										  "rtt": "route-target:as2-nn2:65000:3331402"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331402",
										  "rtt": "route-target:as2-nn2:65000:3331402"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331401",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34168",
							  "rn": "l2-[vxlan-3331401]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331401",
										  "rtt": "route-target:as2-nn2:65000:3331401"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331401",
										  "rtt": "route-target:as2-nn2:65000:3331401"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331400",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34167",
							  "rn": "l2-[vxlan-3331400]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331400",
										  "rtt": "route-target:as2-nn2:65000:3331400"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331400",
										  "rtt": "route-target:as2-nn2:65000:3331400"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331399",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34166",
							  "rn": "l2-[vxlan-3331399]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331399",
										  "rtt": "route-target:as2-nn2:65000:3331399"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331399",
										  "rtt": "route-target:as2-nn2:65000:3331399"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331398",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34165",
							  "rn": "l2-[vxlan-3331398]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331398",
										  "rtt": "route-target:as2-nn2:65000:3331398"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331398",
										  "rtt": "route-target:as2-nn2:65000:3331398"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331397",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34164",
							  "rn": "l2-[vxlan-3331397]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331397",
										  "rtt": "route-target:as2-nn2:65000:3331397"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331397",
										  "rtt": "route-target:as2-nn2:65000:3331397"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331396",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34163",
							  "rn": "l2-[vxlan-3331396]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331396",
										  "rtt": "route-target:as2-nn2:65000:3331396"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331396",
										  "rtt": "route-target:as2-nn2:65000:3331396"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331393",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34160",
							  "rn": "l2-[vxlan-3331393]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331393",
										  "rtt": "route-target:as2-nn2:65000:3331393"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331393",
										  "rtt": "route-target:as2-nn2:65000:3331393"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331391",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34158",
							  "rn": "l2-[vxlan-3331391]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331391",
										  "rtt": "route-target:as2-nn2:65000:3331391"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331391",
										  "rtt": "route-target:as2-nn2:65000:3331391"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331390",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34157",
							  "rn": "l2-[vxlan-3331390]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331390",
										  "rtt": "route-target:as2-nn2:65000:3331390"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331390",
										  "rtt": "route-target:as2-nn2:65000:3331390"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331387",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34154",
							  "rn": "l2-[vxlan-3331387]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331387",
										  "rtt": "route-target:as2-nn2:65000:3331387"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331387",
										  "rtt": "route-target:as2-nn2:65000:3331387"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331386",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34153",
							  "rn": "l2-[vxlan-3331386]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331386",
										  "rtt": "route-target:as2-nn2:65000:3331386"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331386",
										  "rtt": "route-target:as2-nn2:65000:3331386"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331384",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34151",
							  "rn": "l2-[vxlan-3331384]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331384",
										  "rtt": "route-target:as2-nn2:65000:3331384"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331384",
										  "rtt": "route-target:as2-nn2:65000:3331384"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331381",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34148",
							  "rn": "l2-[vxlan-3331381]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331381",
										  "rtt": "route-target:as2-nn2:65000:3331381"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331381",
										  "rtt": "route-target:as2-nn2:65000:3331381"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331456",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34223",
							  "rn": "l2-[vxlan-3331456]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331456",
										  "rtt": "route-target:as2-nn2:65000:3331456"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331456",
										  "rtt": "route-target:as2-nn2:65000:3331456"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331378",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34145",
							  "rn": "l2-[vxlan-3331378]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331378",
										  "rtt": "route-target:as2-nn2:65000:3331378"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331378",
										  "rtt": "route-target:as2-nn2:65000:3331378"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331377",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34144",
							  "rn": "l2-[vxlan-3331377]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331377",
										  "rtt": "route-target:as2-nn2:65000:3331377"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331377",
										  "rtt": "route-target:as2-nn2:65000:3331377"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331392",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34159",
							  "rn": "l2-[vxlan-3331392]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331392",
										  "rtt": "route-target:as2-nn2:65000:3331392"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331392",
										  "rtt": "route-target:as2-nn2:65000:3331392"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331376",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34143",
							  "rn": "l2-[vxlan-3331376]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331376",
										  "rtt": "route-target:as2-nn2:65000:3331376"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331376",
										  "rtt": "route-target:as2-nn2:65000:3331376"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331375",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34142",
							  "rn": "l2-[vxlan-3331375]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331375",
										  "rtt": "route-target:as2-nn2:65000:3331375"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331375",
										  "rtt": "route-target:as2-nn2:65000:3331375"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331374",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34141",
							  "rn": "l2-[vxlan-3331374]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331374",
										  "rtt": "route-target:as2-nn2:65000:3331374"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331374",
										  "rtt": "route-target:as2-nn2:65000:3331374"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331372",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34139",
							  "rn": "l2-[vxlan-3331372]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331372",
										  "rtt": "route-target:as2-nn2:65000:3331372"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331372",
										  "rtt": "route-target:as2-nn2:65000:3331372"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331371",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34138",
							  "rn": "l2-[vxlan-3331371]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331371",
										  "rtt": "route-target:as2-nn2:65000:3331371"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331371",
										  "rtt": "route-target:as2-nn2:65000:3331371"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331369",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34136",
							  "rn": "l2-[vxlan-3331369]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331369",
										  "rtt": "route-target:as2-nn2:65000:3331369"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331369",
										  "rtt": "route-target:as2-nn2:65000:3331369"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331368",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34135",
							  "rn": "l2-[vxlan-3331368]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331368",
										  "rtt": "route-target:as2-nn2:65000:3331368"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331368",
										  "rtt": "route-target:as2-nn2:65000:3331368"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331366",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34133",
							  "rn": "l2-[vxlan-3331366]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331366",
										  "rtt": "route-target:as2-nn2:65000:3331366"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331366",
										  "rtt": "route-target:as2-nn2:65000:3331366"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331360",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34127",
							  "rn": "l2-[vxlan-3331360]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331360",
										  "rtt": "route-target:as2-nn2:65000:3331360"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331360",
										  "rtt": "route-target:as2-nn2:65000:3331360"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331359",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34126",
							  "rn": "l2-[vxlan-3331359]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331359",
										  "rtt": "route-target:as2-nn2:65000:3331359"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331359",
										  "rtt": "route-target:as2-nn2:65000:3331359"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331354",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34121",
							  "rn": "l2-[vxlan-3331354]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331354",
										  "rtt": "route-target:as2-nn2:65000:3331354"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331354",
										  "rtt": "route-target:as2-nn2:65000:3331354"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331352",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34119",
							  "rn": "l2-[vxlan-3331352]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331352",
										  "rtt": "route-target:as2-nn2:65000:3331352"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331352",
										  "rtt": "route-target:as2-nn2:65000:3331352"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331351",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34118",
							  "rn": "l2-[vxlan-3331351]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331351",
										  "rtt": "route-target:as2-nn2:65000:3331351"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331351",
										  "rtt": "route-target:as2-nn2:65000:3331351"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331380",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34147",
							  "rn": "l2-[vxlan-3331380]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331380",
										  "rtt": "route-target:as2-nn2:65000:3331380"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331380",
										  "rtt": "route-target:as2-nn2:65000:3331380"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331349",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34116",
							  "rn": "l2-[vxlan-3331349]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331349",
										  "rtt": "route-target:as2-nn2:65000:3331349"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331349",
										  "rtt": "route-target:as2-nn2:65000:3331349"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331348",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34115",
							  "rn": "l2-[vxlan-3331348]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331348",
										  "rtt": "route-target:as2-nn2:65000:3331348"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331348",
										  "rtt": "route-target:as2-nn2:65000:3331348"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331347",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34114",
							  "rn": "l2-[vxlan-3331347]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331347",
										  "rtt": "route-target:as2-nn2:65000:3331347"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331347",
										  "rtt": "route-target:as2-nn2:65000:3331347"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331345",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34112",
							  "rn": "l2-[vxlan-3331345]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331345",
										  "rtt": "route-target:as2-nn2:65000:3331345"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331345",
										  "rtt": "route-target:as2-nn2:65000:3331345"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331404",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34171",
							  "rn": "l2-[vxlan-3331404]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331404",
										  "rtt": "route-target:as2-nn2:65000:3331404"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331404",
										  "rtt": "route-target:as2-nn2:65000:3331404"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331344",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34111",
							  "rn": "l2-[vxlan-3331344]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331344",
										  "rtt": "route-target:as2-nn2:65000:3331344"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331344",
										  "rtt": "route-target:as2-nn2:65000:3331344"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331343",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34110",
							  "rn": "l2-[vxlan-3331343]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331343",
										  "rtt": "route-target:as2-nn2:65000:3331343"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331343",
										  "rtt": "route-target:as2-nn2:65000:3331343"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331358",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34125",
							  "rn": "l2-[vxlan-3331358]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331358",
										  "rtt": "route-target:as2-nn2:65000:3331358"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331358",
										  "rtt": "route-target:as2-nn2:65000:3331358"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331338",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34105",
							  "rn": "l2-[vxlan-3331338]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331338",
										  "rtt": "route-target:as2-nn2:65000:3331338"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331338",
										  "rtt": "route-target:as2-nn2:65000:3331338"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331337",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34104",
							  "rn": "l2-[vxlan-3331337]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331337",
										  "rtt": "route-target:as2-nn2:65000:3331337"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331337",
										  "rtt": "route-target:as2-nn2:65000:3331337"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331336",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34103",
							  "rn": "l2-[vxlan-3331336]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331336",
										  "rtt": "route-target:as2-nn2:65000:3331336"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331336",
										  "rtt": "route-target:as2-nn2:65000:3331336"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331335",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34102",
							  "rn": "l2-[vxlan-3331335]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331335",
										  "rtt": "route-target:as2-nn2:65000:3331335"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331335",
										  "rtt": "route-target:as2-nn2:65000:3331335"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331334",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34101",
							  "rn": "l2-[vxlan-3331334]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331334",
										  "rtt": "route-target:as2-nn2:65000:3331334"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331334",
										  "rtt": "route-target:as2-nn2:65000:3331334"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331331",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34098",
							  "rn": "l2-[vxlan-3331331]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331331",
										  "rtt": "route-target:as2-nn2:65000:3331331"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331331",
										  "rtt": "route-target:as2-nn2:65000:3331331"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331329",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34096",
							  "rn": "l2-[vxlan-3331329]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331329",
										  "rtt": "route-target:as2-nn2:65000:3331329"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331329",
										  "rtt": "route-target:as2-nn2:65000:3331329"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331327",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34094",
							  "rn": "l2-[vxlan-3331327]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331327",
										  "rtt": "route-target:as2-nn2:65000:3331327"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331327",
										  "rtt": "route-target:as2-nn2:65000:3331327"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331326",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34093",
							  "rn": "l2-[vxlan-3331326]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331326",
										  "rtt": "route-target:as2-nn2:65000:3331326"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331326",
										  "rtt": "route-target:as2-nn2:65000:3331326"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331323",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34090",
							  "rn": "l2-[vxlan-3331323]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331323",
										  "rtt": "route-target:as2-nn2:65000:3331323"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331323",
										  "rtt": "route-target:as2-nn2:65000:3331323"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331320",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34087",
							  "rn": "l2-[vxlan-3331320]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331320",
										  "rtt": "route-target:as2-nn2:65000:3331320"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331320",
										  "rtt": "route-target:as2-nn2:65000:3331320"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331318",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34085",
							  "rn": "l2-[vxlan-3331318]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331318",
										  "rtt": "route-target:as2-nn2:65000:3331318"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331318",
										  "rtt": "route-target:as2-nn2:65000:3331318"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331315",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34082",
							  "rn": "l2-[vxlan-3331315]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331315",
										  "rtt": "route-target:as2-nn2:65000:3331315"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331315",
										  "rtt": "route-target:as2-nn2:65000:3331315"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331340",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34107",
							  "rn": "l2-[vxlan-3331340]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331340",
										  "rtt": "route-target:as2-nn2:65000:3331340"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331340",
										  "rtt": "route-target:as2-nn2:65000:3331340"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331314",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34081",
							  "rn": "l2-[vxlan-3331314]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331314",
										  "rtt": "route-target:as2-nn2:65000:3331314"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331314",
										  "rtt": "route-target:as2-nn2:65000:3331314"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331312",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34079",
							  "rn": "l2-[vxlan-3331312]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331312",
										  "rtt": "route-target:as2-nn2:65000:3331312"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331312",
										  "rtt": "route-target:as2-nn2:65000:3331312"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331394",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34161",
							  "rn": "l2-[vxlan-3331394]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331394",
										  "rtt": "route-target:as2-nn2:65000:3331394"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331394",
										  "rtt": "route-target:as2-nn2:65000:3331394"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331311",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34078",
							  "rn": "l2-[vxlan-3331311]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331311",
										  "rtt": "route-target:as2-nn2:65000:3331311"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331311",
										  "rtt": "route-target:as2-nn2:65000:3331311"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331310",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34077",
							  "rn": "l2-[vxlan-3331310]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331310",
										  "rtt": "route-target:as2-nn2:65000:3331310"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331310",
										  "rtt": "route-target:as2-nn2:65000:3331310"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331309",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34076",
							  "rn": "l2-[vxlan-3331309]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331309",
										  "rtt": "route-target:as2-nn2:65000:3331309"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331309",
										  "rtt": "route-target:as2-nn2:65000:3331309"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331307",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34074",
							  "rn": "l2-[vxlan-3331307]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331307",
										  "rtt": "route-target:as2-nn2:65000:3331307"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331307",
										  "rtt": "route-target:as2-nn2:65000:3331307"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331306",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34073",
							  "rn": "l2-[vxlan-3331306]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331306",
										  "rtt": "route-target:as2-nn2:65000:3331306"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331306",
										  "rtt": "route-target:as2-nn2:65000:3331306"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331305",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34072",
							  "rn": "l2-[vxlan-3331305]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331305",
										  "rtt": "route-target:as2-nn2:65000:3331305"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331305",
										  "rtt": "route-target:as2-nn2:65000:3331305"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331301",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34068",
							  "rn": "l2-[vxlan-3331301]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331301",
										  "rtt": "route-target:as2-nn2:65000:3331301"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331301",
										  "rtt": "route-target:as2-nn2:65000:3331301"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331300",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34067",
							  "rn": "l2-[vxlan-3331300]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331300",
										  "rtt": "route-target:as2-nn2:65000:3331300"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331300",
										  "rtt": "route-target:as2-nn2:65000:3331300"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331413",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34180",
							  "rn": "l2-[vxlan-3331413]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331413",
										  "rtt": "route-target:as2-nn2:65000:3331413"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331413",
										  "rtt": "route-target:as2-nn2:65000:3331413"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331298",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34065",
							  "rn": "l2-[vxlan-3331298]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331298",
										  "rtt": "route-target:as2-nn2:65000:3331298"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331298",
										  "rtt": "route-target:as2-nn2:65000:3331298"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331294",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34061",
							  "rn": "l2-[vxlan-3331294]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331294",
										  "rtt": "route-target:as2-nn2:65000:3331294"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331294",
										  "rtt": "route-target:as2-nn2:65000:3331294"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331292",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34059",
							  "rn": "l2-[vxlan-3331292]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331292",
										  "rtt": "route-target:as2-nn2:65000:3331292"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331292",
										  "rtt": "route-target:as2-nn2:65000:3331292"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331290",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34057",
							  "rn": "l2-[vxlan-3331290]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331290",
										  "rtt": "route-target:as2-nn2:65000:3331290"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331290",
										  "rtt": "route-target:as2-nn2:65000:3331290"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331288",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34055",
							  "rn": "l2-[vxlan-3331288]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331288",
										  "rtt": "route-target:as2-nn2:65000:3331288"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331288",
										  "rtt": "route-target:as2-nn2:65000:3331288"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331287",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34054",
							  "rn": "l2-[vxlan-3331287]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331287",
										  "rtt": "route-target:as2-nn2:65000:3331287"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331287",
										  "rtt": "route-target:as2-nn2:65000:3331287"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331286",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34053",
							  "rn": "l2-[vxlan-3331286]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331286",
										  "rtt": "route-target:as2-nn2:65000:3331286"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331286",
										  "rtt": "route-target:as2-nn2:65000:3331286"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331285",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34052",
							  "rn": "l2-[vxlan-3331285]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331285",
										  "rtt": "route-target:as2-nn2:65000:3331285"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331285",
										  "rtt": "route-target:as2-nn2:65000:3331285"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331283",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34050",
							  "rn": "l2-[vxlan-3331283]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331283",
										  "rtt": "route-target:as2-nn2:65000:3331283"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331283",
										  "rtt": "route-target:as2-nn2:65000:3331283"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331279",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34046",
							  "rn": "l2-[vxlan-3331279]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331279",
										  "rtt": "route-target:as2-nn2:65000:3331279"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331279",
										  "rtt": "route-target:as2-nn2:65000:3331279"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331278",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34045",
							  "rn": "l2-[vxlan-3331278]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331278",
										  "rtt": "route-target:as2-nn2:65000:3331278"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331278",
										  "rtt": "route-target:as2-nn2:65000:3331278"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331313",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34080",
							  "rn": "l2-[vxlan-3331313]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331313",
										  "rtt": "route-target:as2-nn2:65000:3331313"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331313",
										  "rtt": "route-target:as2-nn2:65000:3331313"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331275",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34042",
							  "rn": "l2-[vxlan-3331275]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331275",
										  "rtt": "route-target:as2-nn2:65000:3331275"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331275",
										  "rtt": "route-target:as2-nn2:65000:3331275"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331272",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34039",
							  "rn": "l2-[vxlan-3331272]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331272",
										  "rtt": "route-target:as2-nn2:65000:3331272"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331272",
										  "rtt": "route-target:as2-nn2:65000:3331272"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331261",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34028",
							  "rn": "l2-[vxlan-3331261]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331261",
										  "rtt": "route-target:as2-nn2:65000:3331261"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331261",
										  "rtt": "route-target:as2-nn2:65000:3331261"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331260",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34027",
							  "rn": "l2-[vxlan-3331260]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331260",
										  "rtt": "route-target:as2-nn2:65000:3331260"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331260",
										  "rtt": "route-target:as2-nn2:65000:3331260"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331458",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34225",
							  "rn": "l2-[vxlan-3331458]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331458",
										  "rtt": "route-target:as2-nn2:65000:3331458"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331458",
										  "rtt": "route-target:as2-nn2:65000:3331458"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331258",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34025",
							  "rn": "l2-[vxlan-3331258]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331258",
										  "rtt": "route-target:as2-nn2:65000:3331258"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331258",
										  "rtt": "route-target:as2-nn2:65000:3331258"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331256",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34023",
							  "rn": "l2-[vxlan-3331256]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331256",
										  "rtt": "route-target:as2-nn2:65000:3331256"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331256",
										  "rtt": "route-target:as2-nn2:65000:3331256"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331255",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34022",
							  "rn": "l2-[vxlan-3331255]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331255",
										  "rtt": "route-target:as2-nn2:65000:3331255"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331255",
										  "rtt": "route-target:as2-nn2:65000:3331255"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333999",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33766",
							  "rn": "l2-[vxlan-333999]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333999",
										  "rtt": "route-target:as2-nn2:65000:333999"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333999",
										  "rtt": "route-target:as2-nn2:65000:333999"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333998",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33765",
							  "rn": "l2-[vxlan-333998]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333998",
										  "rtt": "route-target:as2-nn2:65000:333998"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333998",
										  "rtt": "route-target:as2-nn2:65000:333998"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333997",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33764",
							  "rn": "l2-[vxlan-333997]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333997",
										  "rtt": "route-target:as2-nn2:65000:333997"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333997",
										  "rtt": "route-target:as2-nn2:65000:333997"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331421",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34188",
							  "rn": "l2-[vxlan-3331421]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331421",
										  "rtt": "route-target:as2-nn2:65000:3331421"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331421",
										  "rtt": "route-target:as2-nn2:65000:3331421"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333996",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33763",
							  "rn": "l2-[vxlan-333996]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333996",
										  "rtt": "route-target:as2-nn2:65000:333996"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333996",
										  "rtt": "route-target:as2-nn2:65000:333996"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333993",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33760",
							  "rn": "l2-[vxlan-333993]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333993",
										  "rtt": "route-target:as2-nn2:65000:333993"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333993",
										  "rtt": "route-target:as2-nn2:65000:333993"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331270",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34037",
							  "rn": "l2-[vxlan-3331270]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331270",
										  "rtt": "route-target:as2-nn2:65000:3331270"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331270",
										  "rtt": "route-target:as2-nn2:65000:3331270"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333992",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33759",
							  "rn": "l2-[vxlan-333992]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333992",
										  "rtt": "route-target:as2-nn2:65000:333992"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333992",
										  "rtt": "route-target:as2-nn2:65000:333992"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333991",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33758",
							  "rn": "l2-[vxlan-333991]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333991",
										  "rtt": "route-target:as2-nn2:65000:333991"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333991",
										  "rtt": "route-target:as2-nn2:65000:333991"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331324",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34091",
							  "rn": "l2-[vxlan-3331324]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331324",
										  "rtt": "route-target:as2-nn2:65000:3331324"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331324",
										  "rtt": "route-target:as2-nn2:65000:3331324"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333990",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33757",
							  "rn": "l2-[vxlan-333990]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333990",
										  "rtt": "route-target:as2-nn2:65000:333990"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333990",
										  "rtt": "route-target:as2-nn2:65000:333990"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331370",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34137",
							  "rn": "l2-[vxlan-3331370]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331370",
										  "rtt": "route-target:as2-nn2:65000:3331370"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331370",
										  "rtt": "route-target:as2-nn2:65000:3331370"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333987",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33754",
							  "rn": "l2-[vxlan-333987]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333987",
										  "rtt": "route-target:as2-nn2:65000:333987"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333987",
										  "rtt": "route-target:as2-nn2:65000:333987"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331322",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34089",
							  "rn": "l2-[vxlan-3331322]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331322",
										  "rtt": "route-target:as2-nn2:65000:3331322"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331322",
										  "rtt": "route-target:as2-nn2:65000:3331322"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331289",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34056",
							  "rn": "l2-[vxlan-3331289]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331289",
										  "rtt": "route-target:as2-nn2:65000:3331289"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331289",
										  "rtt": "route-target:as2-nn2:65000:3331289"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333986",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33753",
							  "rn": "l2-[vxlan-333986]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333986",
										  "rtt": "route-target:as2-nn2:65000:333986"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333986",
										  "rtt": "route-target:as2-nn2:65000:333986"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333985",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33752",
							  "rn": "l2-[vxlan-333985]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333985",
										  "rtt": "route-target:as2-nn2:65000:333985"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333985",
										  "rtt": "route-target:as2-nn2:65000:333985"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333984",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33751",
							  "rn": "l2-[vxlan-333984]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333984",
										  "rtt": "route-target:as2-nn2:65000:333984"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333984",
										  "rtt": "route-target:as2-nn2:65000:333984"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333983",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33750",
							  "rn": "l2-[vxlan-333983]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333983",
										  "rtt": "route-target:as2-nn2:65000:333983"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333983",
										  "rtt": "route-target:as2-nn2:65000:333983"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333978",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33745",
							  "rn": "l2-[vxlan-333978]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333978",
										  "rtt": "route-target:as2-nn2:65000:333978"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333978",
										  "rtt": "route-target:as2-nn2:65000:333978"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331362",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34129",
							  "rn": "l2-[vxlan-3331362]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331362",
										  "rtt": "route-target:as2-nn2:65000:3331362"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331362",
										  "rtt": "route-target:as2-nn2:65000:3331362"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333977",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33744",
							  "rn": "l2-[vxlan-333977]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333977",
										  "rtt": "route-target:as2-nn2:65000:333977"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333977",
										  "rtt": "route-target:as2-nn2:65000:333977"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333976",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33743",
							  "rn": "l2-[vxlan-333976]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333976",
										  "rtt": "route-target:as2-nn2:65000:333976"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333976",
										  "rtt": "route-target:as2-nn2:65000:333976"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333975",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33742",
							  "rn": "l2-[vxlan-333975]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333975",
										  "rtt": "route-target:as2-nn2:65000:333975"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333975",
										  "rtt": "route-target:as2-nn2:65000:333975"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331367",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34134",
							  "rn": "l2-[vxlan-3331367]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331367",
										  "rtt": "route-target:as2-nn2:65000:3331367"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331367",
										  "rtt": "route-target:as2-nn2:65000:3331367"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333974",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33741",
							  "rn": "l2-[vxlan-333974]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333974",
										  "rtt": "route-target:as2-nn2:65000:333974"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333974",
										  "rtt": "route-target:as2-nn2:65000:333974"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333973",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33740",
							  "rn": "l2-[vxlan-333973]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333973",
										  "rtt": "route-target:as2-nn2:65000:333973"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333973",
										  "rtt": "route-target:as2-nn2:65000:333973"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331271",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34038",
							  "rn": "l2-[vxlan-3331271]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331271",
										  "rtt": "route-target:as2-nn2:65000:3331271"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331271",
										  "rtt": "route-target:as2-nn2:65000:3331271"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333971",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33738",
							  "rn": "l2-[vxlan-333971]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333971",
										  "rtt": "route-target:as2-nn2:65000:333971"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333971",
										  "rtt": "route-target:as2-nn2:65000:333971"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333970",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33737",
							  "rn": "l2-[vxlan-333970]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333970",
										  "rtt": "route-target:as2-nn2:65000:333970"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333970",
										  "rtt": "route-target:as2-nn2:65000:333970"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333969",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33736",
							  "rn": "l2-[vxlan-333969]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333969",
										  "rtt": "route-target:as2-nn2:65000:333969"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333969",
										  "rtt": "route-target:as2-nn2:65000:333969"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333960",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33727",
							  "rn": "l2-[vxlan-333960]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333960",
										  "rtt": "route-target:as2-nn2:65000:333960"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333960",
										  "rtt": "route-target:as2-nn2:65000:333960"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333959",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33726",
							  "rn": "l2-[vxlan-333959]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333959",
										  "rtt": "route-target:as2-nn2:65000:333959"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333959",
										  "rtt": "route-target:as2-nn2:65000:333959"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333958",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33725",
							  "rn": "l2-[vxlan-333958]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333958",
										  "rtt": "route-target:as2-nn2:65000:333958"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333958",
										  "rtt": "route-target:as2-nn2:65000:333958"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333956",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33723",
							  "rn": "l2-[vxlan-333956]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333956",
										  "rtt": "route-target:as2-nn2:65000:333956"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333956",
										  "rtt": "route-target:as2-nn2:65000:333956"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333966",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33733",
							  "rn": "l2-[vxlan-333966]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333966",
										  "rtt": "route-target:as2-nn2:65000:333966"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333966",
										  "rtt": "route-target:as2-nn2:65000:333966"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333955",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33722",
							  "rn": "l2-[vxlan-333955]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333955",
										  "rtt": "route-target:as2-nn2:65000:333955"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333955",
										  "rtt": "route-target:as2-nn2:65000:333955"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333954",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33721",
							  "rn": "l2-[vxlan-333954]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333954",
										  "rtt": "route-target:as2-nn2:65000:333954"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333954",
										  "rtt": "route-target:as2-nn2:65000:333954"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333953",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33720",
							  "rn": "l2-[vxlan-333953]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333953",
										  "rtt": "route-target:as2-nn2:65000:333953"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333953",
										  "rtt": "route-target:as2-nn2:65000:333953"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333952",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33719",
							  "rn": "l2-[vxlan-333952]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333952",
										  "rtt": "route-target:as2-nn2:65000:333952"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333952",
										  "rtt": "route-target:as2-nn2:65000:333952"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333951",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33718",
							  "rn": "l2-[vxlan-333951]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333951",
										  "rtt": "route-target:as2-nn2:65000:333951"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333951",
										  "rtt": "route-target:as2-nn2:65000:333951"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333948",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33715",
							  "rn": "l2-[vxlan-333948]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333948",
										  "rtt": "route-target:as2-nn2:65000:333948"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333948",
										  "rtt": "route-target:as2-nn2:65000:333948"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331357",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34124",
							  "rn": "l2-[vxlan-3331357]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331357",
										  "rtt": "route-target:as2-nn2:65000:3331357"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331357",
										  "rtt": "route-target:as2-nn2:65000:3331357"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333946",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33713",
							  "rn": "l2-[vxlan-333946]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333946",
										  "rtt": "route-target:as2-nn2:65000:333946"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333946",
										  "rtt": "route-target:as2-nn2:65000:333946"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333945",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33712",
							  "rn": "l2-[vxlan-333945]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333945",
										  "rtt": "route-target:as2-nn2:65000:333945"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333945",
										  "rtt": "route-target:as2-nn2:65000:333945"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333944",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33711",
							  "rn": "l2-[vxlan-333944]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333944",
										  "rtt": "route-target:as2-nn2:65000:333944"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333944",
										  "rtt": "route-target:as2-nn2:65000:333944"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333943",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33710",
							  "rn": "l2-[vxlan-333943]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333943",
										  "rtt": "route-target:as2-nn2:65000:333943"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333943",
										  "rtt": "route-target:as2-nn2:65000:333943"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333941",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33708",
							  "rn": "l2-[vxlan-333941]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333941",
										  "rtt": "route-target:as2-nn2:65000:333941"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333941",
										  "rtt": "route-target:as2-nn2:65000:333941"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333940",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33707",
							  "rn": "l2-[vxlan-333940]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333940",
										  "rtt": "route-target:as2-nn2:65000:333940"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333940",
										  "rtt": "route-target:as2-nn2:65000:333940"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333936",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33703",
							  "rn": "l2-[vxlan-333936]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333936",
										  "rtt": "route-target:as2-nn2:65000:333936"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333936",
										  "rtt": "route-target:as2-nn2:65000:333936"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333935",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33702",
							  "rn": "l2-[vxlan-333935]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333935",
										  "rtt": "route-target:as2-nn2:65000:333935"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333935",
										  "rtt": "route-target:as2-nn2:65000:333935"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333934",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33701",
							  "rn": "l2-[vxlan-333934]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333934",
										  "rtt": "route-target:as2-nn2:65000:333934"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333934",
										  "rtt": "route-target:as2-nn2:65000:333934"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331356",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34123",
							  "rn": "l2-[vxlan-3331356]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331356",
										  "rtt": "route-target:as2-nn2:65000:3331356"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331356",
										  "rtt": "route-target:as2-nn2:65000:3331356"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333933",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33700",
							  "rn": "l2-[vxlan-333933]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333933",
										  "rtt": "route-target:as2-nn2:65000:333933"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333933",
										  "rtt": "route-target:as2-nn2:65000:333933"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331281",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34048",
							  "rn": "l2-[vxlan-3331281]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331281",
										  "rtt": "route-target:as2-nn2:65000:3331281"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331281",
										  "rtt": "route-target:as2-nn2:65000:3331281"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333932",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33699",
							  "rn": "l2-[vxlan-333932]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333932",
										  "rtt": "route-target:as2-nn2:65000:333932"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333932",
										  "rtt": "route-target:as2-nn2:65000:333932"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333929",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33696",
							  "rn": "l2-[vxlan-333929]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333929",
										  "rtt": "route-target:as2-nn2:65000:333929"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333929",
										  "rtt": "route-target:as2-nn2:65000:333929"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333927",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33694",
							  "rn": "l2-[vxlan-333927]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333927",
										  "rtt": "route-target:as2-nn2:65000:333927"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333927",
										  "rtt": "route-target:as2-nn2:65000:333927"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333926",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33693",
							  "rn": "l2-[vxlan-333926]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333926",
										  "rtt": "route-target:as2-nn2:65000:333926"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333926",
										  "rtt": "route-target:as2-nn2:65000:333926"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333925",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33692",
							  "rn": "l2-[vxlan-333925]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333925",
										  "rtt": "route-target:as2-nn2:65000:333925"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333925",
										  "rtt": "route-target:as2-nn2:65000:333925"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333924",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33691",
							  "rn": "l2-[vxlan-333924]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333924",
										  "rtt": "route-target:as2-nn2:65000:333924"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333924",
										  "rtt": "route-target:as2-nn2:65000:333924"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333922",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33689",
							  "rn": "l2-[vxlan-333922]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333922",
										  "rtt": "route-target:as2-nn2:65000:333922"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333922",
										  "rtt": "route-target:as2-nn2:65000:333922"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333920",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33687",
							  "rn": "l2-[vxlan-333920]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333920",
										  "rtt": "route-target:as2-nn2:65000:333920"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333920",
										  "rtt": "route-target:as2-nn2:65000:333920"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333919",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33686",
							  "rn": "l2-[vxlan-333919]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333919",
										  "rtt": "route-target:as2-nn2:65000:333919"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333919",
										  "rtt": "route-target:as2-nn2:65000:333919"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333982",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33749",
							  "rn": "l2-[vxlan-333982]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333982",
										  "rtt": "route-target:as2-nn2:65000:333982"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333982",
										  "rtt": "route-target:as2-nn2:65000:333982"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333918",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33685",
							  "rn": "l2-[vxlan-333918]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333918",
										  "rtt": "route-target:as2-nn2:65000:333918"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333918",
										  "rtt": "route-target:as2-nn2:65000:333918"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333916",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33683",
							  "rn": "l2-[vxlan-333916]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333916",
										  "rtt": "route-target:as2-nn2:65000:333916"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333916",
										  "rtt": "route-target:as2-nn2:65000:333916"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331274",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34041",
							  "rn": "l2-[vxlan-3331274]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331274",
										  "rtt": "route-target:as2-nn2:65000:3331274"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331274",
										  "rtt": "route-target:as2-nn2:65000:3331274"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331268",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34035",
							  "rn": "l2-[vxlan-3331268]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331268",
										  "rtt": "route-target:as2-nn2:65000:3331268"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331268",
										  "rtt": "route-target:as2-nn2:65000:3331268"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333914",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33681",
							  "rn": "l2-[vxlan-333914]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333914",
										  "rtt": "route-target:as2-nn2:65000:333914"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333914",
										  "rtt": "route-target:as2-nn2:65000:333914"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333913",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33680",
							  "rn": "l2-[vxlan-333913]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333913",
										  "rtt": "route-target:as2-nn2:65000:333913"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333913",
										  "rtt": "route-target:as2-nn2:65000:333913"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333912",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33679",
							  "rn": "l2-[vxlan-333912]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333912",
										  "rtt": "route-target:as2-nn2:65000:333912"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333912",
										  "rtt": "route-target:as2-nn2:65000:333912"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333911",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33678",
							  "rn": "l2-[vxlan-333911]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333911",
										  "rtt": "route-target:as2-nn2:65000:333911"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333911",
										  "rtt": "route-target:as2-nn2:65000:333911"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333910",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33677",
							  "rn": "l2-[vxlan-333910]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333910",
										  "rtt": "route-target:as2-nn2:65000:333910"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333910",
										  "rtt": "route-target:as2-nn2:65000:333910"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331409",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34176",
							  "rn": "l2-[vxlan-3331409]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331409",
										  "rtt": "route-target:as2-nn2:65000:3331409"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331409",
										  "rtt": "route-target:as2-nn2:65000:3331409"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333906",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33673",
							  "rn": "l2-[vxlan-333906]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333906",
										  "rtt": "route-target:as2-nn2:65000:333906"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333906",
										  "rtt": "route-target:as2-nn2:65000:333906"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333903",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33670",
							  "rn": "l2-[vxlan-333903]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333903",
										  "rtt": "route-target:as2-nn2:65000:333903"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333903",
										  "rtt": "route-target:as2-nn2:65000:333903"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333902",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33669",
							  "rn": "l2-[vxlan-333902]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333902",
										  "rtt": "route-target:as2-nn2:65000:333902"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333902",
										  "rtt": "route-target:as2-nn2:65000:333902"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333901",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33668",
							  "rn": "l2-[vxlan-333901]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333901",
										  "rtt": "route-target:as2-nn2:65000:333901"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333901",
										  "rtt": "route-target:as2-nn2:65000:333901"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333900",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33667",
							  "rn": "l2-[vxlan-333900]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333900",
										  "rtt": "route-target:as2-nn2:65000:333900"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333900",
										  "rtt": "route-target:as2-nn2:65000:333900"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333899",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33666",
							  "rn": "l2-[vxlan-333899]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333899",
										  "rtt": "route-target:as2-nn2:65000:333899"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333899",
										  "rtt": "route-target:as2-nn2:65000:333899"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333897",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33664",
							  "rn": "l2-[vxlan-333897]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333897",
										  "rtt": "route-target:as2-nn2:65000:333897"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333897",
										  "rtt": "route-target:as2-nn2:65000:333897"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331262",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34029",
							  "rn": "l2-[vxlan-3331262]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331262",
										  "rtt": "route-target:as2-nn2:65000:3331262"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331262",
										  "rtt": "route-target:as2-nn2:65000:3331262"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333895",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33662",
							  "rn": "l2-[vxlan-333895]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333895",
										  "rtt": "route-target:as2-nn2:65000:333895"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333895",
										  "rtt": "route-target:as2-nn2:65000:333895"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333893",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33660",
							  "rn": "l2-[vxlan-333893]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333893",
										  "rtt": "route-target:as2-nn2:65000:333893"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333893",
										  "rtt": "route-target:as2-nn2:65000:333893"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331450",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34217",
							  "rn": "l2-[vxlan-3331450]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331450",
										  "rtt": "route-target:as2-nn2:65000:3331450"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331450",
										  "rtt": "route-target:as2-nn2:65000:3331450"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333892",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33659",
							  "rn": "l2-[vxlan-333892]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333892",
										  "rtt": "route-target:as2-nn2:65000:333892"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333892",
										  "rtt": "route-target:as2-nn2:65000:333892"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333937",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33704",
							  "rn": "l2-[vxlan-333937]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333937",
										  "rtt": "route-target:as2-nn2:65000:333937"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333937",
										  "rtt": "route-target:as2-nn2:65000:333937"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333891",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33658",
							  "rn": "l2-[vxlan-333891]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333891",
										  "rtt": "route-target:as2-nn2:65000:333891"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333891",
										  "rtt": "route-target:as2-nn2:65000:333891"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333889",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33656",
							  "rn": "l2-[vxlan-333889]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333889",
										  "rtt": "route-target:as2-nn2:65000:333889"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333889",
										  "rtt": "route-target:as2-nn2:65000:333889"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333888",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33655",
							  "rn": "l2-[vxlan-333888]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333888",
										  "rtt": "route-target:as2-nn2:65000:333888"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333888",
										  "rtt": "route-target:as2-nn2:65000:333888"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333886",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33653",
							  "rn": "l2-[vxlan-333886]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333886",
										  "rtt": "route-target:as2-nn2:65000:333886"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333886",
										  "rtt": "route-target:as2-nn2:65000:333886"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333962",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33729",
							  "rn": "l2-[vxlan-333962]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333962",
										  "rtt": "route-target:as2-nn2:65000:333962"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333962",
										  "rtt": "route-target:as2-nn2:65000:333962"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333885",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33652",
							  "rn": "l2-[vxlan-333885]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333885",
										  "rtt": "route-target:as2-nn2:65000:333885"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333885",
										  "rtt": "route-target:as2-nn2:65000:333885"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331291",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34058",
							  "rn": "l2-[vxlan-3331291]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331291",
										  "rtt": "route-target:as2-nn2:65000:3331291"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331291",
										  "rtt": "route-target:as2-nn2:65000:3331291"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333422",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33189",
							  "rn": "l2-[vxlan-333422]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333422",
										  "rtt": "route-target:as2-nn2:65000:333422"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333422",
										  "rtt": "route-target:as2-nn2:65000:333422"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333344",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33111",
							  "rn": "l2-[vxlan-333344]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333344",
										  "rtt": "route-target:as2-nn2:65000:333344"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333344",
										  "rtt": "route-target:as2-nn2:65000:333344"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333412",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33179",
							  "rn": "l2-[vxlan-333412]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333412",
										  "rtt": "route-target:as2-nn2:65000:333412"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333412",
										  "rtt": "route-target:as2-nn2:65000:333412"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333406",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33173",
							  "rn": "l2-[vxlan-333406]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333406",
										  "rtt": "route-target:as2-nn2:65000:333406"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333406",
										  "rtt": "route-target:as2-nn2:65000:333406"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331490",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34257",
							  "rn": "l2-[vxlan-3331490]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331490",
										  "rtt": "route-target:as2-nn2:65000:3331490"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331490",
										  "rtt": "route-target:as2-nn2:65000:3331490"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333405",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33172",
							  "rn": "l2-[vxlan-333405]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333405",
										  "rtt": "route-target:as2-nn2:65000:333405"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333405",
										  "rtt": "route-target:as2-nn2:65000:333405"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333373",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33140",
							  "rn": "l2-[vxlan-333373]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333373",
										  "rtt": "route-target:as2-nn2:65000:333373"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333373",
										  "rtt": "route-target:as2-nn2:65000:333373"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333404",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33171",
							  "rn": "l2-[vxlan-333404]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333404",
										  "rtt": "route-target:as2-nn2:65000:333404"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333404",
										  "rtt": "route-target:as2-nn2:65000:333404"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331433",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34200",
							  "rn": "l2-[vxlan-3331433]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331433",
										  "rtt": "route-target:as2-nn2:65000:3331433"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331433",
										  "rtt": "route-target:as2-nn2:65000:3331433"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333880",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33647",
							  "rn": "l2-[vxlan-333880]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333880",
										  "rtt": "route-target:as2-nn2:65000:333880"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333880",
										  "rtt": "route-target:as2-nn2:65000:333880"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331432",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34199",
							  "rn": "l2-[vxlan-3331432]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331432",
										  "rtt": "route-target:as2-nn2:65000:3331432"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331432",
										  "rtt": "route-target:as2-nn2:65000:3331432"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333416",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33183",
							  "rn": "l2-[vxlan-333416]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333416",
										  "rtt": "route-target:as2-nn2:65000:333416"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333416",
										  "rtt": "route-target:as2-nn2:65000:333416"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333369",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33136",
							  "rn": "l2-[vxlan-333369]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333369",
										  "rtt": "route-target:as2-nn2:65000:333369"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333369",
										  "rtt": "route-target:as2-nn2:65000:333369"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333831",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33598",
							  "rn": "l2-[vxlan-333831]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333831",
										  "rtt": "route-target:as2-nn2:65000:333831"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333831",
										  "rtt": "route-target:as2-nn2:65000:333831"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333393",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33160",
							  "rn": "l2-[vxlan-333393]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333393",
										  "rtt": "route-target:as2-nn2:65000:333393"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333393",
										  "rtt": "route-target:as2-nn2:65000:333393"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333480",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33247",
							  "rn": "l2-[vxlan-333480]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333480",
										  "rtt": "route-target:as2-nn2:65000:333480"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333480",
										  "rtt": "route-target:as2-nn2:65000:333480"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331264",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34031",
							  "rn": "l2-[vxlan-3331264]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331264",
										  "rtt": "route-target:as2-nn2:65000:3331264"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331264",
										  "rtt": "route-target:as2-nn2:65000:3331264"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333392",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33159",
							  "rn": "l2-[vxlan-333392]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333392",
										  "rtt": "route-target:as2-nn2:65000:333392"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333392",
										  "rtt": "route-target:as2-nn2:65000:333392"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333844",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33611",
							  "rn": "l2-[vxlan-333844]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333844",
										  "rtt": "route-target:as2-nn2:65000:333844"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333844",
										  "rtt": "route-target:as2-nn2:65000:333844"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333386",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33153",
							  "rn": "l2-[vxlan-333386]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333386",
										  "rtt": "route-target:as2-nn2:65000:333386"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333386",
										  "rtt": "route-target:as2-nn2:65000:333386"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333807",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33574",
							  "rn": "l2-[vxlan-333807]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333807",
										  "rtt": "route-target:as2-nn2:65000:333807"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333807",
										  "rtt": "route-target:as2-nn2:65000:333807"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331353",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34120",
							  "rn": "l2-[vxlan-3331353]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331353",
										  "rtt": "route-target:as2-nn2:65000:3331353"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331353",
										  "rtt": "route-target:as2-nn2:65000:3331353"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333863",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33630",
							  "rn": "l2-[vxlan-333863]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333863",
										  "rtt": "route-target:as2-nn2:65000:333863"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333863",
										  "rtt": "route-target:as2-nn2:65000:333863"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333376",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33143",
							  "rn": "l2-[vxlan-333376]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333376",
										  "rtt": "route-target:as2-nn2:65000:333376"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333376",
										  "rtt": "route-target:as2-nn2:65000:333376"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333915",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33682",
							  "rn": "l2-[vxlan-333915]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333915",
										  "rtt": "route-target:as2-nn2:65000:333915"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333915",
										  "rtt": "route-target:as2-nn2:65000:333915"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-1004",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:32771",
							  "rn": "l2-[vxlan-1004]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:1004",
										  "rtt": "route-target:as2-nn2:1:1004"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:1004",
										  "rtt": "route-target:as2-nn2:1:1004"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331429",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34196",
							  "rn": "l2-[vxlan-3331429]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331429",
										  "rtt": "route-target:as2-nn2:65000:3331429"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331429",
										  "rtt": "route-target:as2-nn2:65000:3331429"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333332",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33099",
							  "rn": "l2-[vxlan-333332]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333332",
										  "rtt": "route-target:as2-nn2:65000:333332"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333332",
										  "rtt": "route-target:as2-nn2:65000:333332"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331451",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34218",
							  "rn": "l2-[vxlan-3331451]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331451",
										  "rtt": "route-target:as2-nn2:65000:3331451"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331451",
										  "rtt": "route-target:as2-nn2:65000:3331451"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333267",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33034",
							  "rn": "l2-[vxlan-333267]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333267",
										  "rtt": "route-target:as2-nn2:65000:333267"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333267",
										  "rtt": "route-target:as2-nn2:65000:333267"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333413",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33180",
							  "rn": "l2-[vxlan-333413]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333413",
										  "rtt": "route-target:as2-nn2:65000:333413"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333413",
										  "rtt": "route-target:as2-nn2:65000:333413"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333370",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33137",
							  "rn": "l2-[vxlan-333370]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333370",
										  "rtt": "route-target:as2-nn2:65000:333370"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333370",
										  "rtt": "route-target:as2-nn2:65000:333370"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333284",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33051",
							  "rn": "l2-[vxlan-333284]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333284",
										  "rtt": "route-target:as2-nn2:65000:333284"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333284",
										  "rtt": "route-target:as2-nn2:65000:333284"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331416",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34183",
							  "rn": "l2-[vxlan-3331416]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331416",
										  "rtt": "route-target:as2-nn2:65000:3331416"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331416",
										  "rtt": "route-target:as2-nn2:65000:3331416"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333478",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33245",
							  "rn": "l2-[vxlan-333478]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333478",
										  "rtt": "route-target:as2-nn2:65000:333478"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333478",
										  "rtt": "route-target:as2-nn2:65000:333478"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333336",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33103",
							  "rn": "l2-[vxlan-333336]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333336",
										  "rtt": "route-target:as2-nn2:65000:333336"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333336",
										  "rtt": "route-target:as2-nn2:65000:333336"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333756",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33523",
							  "rn": "l2-[vxlan-333756]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333756",
										  "rtt": "route-target:as2-nn2:65000:333756"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333756",
										  "rtt": "route-target:as2-nn2:65000:333756"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333850",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33617",
							  "rn": "l2-[vxlan-333850]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333850",
										  "rtt": "route-target:as2-nn2:65000:333850"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333850",
										  "rtt": "route-target:as2-nn2:65000:333850"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333367",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33134",
							  "rn": "l2-[vxlan-333367]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333367",
										  "rtt": "route-target:as2-nn2:65000:333367"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333367",
										  "rtt": "route-target:as2-nn2:65000:333367"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333431",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33198",
							  "rn": "l2-[vxlan-333431]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333431",
										  "rtt": "route-target:as2-nn2:65000:333431"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333431",
										  "rtt": "route-target:as2-nn2:65000:333431"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331468",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34235",
							  "rn": "l2-[vxlan-3331468]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331468",
										  "rtt": "route-target:as2-nn2:65000:3331468"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331468",
										  "rtt": "route-target:as2-nn2:65000:3331468"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331415",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34182",
							  "rn": "l2-[vxlan-3331415]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331415",
										  "rtt": "route-target:as2-nn2:65000:3331415"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331415",
										  "rtt": "route-target:as2-nn2:65000:3331415"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333942",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33709",
							  "rn": "l2-[vxlan-333942]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333942",
										  "rtt": "route-target:as2-nn2:65000:333942"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333942",
										  "rtt": "route-target:as2-nn2:65000:333942"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333362",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33129",
							  "rn": "l2-[vxlan-333362]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333362",
										  "rtt": "route-target:as2-nn2:65000:333362"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333362",
										  "rtt": "route-target:as2-nn2:65000:333362"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333360",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33127",
							  "rn": "l2-[vxlan-333360]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333360",
										  "rtt": "route-target:as2-nn2:65000:333360"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333360",
										  "rtt": "route-target:as2-nn2:65000:333360"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-2010101",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:32868",
							  "rn": "l2-[vxlan-2010101]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:2010101",
										  "rtt": "route-target:as2-nn2:1:2010101"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:2010101",
										  "rtt": "route-target:as2-nn2:1:2010101"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333905",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33672",
							  "rn": "l2-[vxlan-333905]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333905",
										  "rtt": "route-target:as2-nn2:65000:333905"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333905",
										  "rtt": "route-target:as2-nn2:65000:333905"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333776",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33543",
							  "rn": "l2-[vxlan-333776]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333776",
										  "rtt": "route-target:as2-nn2:65000:333776"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333776",
										  "rtt": "route-target:as2-nn2:65000:333776"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333353",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33120",
							  "rn": "l2-[vxlan-333353]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333353",
										  "rtt": "route-target:as2-nn2:65000:333353"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333353",
										  "rtt": "route-target:as2-nn2:65000:333353"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333352",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33119",
							  "rn": "l2-[vxlan-333352]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333352",
										  "rtt": "route-target:as2-nn2:65000:333352"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333352",
										  "rtt": "route-target:as2-nn2:65000:333352"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331296",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34063",
							  "rn": "l2-[vxlan-3331296]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331296",
										  "rtt": "route-target:as2-nn2:65000:3331296"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331296",
										  "rtt": "route-target:as2-nn2:65000:3331296"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333499",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33266",
							  "rn": "l2-[vxlan-333499]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333499",
										  "rtt": "route-target:as2-nn2:65000:333499"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333499",
										  "rtt": "route-target:as2-nn2:65000:333499"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333830",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33597",
							  "rn": "l2-[vxlan-333830]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333830",
										  "rtt": "route-target:as2-nn2:65000:333830"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333830",
										  "rtt": "route-target:as2-nn2:65000:333830"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333266",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33033",
							  "rn": "l2-[vxlan-333266]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333266",
										  "rtt": "route-target:as2-nn2:65000:333266"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333266",
										  "rtt": "route-target:as2-nn2:65000:333266"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333768",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33535",
							  "rn": "l2-[vxlan-333768]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333768",
										  "rtt": "route-target:as2-nn2:65000:333768"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333768",
										  "rtt": "route-target:as2-nn2:65000:333768"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333846",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33613",
							  "rn": "l2-[vxlan-333846]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333846",
										  "rtt": "route-target:as2-nn2:65000:333846"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333846",
										  "rtt": "route-target:as2-nn2:65000:333846"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333346",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33113",
							  "rn": "l2-[vxlan-333346]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333346",
										  "rtt": "route-target:as2-nn2:65000:333346"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333346",
										  "rtt": "route-target:as2-nn2:65000:333346"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333950",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33717",
							  "rn": "l2-[vxlan-333950]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333950",
										  "rtt": "route-target:as2-nn2:65000:333950"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333950",
										  "rtt": "route-target:as2-nn2:65000:333950"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333270",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33037",
							  "rn": "l2-[vxlan-333270]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333270",
										  "rtt": "route-target:as2-nn2:65000:333270"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333270",
										  "rtt": "route-target:as2-nn2:65000:333270"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331385",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34152",
							  "rn": "l2-[vxlan-3331385]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331385",
										  "rtt": "route-target:as2-nn2:65000:3331385"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331385",
										  "rtt": "route-target:as2-nn2:65000:3331385"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333492",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33259",
							  "rn": "l2-[vxlan-333492]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333492",
										  "rtt": "route-target:as2-nn2:65000:333492"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333492",
										  "rtt": "route-target:as2-nn2:65000:333492"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333397",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33164",
							  "rn": "l2-[vxlan-333397]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333397",
										  "rtt": "route-target:as2-nn2:65000:333397"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333397",
										  "rtt": "route-target:as2-nn2:65000:333397"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331460",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34227",
							  "rn": "l2-[vxlan-3331460]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331460",
										  "rtt": "route-target:as2-nn2:65000:3331460"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331460",
										  "rtt": "route-target:as2-nn2:65000:3331460"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333832",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33599",
							  "rn": "l2-[vxlan-333832]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333832",
										  "rtt": "route-target:as2-nn2:65000:333832"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333832",
										  "rtt": "route-target:as2-nn2:65000:333832"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333298",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33065",
							  "rn": "l2-[vxlan-333298]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333298",
										  "rtt": "route-target:as2-nn2:65000:333298"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333298",
										  "rtt": "route-target:as2-nn2:65000:333298"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333340",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33107",
							  "rn": "l2-[vxlan-333340]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333340",
										  "rtt": "route-target:as2-nn2:65000:333340"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333340",
										  "rtt": "route-target:as2-nn2:65000:333340"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333339",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33106",
							  "rn": "l2-[vxlan-333339]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333339",
										  "rtt": "route-target:as2-nn2:65000:333339"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333339",
										  "rtt": "route-target:as2-nn2:65000:333339"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333338",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33105",
							  "rn": "l2-[vxlan-333338]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333338",
										  "rtt": "route-target:as2-nn2:65000:333338"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333338",
										  "rtt": "route-target:as2-nn2:65000:333338"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333767",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33534",
							  "rn": "l2-[vxlan-333767]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333767",
										  "rtt": "route-target:as2-nn2:65000:333767"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333767",
										  "rtt": "route-target:as2-nn2:65000:333767"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333421",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33188",
							  "rn": "l2-[vxlan-333421]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333421",
										  "rtt": "route-target:as2-nn2:65000:333421"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333421",
										  "rtt": "route-target:as2-nn2:65000:333421"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333472",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33239",
							  "rn": "l2-[vxlan-333472]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333472",
										  "rtt": "route-target:as2-nn2:65000:333472"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333472",
										  "rtt": "route-target:as2-nn2:65000:333472"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333972",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33739",
							  "rn": "l2-[vxlan-333972]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333972",
										  "rtt": "route-target:as2-nn2:65000:333972"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333972",
										  "rtt": "route-target:as2-nn2:65000:333972"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333400",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33167",
							  "rn": "l2-[vxlan-333400]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333400",
										  "rtt": "route-target:as2-nn2:65000:333400"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333400",
										  "rtt": "route-target:as2-nn2:65000:333400"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333337",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33104",
							  "rn": "l2-[vxlan-333337]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333337",
										  "rtt": "route-target:as2-nn2:65000:333337"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333337",
										  "rtt": "route-target:as2-nn2:65000:333337"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-2010100",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:32867",
							  "rn": "l2-[vxlan-2010100]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:2010100",
										  "rtt": "route-target:as2-nn2:1:2010100"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:2010100",
										  "rtt": "route-target:as2-nn2:1:2010100"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333368",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33135",
							  "rn": "l2-[vxlan-333368]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333368",
										  "rtt": "route-target:as2-nn2:65000:333368"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333368",
										  "rtt": "route-target:as2-nn2:65000:333368"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333461",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33228",
							  "rn": "l2-[vxlan-333461]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333461",
										  "rtt": "route-target:as2-nn2:65000:333461"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333461",
										  "rtt": "route-target:as2-nn2:65000:333461"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333357",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33124",
							  "rn": "l2-[vxlan-333357]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333357",
										  "rtt": "route-target:as2-nn2:65000:333357"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333357",
										  "rtt": "route-target:as2-nn2:65000:333357"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333852",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33619",
							  "rn": "l2-[vxlan-333852]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333852",
										  "rtt": "route-target:as2-nn2:65000:333852"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333852",
										  "rtt": "route-target:as2-nn2:65000:333852"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333855",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33622",
							  "rn": "l2-[vxlan-333855]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333855",
										  "rtt": "route-target:as2-nn2:65000:333855"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333855",
										  "rtt": "route-target:as2-nn2:65000:333855"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333331",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33098",
							  "rn": "l2-[vxlan-333331]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333331",
										  "rtt": "route-target:as2-nn2:65000:333331"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333331",
										  "rtt": "route-target:as2-nn2:65000:333331"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331497",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34264",
							  "rn": "l2-[vxlan-3331497]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331497",
										  "rtt": "route-target:as2-nn2:65000:3331497"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331497",
										  "rtt": "route-target:as2-nn2:65000:3331497"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333329",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33096",
							  "rn": "l2-[vxlan-333329]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333329",
										  "rtt": "route-target:as2-nn2:65000:333329"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333329",
										  "rtt": "route-target:as2-nn2:65000:333329"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331282",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34049",
							  "rn": "l2-[vxlan-3331282]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331282",
										  "rtt": "route-target:as2-nn2:65000:3331282"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331282",
										  "rtt": "route-target:as2-nn2:65000:3331282"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333301",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33068",
							  "rn": "l2-[vxlan-333301]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333301",
										  "rtt": "route-target:as2-nn2:65000:333301"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333301",
										  "rtt": "route-target:as2-nn2:65000:333301"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333781",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33548",
							  "rn": "l2-[vxlan-333781]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333781",
										  "rtt": "route-target:as2-nn2:65000:333781"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333781",
										  "rtt": "route-target:as2-nn2:65000:333781"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331269",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34036",
							  "rn": "l2-[vxlan-3331269]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331269",
										  "rtt": "route-target:as2-nn2:65000:3331269"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331269",
										  "rtt": "route-target:as2-nn2:65000:3331269"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333394",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33161",
							  "rn": "l2-[vxlan-333394]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333394",
										  "rtt": "route-target:as2-nn2:65000:333394"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333394",
										  "rtt": "route-target:as2-nn2:65000:333394"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333269",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33036",
							  "rn": "l2-[vxlan-333269]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333269",
										  "rtt": "route-target:as2-nn2:65000:333269"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333269",
										  "rtt": "route-target:as2-nn2:65000:333269"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333364",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33131",
							  "rn": "l2-[vxlan-333364]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333364",
										  "rtt": "route-target:as2-nn2:65000:333364"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333364",
										  "rtt": "route-target:as2-nn2:65000:333364"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333383",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33150",
							  "rn": "l2-[vxlan-333383]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333383",
										  "rtt": "route-target:as2-nn2:65000:333383"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333383",
										  "rtt": "route-target:as2-nn2:65000:333383"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333271",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33038",
							  "rn": "l2-[vxlan-333271]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333271",
										  "rtt": "route-target:as2-nn2:65000:333271"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333271",
										  "rtt": "route-target:as2-nn2:65000:333271"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333291",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33058",
							  "rn": "l2-[vxlan-333291]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333291",
										  "rtt": "route-target:as2-nn2:65000:333291"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333291",
										  "rtt": "route-target:as2-nn2:65000:333291"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333965",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33732",
							  "rn": "l2-[vxlan-333965]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333965",
										  "rtt": "route-target:as2-nn2:65000:333965"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333965",
										  "rtt": "route-target:as2-nn2:65000:333965"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333457",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33224",
							  "rn": "l2-[vxlan-333457]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333457",
										  "rtt": "route-target:as2-nn2:65000:333457"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333457",
										  "rtt": "route-target:as2-nn2:65000:333457"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331441",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34208",
							  "rn": "l2-[vxlan-3331441]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331441",
										  "rtt": "route-target:as2-nn2:65000:3331441"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331441",
										  "rtt": "route-target:as2-nn2:65000:3331441"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333415",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33182",
							  "rn": "l2-[vxlan-333415]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333415",
										  "rtt": "route-target:as2-nn2:65000:333415"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333415",
										  "rtt": "route-target:as2-nn2:65000:333415"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333252",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33019",
							  "rn": "l2-[vxlan-333252]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333252",
										  "rtt": "route-target:as2-nn2:65000:333252"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333252",
										  "rtt": "route-target:as2-nn2:65000:333252"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333968",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33735",
							  "rn": "l2-[vxlan-333968]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333968",
										  "rtt": "route-target:as2-nn2:65000:333968"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333968",
										  "rtt": "route-target:as2-nn2:65000:333968"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333326",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33093",
							  "rn": "l2-[vxlan-333326]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333326",
										  "rtt": "route-target:as2-nn2:65000:333326"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333326",
										  "rtt": "route-target:as2-nn2:65000:333326"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331395",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34162",
							  "rn": "l2-[vxlan-3331395]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331395",
										  "rtt": "route-target:as2-nn2:65000:3331395"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331395",
										  "rtt": "route-target:as2-nn2:65000:3331395"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333486",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33253",
							  "rn": "l2-[vxlan-333486]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333486",
										  "rtt": "route-target:as2-nn2:65000:333486"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333486",
										  "rtt": "route-target:as2-nn2:65000:333486"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333322",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33089",
							  "rn": "l2-[vxlan-333322]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333322",
										  "rtt": "route-target:as2-nn2:65000:333322"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333322",
										  "rtt": "route-target:as2-nn2:65000:333322"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333883",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33650",
							  "rn": "l2-[vxlan-333883]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333883",
										  "rtt": "route-target:as2-nn2:65000:333883"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333883",
										  "rtt": "route-target:as2-nn2:65000:333883"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333752",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33519",
							  "rn": "l2-[vxlan-333752]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333752",
										  "rtt": "route-target:as2-nn2:65000:333752"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333752",
										  "rtt": "route-target:as2-nn2:65000:333752"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333290",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33057",
							  "rn": "l2-[vxlan-333290]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333290",
										  "rtt": "route-target:as2-nn2:65000:333290"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333290",
										  "rtt": "route-target:as2-nn2:65000:333290"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331449",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34216",
							  "rn": "l2-[vxlan-3331449]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331449",
										  "rtt": "route-target:as2-nn2:65000:3331449"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331449",
										  "rtt": "route-target:as2-nn2:65000:3331449"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333770",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33537",
							  "rn": "l2-[vxlan-333770]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333770",
										  "rtt": "route-target:as2-nn2:65000:333770"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333770",
										  "rtt": "route-target:as2-nn2:65000:333770"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333358",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33125",
							  "rn": "l2-[vxlan-333358]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333358",
										  "rtt": "route-target:as2-nn2:65000:333358"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333358",
										  "rtt": "route-target:as2-nn2:65000:333358"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333305",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33072",
							  "rn": "l2-[vxlan-333305]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333305",
										  "rtt": "route-target:as2-nn2:65000:333305"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333305",
										  "rtt": "route-target:as2-nn2:65000:333305"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333904",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33671",
							  "rn": "l2-[vxlan-333904]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333904",
										  "rtt": "route-target:as2-nn2:65000:333904"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333904",
										  "rtt": "route-target:as2-nn2:65000:333904"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333444",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33211",
							  "rn": "l2-[vxlan-333444]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333444",
										  "rtt": "route-target:as2-nn2:65000:333444"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333444",
										  "rtt": "route-target:as2-nn2:65000:333444"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333264",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33031",
							  "rn": "l2-[vxlan-333264]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333264",
										  "rtt": "route-target:as2-nn2:65000:333264"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333264",
										  "rtt": "route-target:as2-nn2:65000:333264"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331379",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34146",
							  "rn": "l2-[vxlan-3331379]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331379",
										  "rtt": "route-target:as2-nn2:65000:3331379"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331379",
										  "rtt": "route-target:as2-nn2:65000:3331379"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-2012006",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34773",
							  "rn": "l2-[vxlan-2012006]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:2012006",
										  "rtt": "route-target:as2-nn2:1:2012006"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:2012006",
										  "rtt": "route-target:as2-nn2:1:2012006"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333372",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33139",
							  "rn": "l2-[vxlan-333372]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333372",
										  "rtt": "route-target:as2-nn2:65000:333372"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333372",
										  "rtt": "route-target:as2-nn2:65000:333372"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333275",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33042",
							  "rn": "l2-[vxlan-333275]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333275",
										  "rtt": "route-target:as2-nn2:65000:333275"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333275",
										  "rtt": "route-target:as2-nn2:65000:333275"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331277",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34044",
							  "rn": "l2-[vxlan-3331277]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331277",
										  "rtt": "route-target:as2-nn2:65000:3331277"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331277",
										  "rtt": "route-target:as2-nn2:65000:3331277"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333452",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33219",
							  "rn": "l2-[vxlan-333452]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333452",
										  "rtt": "route-target:as2-nn2:65000:333452"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333452",
										  "rtt": "route-target:as2-nn2:65000:333452"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333869",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33636",
							  "rn": "l2-[vxlan-333869]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333869",
										  "rtt": "route-target:as2-nn2:65000:333869"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333869",
										  "rtt": "route-target:as2-nn2:65000:333869"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331273",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34040",
							  "rn": "l2-[vxlan-3331273]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331273",
										  "rtt": "route-target:as2-nn2:65000:3331273"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331273",
										  "rtt": "route-target:as2-nn2:65000:3331273"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333411",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33178",
							  "rn": "l2-[vxlan-333411]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333411",
										  "rtt": "route-target:as2-nn2:65000:333411"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333411",
										  "rtt": "route-target:as2-nn2:65000:333411"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333256",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33023",
							  "rn": "l2-[vxlan-333256]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333256",
										  "rtt": "route-target:as2-nn2:65000:333256"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333256",
										  "rtt": "route-target:as2-nn2:65000:333256"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331430",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34197",
							  "rn": "l2-[vxlan-3331430]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331430",
										  "rtt": "route-target:as2-nn2:65000:3331430"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331430",
										  "rtt": "route-target:as2-nn2:65000:3331430"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331405",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34172",
							  "rn": "l2-[vxlan-3331405]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331405",
										  "rtt": "route-target:as2-nn2:65000:3331405"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331405",
										  "rtt": "route-target:as2-nn2:65000:3331405"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331361",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34128",
							  "rn": "l2-[vxlan-3331361]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331361",
										  "rtt": "route-target:as2-nn2:65000:3331361"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331361",
										  "rtt": "route-target:as2-nn2:65000:3331361"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333448",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33215",
							  "rn": "l2-[vxlan-333448]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333448",
										  "rtt": "route-target:as2-nn2:65000:333448"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333448",
										  "rtt": "route-target:as2-nn2:65000:333448"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331467",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34234",
							  "rn": "l2-[vxlan-3331467]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331467",
										  "rtt": "route-target:as2-nn2:65000:3331467"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331467",
										  "rtt": "route-target:as2-nn2:65000:3331467"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331267",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34034",
							  "rn": "l2-[vxlan-3331267]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331267",
										  "rtt": "route-target:as2-nn2:65000:3331267"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331267",
										  "rtt": "route-target:as2-nn2:65000:3331267"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333365",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33132",
							  "rn": "l2-[vxlan-333365]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333365",
										  "rtt": "route-target:as2-nn2:65000:333365"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333365",
										  "rtt": "route-target:as2-nn2:65000:333365"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333251",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33018",
							  "rn": "l2-[vxlan-333251]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333251",
										  "rtt": "route-target:as2-nn2:65000:333251"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333251",
										  "rtt": "route-target:as2-nn2:65000:333251"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333258",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33025",
							  "rn": "l2-[vxlan-333258]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333258",
										  "rtt": "route-target:as2-nn2:65000:333258"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333258",
										  "rtt": "route-target:as2-nn2:65000:333258"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333385",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33152",
							  "rn": "l2-[vxlan-333385]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333385",
										  "rtt": "route-target:as2-nn2:65000:333385"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333385",
										  "rtt": "route-target:as2-nn2:65000:333385"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333778",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33545",
							  "rn": "l2-[vxlan-333778]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333778",
										  "rtt": "route-target:as2-nn2:65000:333778"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333778",
										  "rtt": "route-target:as2-nn2:65000:333778"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333428",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33195",
							  "rn": "l2-[vxlan-333428]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333428",
										  "rtt": "route-target:as2-nn2:65000:333428"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333428",
										  "rtt": "route-target:as2-nn2:65000:333428"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333334",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33101",
							  "rn": "l2-[vxlan-333334]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333334",
										  "rtt": "route-target:as2-nn2:65000:333334"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333334",
										  "rtt": "route-target:as2-nn2:65000:333334"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333371",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33138",
							  "rn": "l2-[vxlan-333371]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333371",
										  "rtt": "route-target:as2-nn2:65000:333371"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333371",
										  "rtt": "route-target:as2-nn2:65000:333371"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333268",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33035",
							  "rn": "l2-[vxlan-333268]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333268",
										  "rtt": "route-target:as2-nn2:65000:333268"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333268",
										  "rtt": "route-target:as2-nn2:65000:333268"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333286",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33053",
							  "rn": "l2-[vxlan-333286]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333286",
										  "rtt": "route-target:as2-nn2:65000:333286"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333286",
										  "rtt": "route-target:as2-nn2:65000:333286"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331484",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34251",
							  "rn": "l2-[vxlan-3331484]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331484",
										  "rtt": "route-target:as2-nn2:65000:3331484"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331484",
										  "rtt": "route-target:as2-nn2:65000:3331484"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331251",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34018",
							  "rn": "l2-[vxlan-3331251]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331251",
										  "rtt": "route-target:as2-nn2:65000:3331251"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331251",
										  "rtt": "route-target:as2-nn2:65000:3331251"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333382",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33149",
							  "rn": "l2-[vxlan-333382]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333382",
										  "rtt": "route-target:as2-nn2:65000:333382"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333382",
										  "rtt": "route-target:as2-nn2:65000:333382"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333321",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33088",
							  "rn": "l2-[vxlan-333321]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333321",
										  "rtt": "route-target:as2-nn2:65000:333321"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333321",
										  "rtt": "route-target:as2-nn2:65000:333321"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333288",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33055",
							  "rn": "l2-[vxlan-333288]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333288",
										  "rtt": "route-target:as2-nn2:65000:333288"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333288",
										  "rtt": "route-target:as2-nn2:65000:333288"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333907",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33674",
							  "rn": "l2-[vxlan-333907]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333907",
										  "rtt": "route-target:as2-nn2:65000:333907"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333907",
										  "rtt": "route-target:as2-nn2:65000:333907"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331480",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34247",
							  "rn": "l2-[vxlan-3331480]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331480",
										  "rtt": "route-target:as2-nn2:65000:3331480"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331480",
										  "rtt": "route-target:as2-nn2:65000:3331480"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333387",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33154",
							  "rn": "l2-[vxlan-333387]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333387",
										  "rtt": "route-target:as2-nn2:65000:333387"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333387",
										  "rtt": "route-target:as2-nn2:65000:333387"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331325",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34092",
							  "rn": "l2-[vxlan-3331325]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331325",
										  "rtt": "route-target:as2-nn2:65000:3331325"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331325",
										  "rtt": "route-target:as2-nn2:65000:3331325"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333938",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33705",
							  "rn": "l2-[vxlan-333938]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333938",
										  "rtt": "route-target:as2-nn2:65000:333938"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333938",
										  "rtt": "route-target:as2-nn2:65000:333938"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333865",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33632",
							  "rn": "l2-[vxlan-333865]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333865",
										  "rtt": "route-target:as2-nn2:65000:333865"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333865",
										  "rtt": "route-target:as2-nn2:65000:333865"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331330",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34097",
							  "rn": "l2-[vxlan-3331330]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331330",
										  "rtt": "route-target:as2-nn2:65000:3331330"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331330",
										  "rtt": "route-target:as2-nn2:65000:3331330"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331488",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34255",
							  "rn": "l2-[vxlan-3331488]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331488",
										  "rtt": "route-target:as2-nn2:65000:3331488"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331488",
										  "rtt": "route-target:as2-nn2:65000:3331488"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333441",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33208",
							  "rn": "l2-[vxlan-333441]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333441",
										  "rtt": "route-target:as2-nn2:65000:333441"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333441",
										  "rtt": "route-target:as2-nn2:65000:333441"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333279",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33046",
							  "rn": "l2-[vxlan-333279]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333279",
										  "rtt": "route-target:as2-nn2:65000:333279"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333279",
										  "rtt": "route-target:as2-nn2:65000:333279"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333980",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33747",
							  "rn": "l2-[vxlan-333980]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333980",
										  "rtt": "route-target:as2-nn2:65000:333980"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333980",
										  "rtt": "route-target:as2-nn2:65000:333980"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333402",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33169",
							  "rn": "l2-[vxlan-333402]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333402",
										  "rtt": "route-target:as2-nn2:65000:333402"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333402",
										  "rtt": "route-target:as2-nn2:65000:333402"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331341",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34108",
							  "rn": "l2-[vxlan-3331341]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331341",
										  "rtt": "route-target:as2-nn2:65000:3331341"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331341",
										  "rtt": "route-target:as2-nn2:65000:3331341"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333294",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33061",
							  "rn": "l2-[vxlan-333294]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333294",
										  "rtt": "route-target:as2-nn2:65000:333294"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333294",
										  "rtt": "route-target:as2-nn2:65000:333294"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333868",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33635",
							  "rn": "l2-[vxlan-333868]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333868",
										  "rtt": "route-target:as2-nn2:65000:333868"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333868",
										  "rtt": "route-target:as2-nn2:65000:333868"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331495",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34262",
							  "rn": "l2-[vxlan-3331495]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331495",
										  "rtt": "route-target:as2-nn2:65000:3331495"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331495",
										  "rtt": "route-target:as2-nn2:65000:3331495"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331462",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34229",
							  "rn": "l2-[vxlan-3331462]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331462",
										  "rtt": "route-target:as2-nn2:65000:3331462"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331462",
										  "rtt": "route-target:as2-nn2:65000:3331462"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331350",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34117",
							  "rn": "l2-[vxlan-3331350]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331350",
										  "rtt": "route-target:as2-nn2:65000:3331350"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331350",
										  "rtt": "route-target:as2-nn2:65000:3331350"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331316",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34083",
							  "rn": "l2-[vxlan-3331316]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331316",
										  "rtt": "route-target:as2-nn2:65000:3331316"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331316",
										  "rtt": "route-target:as2-nn2:65000:3331316"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333425",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33192",
							  "rn": "l2-[vxlan-333425]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333425",
										  "rtt": "route-target:as2-nn2:65000:333425"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333425",
										  "rtt": "route-target:as2-nn2:65000:333425"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333995",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33762",
							  "rn": "l2-[vxlan-333995]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333995",
										  "rtt": "route-target:as2-nn2:65000:333995"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333995",
										  "rtt": "route-target:as2-nn2:65000:333995"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333263",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33030",
							  "rn": "l2-[vxlan-333263]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333263",
										  "rtt": "route-target:as2-nn2:65000:333263"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333263",
										  "rtt": "route-target:as2-nn2:65000:333263"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333923",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33690",
							  "rn": "l2-[vxlan-333923]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333923",
										  "rtt": "route-target:as2-nn2:65000:333923"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333923",
										  "rtt": "route-target:as2-nn2:65000:333923"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333814",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33581",
							  "rn": "l2-[vxlan-333814]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333814",
										  "rtt": "route-target:as2-nn2:65000:333814"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333814",
										  "rtt": "route-target:as2-nn2:65000:333814"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333414",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33181",
							  "rn": "l2-[vxlan-333414]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333414",
										  "rtt": "route-target:as2-nn2:65000:333414"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333414",
										  "rtt": "route-target:as2-nn2:65000:333414"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331489",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34256",
							  "rn": "l2-[vxlan-3331489]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331489",
										  "rtt": "route-target:as2-nn2:65000:3331489"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331489",
										  "rtt": "route-target:as2-nn2:65000:3331489"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331284",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34051",
							  "rn": "l2-[vxlan-3331284]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331284",
										  "rtt": "route-target:as2-nn2:65000:3331284"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331284",
										  "rtt": "route-target:as2-nn2:65000:3331284"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331265",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34032",
							  "rn": "l2-[vxlan-3331265]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331265",
										  "rtt": "route-target:as2-nn2:65000:3331265"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331265",
										  "rtt": "route-target:as2-nn2:65000:3331265"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333308",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33075",
							  "rn": "l2-[vxlan-333308]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333308",
										  "rtt": "route-target:as2-nn2:65000:333308"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333308",
										  "rtt": "route-target:as2-nn2:65000:333308"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333750",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33517",
							  "rn": "l2-[vxlan-333750]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333750",
										  "rtt": "route-target:as2-nn2:65000:333750"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333750",
										  "rtt": "route-target:as2-nn2:65000:333750"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331383",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34150",
							  "rn": "l2-[vxlan-3331383]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331383",
										  "rtt": "route-target:as2-nn2:65000:3331383"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331383",
										  "rtt": "route-target:as2-nn2:65000:3331383"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333250",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33017",
							  "rn": "l2-[vxlan-333250]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333250",
										  "rtt": "route-target:as2-nn2:65000:333250"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333250",
										  "rtt": "route-target:as2-nn2:65000:333250"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333265",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33032",
							  "rn": "l2-[vxlan-333265]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333265",
										  "rtt": "route-target:as2-nn2:65000:333265"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333265",
										  "rtt": "route-target:as2-nn2:65000:333265"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333293",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33060",
							  "rn": "l2-[vxlan-333293]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333293",
										  "rtt": "route-target:as2-nn2:65000:333293"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333293",
										  "rtt": "route-target:as2-nn2:65000:333293"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333475",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33242",
							  "rn": "l2-[vxlan-333475]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333475",
										  "rtt": "route-target:as2-nn2:65000:333475"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333475",
										  "rtt": "route-target:as2-nn2:65000:333475"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331304",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34071",
							  "rn": "l2-[vxlan-3331304]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331304",
										  "rtt": "route-target:as2-nn2:65000:3331304"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331304",
										  "rtt": "route-target:as2-nn2:65000:3331304"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333420",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33187",
							  "rn": "l2-[vxlan-333420]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333420",
										  "rtt": "route-target:as2-nn2:65000:333420"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333420",
										  "rtt": "route-target:as2-nn2:65000:333420"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331388",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34155",
							  "rn": "l2-[vxlan-3331388]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331388",
										  "rtt": "route-target:as2-nn2:65000:3331388"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331388",
										  "rtt": "route-target:as2-nn2:65000:3331388"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333484",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33251",
							  "rn": "l2-[vxlan-333484]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333484",
										  "rtt": "route-target:as2-nn2:65000:333484"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333484",
										  "rtt": "route-target:as2-nn2:65000:333484"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331483",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34250",
							  "rn": "l2-[vxlan-3331483]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331483",
										  "rtt": "route-target:as2-nn2:65000:3331483"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331483",
										  "rtt": "route-target:as2-nn2:65000:3331483"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333794",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33561",
							  "rn": "l2-[vxlan-333794]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333794",
										  "rtt": "route-target:as2-nn2:65000:333794"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333794",
										  "rtt": "route-target:as2-nn2:65000:333794"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333341",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33108",
							  "rn": "l2-[vxlan-333341]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333341",
										  "rtt": "route-target:as2-nn2:65000:333341"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333341",
										  "rtt": "route-target:as2-nn2:65000:333341"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333876",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33643",
							  "rn": "l2-[vxlan-333876]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333876",
										  "rtt": "route-target:as2-nn2:65000:333876"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333876",
										  "rtt": "route-target:as2-nn2:65000:333876"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331257",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34024",
							  "rn": "l2-[vxlan-3331257]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331257",
										  "rtt": "route-target:as2-nn2:65000:3331257"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331257",
										  "rtt": "route-target:as2-nn2:65000:3331257"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331487",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34254",
							  "rn": "l2-[vxlan-3331487]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331487",
										  "rtt": "route-target:as2-nn2:65000:3331487"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331487",
										  "rtt": "route-target:as2-nn2:65000:3331487"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333295",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33062",
							  "rn": "l2-[vxlan-333295]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333295",
										  "rtt": "route-target:as2-nn2:65000:333295"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333295",
										  "rtt": "route-target:as2-nn2:65000:333295"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333285",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33052",
							  "rn": "l2-[vxlan-333285]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333285",
										  "rtt": "route-target:as2-nn2:65000:333285"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333285",
										  "rtt": "route-target:as2-nn2:65000:333285"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333363",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33130",
							  "rn": "l2-[vxlan-333363]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333363",
										  "rtt": "route-target:as2-nn2:65000:333363"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333363",
										  "rtt": "route-target:as2-nn2:65000:333363"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333751",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33518",
							  "rn": "l2-[vxlan-333751]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333751",
										  "rtt": "route-target:as2-nn2:65000:333751"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333751",
										  "rtt": "route-target:as2-nn2:65000:333751"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333253",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33020",
							  "rn": "l2-[vxlan-333253]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333253",
										  "rtt": "route-target:as2-nn2:65000:333253"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333253",
										  "rtt": "route-target:as2-nn2:65000:333253"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333988",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33755",
							  "rn": "l2-[vxlan-333988]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333988",
										  "rtt": "route-target:as2-nn2:65000:333988"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333988",
										  "rtt": "route-target:as2-nn2:65000:333988"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333467",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33234",
							  "rn": "l2-[vxlan-333467]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333467",
										  "rtt": "route-target:as2-nn2:65000:333467"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333467",
										  "rtt": "route-target:as2-nn2:65000:333467"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331479",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34246",
							  "rn": "l2-[vxlan-3331479]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331479",
										  "rtt": "route-target:as2-nn2:65000:3331479"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331479",
										  "rtt": "route-target:as2-nn2:65000:3331479"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331473",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34240",
							  "rn": "l2-[vxlan-3331473]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331473",
										  "rtt": "route-target:as2-nn2:65000:3331473"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331473",
										  "rtt": "route-target:as2-nn2:65000:3331473"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331363",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34130",
							  "rn": "l2-[vxlan-3331363]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331363",
										  "rtt": "route-target:as2-nn2:65000:3331363"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331363",
										  "rtt": "route-target:as2-nn2:65000:3331363"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333939",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33706",
							  "rn": "l2-[vxlan-333939]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333939",
										  "rtt": "route-target:as2-nn2:65000:333939"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333939",
										  "rtt": "route-target:as2-nn2:65000:333939"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-1005",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:32772",
							  "rn": "l2-[vxlan-1005]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:1005",
										  "rtt": "route-target:as2-nn2:1:1005"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:1005",
										  "rtt": "route-target:as2-nn2:1:1005"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333476",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33243",
							  "rn": "l2-[vxlan-333476]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333476",
										  "rtt": "route-target:as2-nn2:65000:333476"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333476",
										  "rtt": "route-target:as2-nn2:65000:333476"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333765",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33532",
							  "rn": "l2-[vxlan-333765]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333765",
										  "rtt": "route-target:as2-nn2:65000:333765"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333765",
										  "rtt": "route-target:as2-nn2:65000:333765"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331485",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34252",
							  "rn": "l2-[vxlan-3331485]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331485",
										  "rtt": "route-target:as2-nn2:65000:3331485"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331485",
										  "rtt": "route-target:as2-nn2:65000:3331485"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331493",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34260",
							  "rn": "l2-[vxlan-3331493]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331493",
										  "rtt": "route-target:as2-nn2:65000:3331493"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331493",
										  "rtt": "route-target:as2-nn2:65000:3331493"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333788",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33555",
							  "rn": "l2-[vxlan-333788]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333788",
										  "rtt": "route-target:as2-nn2:65000:333788"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333788",
										  "rtt": "route-target:as2-nn2:65000:333788"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333909",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33676",
							  "rn": "l2-[vxlan-333909]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333909",
										  "rtt": "route-target:as2-nn2:65000:333909"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333909",
										  "rtt": "route-target:as2-nn2:65000:333909"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333348",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33115",
							  "rn": "l2-[vxlan-333348]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333348",
										  "rtt": "route-target:as2-nn2:65000:333348"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333348",
										  "rtt": "route-target:as2-nn2:65000:333348"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331406",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34173",
							  "rn": "l2-[vxlan-3331406]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331406",
										  "rtt": "route-target:as2-nn2:65000:3331406"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331406",
										  "rtt": "route-target:as2-nn2:65000:3331406"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333789",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33556",
							  "rn": "l2-[vxlan-333789]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333789",
										  "rtt": "route-target:as2-nn2:65000:333789"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333789",
										  "rtt": "route-target:as2-nn2:65000:333789"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333276",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33043",
							  "rn": "l2-[vxlan-333276]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333276",
										  "rtt": "route-target:as2-nn2:65000:333276"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333276",
										  "rtt": "route-target:as2-nn2:65000:333276"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331491",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34258",
							  "rn": "l2-[vxlan-3331491]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331491",
										  "rtt": "route-target:as2-nn2:65000:3331491"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331491",
										  "rtt": "route-target:as2-nn2:65000:3331491"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333917",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33684",
							  "rn": "l2-[vxlan-333917]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333917",
										  "rtt": "route-target:as2-nn2:65000:333917"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333917",
										  "rtt": "route-target:as2-nn2:65000:333917"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333391",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33158",
							  "rn": "l2-[vxlan-333391]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333391",
										  "rtt": "route-target:as2-nn2:65000:333391"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333391",
										  "rtt": "route-target:as2-nn2:65000:333391"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333495",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33262",
							  "rn": "l2-[vxlan-333495]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333495",
										  "rtt": "route-target:as2-nn2:65000:333495"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333495",
										  "rtt": "route-target:as2-nn2:65000:333495"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331424",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34191",
							  "rn": "l2-[vxlan-3331424]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331424",
										  "rtt": "route-target:as2-nn2:65000:3331424"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331424",
										  "rtt": "route-target:as2-nn2:65000:3331424"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333255",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33022",
							  "rn": "l2-[vxlan-333255]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333255",
										  "rtt": "route-target:as2-nn2:65000:333255"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333255",
										  "rtt": "route-target:as2-nn2:65000:333255"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331299",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34066",
							  "rn": "l2-[vxlan-3331299]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331299",
										  "rtt": "route-target:as2-nn2:65000:3331299"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331299",
										  "rtt": "route-target:as2-nn2:65000:3331299"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333894",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33661",
							  "rn": "l2-[vxlan-333894]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333894",
										  "rtt": "route-target:as2-nn2:65000:333894"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333894",
										  "rtt": "route-target:as2-nn2:65000:333894"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333451",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33218",
							  "rn": "l2-[vxlan-333451]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333451",
										  "rtt": "route-target:as2-nn2:65000:333451"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333451",
										  "rtt": "route-target:as2-nn2:65000:333451"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333380",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33147",
							  "rn": "l2-[vxlan-333380]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333380",
										  "rtt": "route-target:as2-nn2:65000:333380"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333380",
										  "rtt": "route-target:as2-nn2:65000:333380"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333374",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33141",
							  "rn": "l2-[vxlan-333374]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333374",
										  "rtt": "route-target:as2-nn2:65000:333374"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333374",
										  "rtt": "route-target:as2-nn2:65000:333374"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331486",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34253",
							  "rn": "l2-[vxlan-3331486]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331486",
										  "rtt": "route-target:as2-nn2:65000:3331486"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331486",
										  "rtt": "route-target:as2-nn2:65000:3331486"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333410",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33177",
							  "rn": "l2-[vxlan-333410]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333410",
										  "rtt": "route-target:as2-nn2:65000:333410"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333410",
										  "rtt": "route-target:as2-nn2:65000:333410"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333408",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33175",
							  "rn": "l2-[vxlan-333408]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333408",
										  "rtt": "route-target:as2-nn2:65000:333408"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333408",
										  "rtt": "route-target:as2-nn2:65000:333408"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333313",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33080",
							  "rn": "l2-[vxlan-333313]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333313",
										  "rtt": "route-target:as2-nn2:65000:333313"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333313",
										  "rtt": "route-target:as2-nn2:65000:333313"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331266",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34033",
							  "rn": "l2-[vxlan-3331266]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331266",
										  "rtt": "route-target:as2-nn2:65000:3331266"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331266",
										  "rtt": "route-target:as2-nn2:65000:3331266"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333333",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33100",
							  "rn": "l2-[vxlan-333333]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333333",
										  "rtt": "route-target:as2-nn2:65000:333333"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333333",
										  "rtt": "route-target:as2-nn2:65000:333333"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333328",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33095",
							  "rn": "l2-[vxlan-333328]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333328",
										  "rtt": "route-target:as2-nn2:65000:333328"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333328",
										  "rtt": "route-target:as2-nn2:65000:333328"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331492",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34259",
							  "rn": "l2-[vxlan-3331492]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331492",
										  "rtt": "route-target:as2-nn2:65000:3331492"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331492",
										  "rtt": "route-target:as2-nn2:65000:3331492"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333379",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33146",
							  "rn": "l2-[vxlan-333379]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333379",
										  "rtt": "route-target:as2-nn2:65000:333379"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333379",
										  "rtt": "route-target:as2-nn2:65000:333379"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333261",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33028",
							  "rn": "l2-[vxlan-333261]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333261",
										  "rtt": "route-target:as2-nn2:65000:333261"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333261",
										  "rtt": "route-target:as2-nn2:65000:333261"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333979",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33746",
							  "rn": "l2-[vxlan-333979]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333979",
										  "rtt": "route-target:as2-nn2:65000:333979"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333979",
										  "rtt": "route-target:as2-nn2:65000:333979"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331494",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34261",
							  "rn": "l2-[vxlan-3331494]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331494",
										  "rtt": "route-target:as2-nn2:65000:3331494"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331494",
										  "rtt": "route-target:as2-nn2:65000:3331494"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333366",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33133",
							  "rn": "l2-[vxlan-333366]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333366",
										  "rtt": "route-target:as2-nn2:65000:333366"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333366",
										  "rtt": "route-target:as2-nn2:65000:333366"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333381",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33148",
							  "rn": "l2-[vxlan-333381]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333381",
										  "rtt": "route-target:as2-nn2:65000:333381"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333381",
										  "rtt": "route-target:as2-nn2:65000:333381"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333489",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33256",
							  "rn": "l2-[vxlan-333489]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333489",
										  "rtt": "route-target:as2-nn2:65000:333489"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333489",
										  "rtt": "route-target:as2-nn2:65000:333489"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333810",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33577",
							  "rn": "l2-[vxlan-333810]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333810",
										  "rtt": "route-target:as2-nn2:65000:333810"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333810",
										  "rtt": "route-target:as2-nn2:65000:333810"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333297",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33064",
							  "rn": "l2-[vxlan-333297]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333297",
										  "rtt": "route-target:as2-nn2:65000:333297"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333297",
										  "rtt": "route-target:as2-nn2:65000:333297"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333399",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33166",
							  "rn": "l2-[vxlan-333399]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333399",
										  "rtt": "route-target:as2-nn2:65000:333399"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333399",
										  "rtt": "route-target:as2-nn2:65000:333399"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333396",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33163",
							  "rn": "l2-[vxlan-333396]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333396",
										  "rtt": "route-target:as2-nn2:65000:333396"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333396",
										  "rtt": "route-target:as2-nn2:65000:333396"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333259",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33026",
							  "rn": "l2-[vxlan-333259]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333259",
										  "rtt": "route-target:as2-nn2:65000:333259"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333259",
										  "rtt": "route-target:as2-nn2:65000:333259"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333254",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33021",
							  "rn": "l2-[vxlan-333254]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333254",
										  "rtt": "route-target:as2-nn2:65000:333254"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333254",
										  "rtt": "route-target:as2-nn2:65000:333254"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331328",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34095",
							  "rn": "l2-[vxlan-3331328]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331328",
										  "rtt": "route-target:as2-nn2:65000:3331328"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331328",
										  "rtt": "route-target:as2-nn2:65000:3331328"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333349",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33116",
							  "rn": "l2-[vxlan-333349]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333349",
										  "rtt": "route-target:as2-nn2:65000:333349"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333349",
										  "rtt": "route-target:as2-nn2:65000:333349"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333761",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33528",
							  "rn": "l2-[vxlan-333761]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333761",
										  "rtt": "route-target:as2-nn2:65000:333761"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333761",
										  "rtt": "route-target:as2-nn2:65000:333761"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333359",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33126",
							  "rn": "l2-[vxlan-333359]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333359",
										  "rtt": "route-target:as2-nn2:65000:333359"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333359",
										  "rtt": "route-target:as2-nn2:65000:333359"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333355",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33122",
							  "rn": "l2-[vxlan-333355]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333355",
										  "rtt": "route-target:as2-nn2:65000:333355"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333355",
										  "rtt": "route-target:as2-nn2:65000:333355"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331496",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34263",
							  "rn": "l2-[vxlan-3331496]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331496",
										  "rtt": "route-target:as2-nn2:65000:3331496"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331496",
										  "rtt": "route-target:as2-nn2:65000:3331496"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333957",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33724",
							  "rn": "l2-[vxlan-333957]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333957",
										  "rtt": "route-target:as2-nn2:65000:333957"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333957",
										  "rtt": "route-target:as2-nn2:65000:333957"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333296",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33063",
							  "rn": "l2-[vxlan-333296]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333296",
										  "rtt": "route-target:as2-nn2:65000:333296"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333296",
										  "rtt": "route-target:as2-nn2:65000:333296"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331263",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34030",
							  "rn": "l2-[vxlan-3331263]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331263",
										  "rtt": "route-target:as2-nn2:65000:3331263"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331263",
										  "rtt": "route-target:as2-nn2:65000:3331263"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-2012007",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34774",
							  "rn": "l2-[vxlan-2012007]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:2012007",
										  "rtt": "route-target:as2-nn2:1:2012007"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:2012007",
										  "rtt": "route-target:as2-nn2:1:2012007"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333277",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33044",
							  "rn": "l2-[vxlan-333277]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333277",
										  "rtt": "route-target:as2-nn2:65000:333277"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333277",
										  "rtt": "route-target:as2-nn2:65000:333277"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331364",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34131",
							  "rn": "l2-[vxlan-3331364]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331364",
										  "rtt": "route-target:as2-nn2:65000:3331364"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331364",
										  "rtt": "route-target:as2-nn2:65000:3331364"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331499",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34266",
							  "rn": "l2-[vxlan-3331499]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331499",
										  "rtt": "route-target:as2-nn2:65000:3331499"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331499",
										  "rtt": "route-target:as2-nn2:65000:3331499"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333345",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33112",
							  "rn": "l2-[vxlan-333345]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333345",
										  "rtt": "route-target:as2-nn2:65000:333345"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333345",
										  "rtt": "route-target:as2-nn2:65000:333345"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333483",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33250",
							  "rn": "l2-[vxlan-333483]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333483",
										  "rtt": "route-target:as2-nn2:65000:333483"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333483",
										  "rtt": "route-target:as2-nn2:65000:333483"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333896",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33663",
							  "rn": "l2-[vxlan-333896]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333896",
										  "rtt": "route-target:as2-nn2:65000:333896"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333896",
										  "rtt": "route-target:as2-nn2:65000:333896"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333273",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33040",
							  "rn": "l2-[vxlan-333273]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333273",
										  "rtt": "route-target:as2-nn2:65000:333273"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333273",
										  "rtt": "route-target:as2-nn2:65000:333273"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333800",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33567",
							  "rn": "l2-[vxlan-333800]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333800",
										  "rtt": "route-target:as2-nn2:65000:333800"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333800",
										  "rtt": "route-target:as2-nn2:65000:333800"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333454",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33221",
							  "rn": "l2-[vxlan-333454]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333454",
										  "rtt": "route-target:as2-nn2:65000:333454"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333454",
										  "rtt": "route-target:as2-nn2:65000:333454"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333963",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33730",
							  "rn": "l2-[vxlan-333963]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333963",
										  "rtt": "route-target:as2-nn2:65000:333963"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333963",
										  "rtt": "route-target:as2-nn2:65000:333963"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333330",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33097",
							  "rn": "l2-[vxlan-333330]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333330",
										  "rtt": "route-target:as2-nn2:65000:333330"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333330",
										  "rtt": "route-target:as2-nn2:65000:333330"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333278",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33045",
							  "rn": "l2-[vxlan-333278]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333278",
										  "rtt": "route-target:as2-nn2:65000:333278"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333278",
										  "rtt": "route-target:as2-nn2:65000:333278"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333432",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33199",
							  "rn": "l2-[vxlan-333432]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333432",
										  "rtt": "route-target:as2-nn2:65000:333432"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333432",
										  "rtt": "route-target:as2-nn2:65000:333432"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331410",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34177",
							  "rn": "l2-[vxlan-3331410]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331410",
										  "rtt": "route-target:as2-nn2:65000:3331410"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331410",
										  "rtt": "route-target:as2-nn2:65000:3331410"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333342",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33109",
							  "rn": "l2-[vxlan-333342]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333342",
										  "rtt": "route-target:as2-nn2:65000:333342"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333342",
										  "rtt": "route-target:as2-nn2:65000:333342"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333455",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33222",
							  "rn": "l2-[vxlan-333455]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333455",
										  "rtt": "route-target:as2-nn2:65000:333455"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333455",
										  "rtt": "route-target:as2-nn2:65000:333455"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333866",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33633",
							  "rn": "l2-[vxlan-333866]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333866",
										  "rtt": "route-target:as2-nn2:65000:333866"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333866",
										  "rtt": "route-target:as2-nn2:65000:333866"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333401",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33168",
							  "rn": "l2-[vxlan-333401]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333401",
										  "rtt": "route-target:as2-nn2:65000:333401"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333401",
										  "rtt": "route-target:as2-nn2:65000:333401"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333282",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33049",
							  "rn": "l2-[vxlan-333282]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333282",
										  "rtt": "route-target:as2-nn2:65000:333282"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333282",
										  "rtt": "route-target:as2-nn2:65000:333282"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333417",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33184",
							  "rn": "l2-[vxlan-333417]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333417",
										  "rtt": "route-target:as2-nn2:65000:333417"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333417",
										  "rtt": "route-target:as2-nn2:65000:333417"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333842",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33609",
							  "rn": "l2-[vxlan-333842]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333842",
										  "rtt": "route-target:as2-nn2:65000:333842"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333842",
										  "rtt": "route-target:as2-nn2:65000:333842"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333325",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33092",
							  "rn": "l2-[vxlan-333325]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333325",
										  "rtt": "route-target:as2-nn2:65000:333325"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333325",
										  "rtt": "route-target:as2-nn2:65000:333325"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333849",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33616",
							  "rn": "l2-[vxlan-333849]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333849",
										  "rtt": "route-target:as2-nn2:65000:333849"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333849",
										  "rtt": "route-target:as2-nn2:65000:333849"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333389",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33156",
							  "rn": "l2-[vxlan-333389]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333389",
										  "rtt": "route-target:as2-nn2:65000:333389"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333389",
										  "rtt": "route-target:as2-nn2:65000:333389"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333283",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33050",
							  "rn": "l2-[vxlan-333283]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333283",
										  "rtt": "route-target:as2-nn2:65000:333283"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333283",
										  "rtt": "route-target:as2-nn2:65000:333283"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333967",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33734",
							  "rn": "l2-[vxlan-333967]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333967",
										  "rtt": "route-target:as2-nn2:65000:333967"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333967",
										  "rtt": "route-target:as2-nn2:65000:333967"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333791",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33558",
							  "rn": "l2-[vxlan-333791]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333791",
										  "rtt": "route-target:as2-nn2:65000:333791"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333791",
										  "rtt": "route-target:as2-nn2:65000:333791"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333292",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33059",
							  "rn": "l2-[vxlan-333292]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333292",
										  "rtt": "route-target:as2-nn2:65000:333292"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333292",
										  "rtt": "route-target:as2-nn2:65000:333292"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333419",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33186",
							  "rn": "l2-[vxlan-333419]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333419",
										  "rtt": "route-target:as2-nn2:65000:333419"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333419",
										  "rtt": "route-target:as2-nn2:65000:333419"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333782",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33549",
							  "rn": "l2-[vxlan-333782]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333782",
										  "rtt": "route-target:as2-nn2:65000:333782"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333782",
										  "rtt": "route-target:as2-nn2:65000:333782"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333289",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33056",
							  "rn": "l2-[vxlan-333289]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333289",
										  "rtt": "route-target:as2-nn2:65000:333289"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333289",
										  "rtt": "route-target:as2-nn2:65000:333289"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331252",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34019",
							  "rn": "l2-[vxlan-3331252]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331252",
										  "rtt": "route-target:as2-nn2:65000:3331252"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331252",
										  "rtt": "route-target:as2-nn2:65000:3331252"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333403",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33170",
							  "rn": "l2-[vxlan-333403]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333403",
										  "rtt": "route-target:as2-nn2:65000:333403"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333403",
										  "rtt": "route-target:as2-nn2:65000:333403"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333820",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33587",
							  "rn": "l2-[vxlan-333820]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333820",
										  "rtt": "route-target:as2-nn2:65000:333820"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333820",
										  "rtt": "route-target:as2-nn2:65000:333820"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333314",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33081",
							  "rn": "l2-[vxlan-333314]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333314",
										  "rtt": "route-target:as2-nn2:65000:333314"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333314",
										  "rtt": "route-target:as2-nn2:65000:333314"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333481",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33248",
							  "rn": "l2-[vxlan-333481]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333481",
										  "rtt": "route-target:as2-nn2:65000:333481"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333481",
										  "rtt": "route-target:as2-nn2:65000:333481"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333390",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33157",
							  "rn": "l2-[vxlan-333390]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333390",
										  "rtt": "route-target:as2-nn2:65000:333390"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333390",
										  "rtt": "route-target:as2-nn2:65000:333390"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333299",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33066",
							  "rn": "l2-[vxlan-333299]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333299",
										  "rtt": "route-target:as2-nn2:65000:333299"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333299",
										  "rtt": "route-target:as2-nn2:65000:333299"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331446",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34213",
							  "rn": "l2-[vxlan-3331446]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331446",
										  "rtt": "route-target:as2-nn2:65000:3331446"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331446",
										  "rtt": "route-target:as2-nn2:65000:3331446"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333494",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33261",
							  "rn": "l2-[vxlan-333494]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333494",
										  "rtt": "route-target:as2-nn2:65000:333494"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333494",
										  "rtt": "route-target:as2-nn2:65000:333494"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333882",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33649",
							  "rn": "l2-[vxlan-333882]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333882",
										  "rtt": "route-target:as2-nn2:65000:333882"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333882",
										  "rtt": "route-target:as2-nn2:65000:333882"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333304",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33071",
							  "rn": "l2-[vxlan-333304]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333304",
										  "rtt": "route-target:as2-nn2:65000:333304"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333304",
										  "rtt": "route-target:as2-nn2:65000:333304"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331303",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34070",
							  "rn": "l2-[vxlan-3331303]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331303",
										  "rtt": "route-target:as2-nn2:65000:3331303"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331303",
										  "rtt": "route-target:as2-nn2:65000:3331303"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331250",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34017",
							  "rn": "l2-[vxlan-3331250]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331250",
										  "rtt": "route-target:as2-nn2:65000:3331250"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331250",
										  "rtt": "route-target:as2-nn2:65000:3331250"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333423",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33190",
							  "rn": "l2-[vxlan-333423]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333423",
										  "rtt": "route-target:as2-nn2:65000:333423"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333423",
										  "rtt": "route-target:as2-nn2:65000:333423"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333375",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33142",
							  "rn": "l2-[vxlan-333375]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333375",
										  "rtt": "route-target:as2-nn2:65000:333375"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333375",
										  "rtt": "route-target:as2-nn2:65000:333375"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331365",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34132",
							  "rn": "l2-[vxlan-3331365]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331365",
										  "rtt": "route-target:as2-nn2:65000:3331365"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331365",
										  "rtt": "route-target:as2-nn2:65000:3331365"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333309",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33076",
							  "rn": "l2-[vxlan-333309]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333309",
										  "rtt": "route-target:as2-nn2:65000:333309"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333309",
										  "rtt": "route-target:as2-nn2:65000:333309"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333479",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33246",
							  "rn": "l2-[vxlan-333479]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333479",
										  "rtt": "route-target:as2-nn2:65000:333479"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333479",
										  "rtt": "route-target:as2-nn2:65000:333479"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333407",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33174",
							  "rn": "l2-[vxlan-333407]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333407",
										  "rtt": "route-target:as2-nn2:65000:333407"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333407",
										  "rtt": "route-target:as2-nn2:65000:333407"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333458",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33225",
							  "rn": "l2-[vxlan-333458]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333458",
										  "rtt": "route-target:as2-nn2:65000:333458"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333458",
										  "rtt": "route-target:as2-nn2:65000:333458"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333771",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33538",
							  "rn": "l2-[vxlan-333771]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333771",
										  "rtt": "route-target:as2-nn2:65000:333771"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333771",
										  "rtt": "route-target:as2-nn2:65000:333771"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333262",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33029",
							  "rn": "l2-[vxlan-333262]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333262",
										  "rtt": "route-target:as2-nn2:65000:333262"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333262",
										  "rtt": "route-target:as2-nn2:65000:333262"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333310",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33077",
							  "rn": "l2-[vxlan-333310]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333310",
										  "rtt": "route-target:as2-nn2:65000:333310"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333310",
										  "rtt": "route-target:as2-nn2:65000:333310"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333859",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33626",
							  "rn": "l2-[vxlan-333859]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333859",
										  "rtt": "route-target:as2-nn2:65000:333859"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333859",
										  "rtt": "route-target:as2-nn2:65000:333859"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333418",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33185",
							  "rn": "l2-[vxlan-333418]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333418",
										  "rtt": "route-target:as2-nn2:65000:333418"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333418",
										  "rtt": "route-target:as2-nn2:65000:333418"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333323",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33090",
							  "rn": "l2-[vxlan-333323]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333323",
										  "rtt": "route-target:as2-nn2:65000:333323"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333323",
										  "rtt": "route-target:as2-nn2:65000:333323"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333377",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33144",
							  "rn": "l2-[vxlan-333377]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333377",
										  "rtt": "route-target:as2-nn2:65000:333377"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333377",
										  "rtt": "route-target:as2-nn2:65000:333377"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333783",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33550",
							  "rn": "l2-[vxlan-333783]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333783",
										  "rtt": "route-target:as2-nn2:65000:333783"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333783",
										  "rtt": "route-target:as2-nn2:65000:333783"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333851",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33618",
							  "rn": "l2-[vxlan-333851]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333851",
										  "rtt": "route-target:as2-nn2:65000:333851"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333851",
										  "rtt": "route-target:as2-nn2:65000:333851"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333354",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33121",
							  "rn": "l2-[vxlan-333354]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333354",
										  "rtt": "route-target:as2-nn2:65000:333354"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333354",
										  "rtt": "route-target:as2-nn2:65000:333354"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333315",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33082",
							  "rn": "l2-[vxlan-333315]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333315",
										  "rtt": "route-target:as2-nn2:65000:333315"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333315",
										  "rtt": "route-target:as2-nn2:65000:333315"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331297",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34064",
							  "rn": "l2-[vxlan-3331297]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331297",
										  "rtt": "route-target:as2-nn2:65000:3331297"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331297",
										  "rtt": "route-target:as2-nn2:65000:3331297"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333384",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33151",
							  "rn": "l2-[vxlan-333384]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333384",
										  "rtt": "route-target:as2-nn2:65000:333384"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333384",
										  "rtt": "route-target:as2-nn2:65000:333384"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333312",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33079",
							  "rn": "l2-[vxlan-333312]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333312",
										  "rtt": "route-target:as2-nn2:65000:333312"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333312",
										  "rtt": "route-target:as2-nn2:65000:333312"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333316",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33083",
							  "rn": "l2-[vxlan-333316]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333316",
										  "rtt": "route-target:as2-nn2:65000:333316"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333316",
										  "rtt": "route-target:as2-nn2:65000:333316"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333378",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33145",
							  "rn": "l2-[vxlan-333378]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333378",
										  "rtt": "route-target:as2-nn2:65000:333378"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333378",
										  "rtt": "route-target:as2-nn2:65000:333378"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331453",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34220",
							  "rn": "l2-[vxlan-3331453]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331453",
										  "rtt": "route-target:as2-nn2:65000:3331453"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331453",
										  "rtt": "route-target:as2-nn2:65000:3331453"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333350",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33117",
							  "rn": "l2-[vxlan-333350]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333350",
										  "rtt": "route-target:as2-nn2:65000:333350"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333350",
										  "rtt": "route-target:as2-nn2:65000:333350"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333827",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33594",
							  "rn": "l2-[vxlan-333827]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333827",
										  "rtt": "route-target:as2-nn2:65000:333827"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333827",
										  "rtt": "route-target:as2-nn2:65000:333827"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333317",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33084",
							  "rn": "l2-[vxlan-333317]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333317",
										  "rtt": "route-target:as2-nn2:65000:333317"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333317",
										  "rtt": "route-target:as2-nn2:65000:333317"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333300",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33067",
							  "rn": "l2-[vxlan-333300]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333300",
										  "rtt": "route-target:as2-nn2:65000:333300"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333300",
										  "rtt": "route-target:as2-nn2:65000:333300"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333287",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33054",
							  "rn": "l2-[vxlan-333287]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333287",
										  "rtt": "route-target:as2-nn2:65000:333287"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333287",
										  "rtt": "route-target:as2-nn2:65000:333287"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333303",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33070",
							  "rn": "l2-[vxlan-333303]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333303",
										  "rtt": "route-target:as2-nn2:65000:333303"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333303",
										  "rtt": "route-target:as2-nn2:65000:333303"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331472",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34239",
							  "rn": "l2-[vxlan-3331472]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331472",
										  "rtt": "route-target:as2-nn2:65000:3331472"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331472",
										  "rtt": "route-target:as2-nn2:65000:3331472"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333351",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33118",
							  "rn": "l2-[vxlan-333351]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333351",
										  "rtt": "route-target:as2-nn2:65000:333351"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333351",
										  "rtt": "route-target:as2-nn2:65000:333351"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333487",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33254",
							  "rn": "l2-[vxlan-333487]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333487",
										  "rtt": "route-target:as2-nn2:65000:333487"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333487",
										  "rtt": "route-target:as2-nn2:65000:333487"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333320",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33087",
							  "rn": "l2-[vxlan-333320]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333320",
										  "rtt": "route-target:as2-nn2:65000:333320"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333320",
										  "rtt": "route-target:as2-nn2:65000:333320"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333426",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33193",
							  "rn": "l2-[vxlan-333426]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333426",
										  "rtt": "route-target:as2-nn2:65000:333426"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333426",
										  "rtt": "route-target:as2-nn2:65000:333426"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333427",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33194",
							  "rn": "l2-[vxlan-333427]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333427",
										  "rtt": "route-target:as2-nn2:65000:333427"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333427",
										  "rtt": "route-target:as2-nn2:65000:333427"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333260",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33027",
							  "rn": "l2-[vxlan-333260]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333260",
										  "rtt": "route-target:as2-nn2:65000:333260"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333260",
										  "rtt": "route-target:as2-nn2:65000:333260"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333498",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33265",
							  "rn": "l2-[vxlan-333498]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333498",
										  "rtt": "route-target:as2-nn2:65000:333498"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333498",
										  "rtt": "route-target:as2-nn2:65000:333498"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333430",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33197",
							  "rn": "l2-[vxlan-333430]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333430",
										  "rtt": "route-target:as2-nn2:65000:333430"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333430",
										  "rtt": "route-target:as2-nn2:65000:333430"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331293",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34060",
							  "rn": "l2-[vxlan-3331293]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331293",
										  "rtt": "route-target:as2-nn2:65000:3331293"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331293",
										  "rtt": "route-target:as2-nn2:65000:3331293"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-1006",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:32773",
							  "rn": "l2-[vxlan-1006]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:1006",
										  "rtt": "route-target:as2-nn2:1:1006"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:1:1006",
										  "rtt": "route-target:as2-nn2:1:1006"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333774",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33541",
							  "rn": "l2-[vxlan-333774]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333774",
										  "rtt": "route-target:as2-nn2:65000:333774"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333774",
										  "rtt": "route-target:as2-nn2:65000:333774"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333433",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33200",
							  "rn": "l2-[vxlan-333433]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333433",
										  "rtt": "route-target:as2-nn2:65000:333433"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333433",
										  "rtt": "route-target:as2-nn2:65000:333433"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331276",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34043",
							  "rn": "l2-[vxlan-3331276]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331276",
										  "rtt": "route-target:as2-nn2:65000:3331276"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331276",
										  "rtt": "route-target:as2-nn2:65000:3331276"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333453",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33220",
							  "rn": "l2-[vxlan-333453]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333453",
										  "rtt": "route-target:as2-nn2:65000:333453"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333453",
										  "rtt": "route-target:as2-nn2:65000:333453"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333981",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33748",
							  "rn": "l2-[vxlan-333981]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333981",
										  "rtt": "route-target:as2-nn2:65000:333981"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333981",
										  "rtt": "route-target:as2-nn2:65000:333981"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333434",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33201",
							  "rn": "l2-[vxlan-333434]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333434",
										  "rtt": "route-target:as2-nn2:65000:333434"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333434",
										  "rtt": "route-target:as2-nn2:65000:333434"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333435",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33202",
							  "rn": "l2-[vxlan-333435]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333435",
										  "rtt": "route-target:as2-nn2:65000:333435"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333435",
										  "rtt": "route-target:as2-nn2:65000:333435"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333439",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33206",
							  "rn": "l2-[vxlan-333439]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333439",
										  "rtt": "route-target:as2-nn2:65000:333439"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333439",
										  "rtt": "route-target:as2-nn2:65000:333439"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333447",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33214",
							  "rn": "l2-[vxlan-333447]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333447",
										  "rtt": "route-target:as2-nn2:65000:333447"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333447",
										  "rtt": "route-target:as2-nn2:65000:333447"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333758",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33525",
							  "rn": "l2-[vxlan-333758]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333758",
										  "rtt": "route-target:as2-nn2:65000:333758"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333758",
										  "rtt": "route-target:as2-nn2:65000:333758"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333436",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33203",
							  "rn": "l2-[vxlan-333436]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333436",
										  "rtt": "route-target:as2-nn2:65000:333436"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333436",
										  "rtt": "route-target:as2-nn2:65000:333436"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333931",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33698",
							  "rn": "l2-[vxlan-333931]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333931",
										  "rtt": "route-target:as2-nn2:65000:333931"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333931",
										  "rtt": "route-target:as2-nn2:65000:333931"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333465",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33232",
							  "rn": "l2-[vxlan-333465]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333465",
										  "rtt": "route-target:as2-nn2:65000:333465"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333465",
										  "rtt": "route-target:as2-nn2:65000:333465"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333485",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33252",
							  "rn": "l2-[vxlan-333485]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333485",
										  "rtt": "route-target:as2-nn2:65000:333485"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333485",
										  "rtt": "route-target:as2-nn2:65000:333485"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333361",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33128",
							  "rn": "l2-[vxlan-333361]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333361",
										  "rtt": "route-target:as2-nn2:65000:333361"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333361",
										  "rtt": "route-target:as2-nn2:65000:333361"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333473",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33240",
							  "rn": "l2-[vxlan-333473]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333473",
										  "rtt": "route-target:as2-nn2:65000:333473"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333473",
										  "rtt": "route-target:as2-nn2:65000:333473"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331321",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34088",
							  "rn": "l2-[vxlan-3331321]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331321",
										  "rtt": "route-target:as2-nn2:65000:3331321"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331321",
										  "rtt": "route-target:as2-nn2:65000:3331321"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333930",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33697",
							  "rn": "l2-[vxlan-333930]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333930",
										  "rtt": "route-target:as2-nn2:65000:333930"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333930",
										  "rtt": "route-target:as2-nn2:65000:333930"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333440",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33207",
							  "rn": "l2-[vxlan-333440]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333440",
										  "rtt": "route-target:as2-nn2:65000:333440"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333440",
										  "rtt": "route-target:as2-nn2:65000:333440"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331373",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34140",
							  "rn": "l2-[vxlan-3331373]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331373",
										  "rtt": "route-target:as2-nn2:65000:3331373"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331373",
										  "rtt": "route-target:as2-nn2:65000:3331373"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333424",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33191",
							  "rn": "l2-[vxlan-333424]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333424",
										  "rtt": "route-target:as2-nn2:65000:333424"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333424",
										  "rtt": "route-target:as2-nn2:65000:333424"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333864",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33631",
							  "rn": "l2-[vxlan-333864]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333864",
										  "rtt": "route-target:as2-nn2:65000:333864"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333864",
										  "rtt": "route-target:as2-nn2:65000:333864"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331461",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34228",
							  "rn": "l2-[vxlan-3331461]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331461",
										  "rtt": "route-target:as2-nn2:65000:3331461"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331461",
										  "rtt": "route-target:as2-nn2:65000:3331461"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333442",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33209",
							  "rn": "l2-[vxlan-333442]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333442",
										  "rtt": "route-target:as2-nn2:65000:333442"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333442",
										  "rtt": "route-target:as2-nn2:65000:333442"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333443",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33210",
							  "rn": "l2-[vxlan-333443]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333443",
										  "rtt": "route-target:as2-nn2:65000:333443"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333443",
										  "rtt": "route-target:as2-nn2:65000:333443"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333474",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33241",
							  "rn": "l2-[vxlan-333474]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333474",
										  "rtt": "route-target:as2-nn2:65000:333474"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333474",
										  "rtt": "route-target:as2-nn2:65000:333474"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333799",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33566",
							  "rn": "l2-[vxlan-333799]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333799",
										  "rtt": "route-target:as2-nn2:65000:333799"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333799",
										  "rtt": "route-target:as2-nn2:65000:333799"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333445",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33212",
							  "rn": "l2-[vxlan-333445]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333445",
										  "rtt": "route-target:as2-nn2:65000:333445"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333445",
										  "rtt": "route-target:as2-nn2:65000:333445"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331482",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34249",
							  "rn": "l2-[vxlan-3331482]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331482",
										  "rtt": "route-target:as2-nn2:65000:3331482"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331482",
										  "rtt": "route-target:as2-nn2:65000:3331482"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333857",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33624",
							  "rn": "l2-[vxlan-333857]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333857",
										  "rtt": "route-target:as2-nn2:65000:333857"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333857",
										  "rtt": "route-target:as2-nn2:65000:333857"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331444",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34211",
							  "rn": "l2-[vxlan-3331444]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331444",
										  "rtt": "route-target:as2-nn2:65000:3331444"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331444",
										  "rtt": "route-target:as2-nn2:65000:3331444"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333446",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33213",
							  "rn": "l2-[vxlan-333446]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333446",
										  "rtt": "route-target:as2-nn2:65000:333446"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333446",
										  "rtt": "route-target:as2-nn2:65000:333446"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333347",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33114",
							  "rn": "l2-[vxlan-333347]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333347",
										  "rtt": "route-target:as2-nn2:65000:333347"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333347",
										  "rtt": "route-target:as2-nn2:65000:333347"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333843",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33610",
							  "rn": "l2-[vxlan-333843]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333843",
										  "rtt": "route-target:as2-nn2:65000:333843"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333843",
										  "rtt": "route-target:as2-nn2:65000:333843"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333449",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33216",
							  "rn": "l2-[vxlan-333449]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333449",
										  "rtt": "route-target:as2-nn2:65000:333449"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333449",
										  "rtt": "route-target:as2-nn2:65000:333449"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331295",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34062",
							  "rn": "l2-[vxlan-3331295]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331295",
										  "rtt": "route-target:as2-nn2:65000:3331295"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331295",
										  "rtt": "route-target:as2-nn2:65000:3331295"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333450",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33217",
							  "rn": "l2-[vxlan-333450]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333450",
										  "rtt": "route-target:as2-nn2:65000:333450"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333450",
										  "rtt": "route-target:as2-nn2:65000:333450"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333757",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33524",
							  "rn": "l2-[vxlan-333757]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333757",
										  "rtt": "route-target:as2-nn2:65000:333757"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333757",
										  "rtt": "route-target:as2-nn2:65000:333757"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333785",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33552",
							  "rn": "l2-[vxlan-333785]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333785",
										  "rtt": "route-target:as2-nn2:65000:333785"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333785",
										  "rtt": "route-target:as2-nn2:65000:333785"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333456",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33223",
							  "rn": "l2-[vxlan-333456]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333456",
										  "rtt": "route-target:as2-nn2:65000:333456"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333456",
										  "rtt": "route-target:as2-nn2:65000:333456"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331302",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34069",
							  "rn": "l2-[vxlan-3331302]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331302",
										  "rtt": "route-target:as2-nn2:65000:3331302"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331302",
										  "rtt": "route-target:as2-nn2:65000:3331302"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333459",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33226",
							  "rn": "l2-[vxlan-333459]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333459",
										  "rtt": "route-target:as2-nn2:65000:333459"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333459",
										  "rtt": "route-target:as2-nn2:65000:333459"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333437",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33204",
							  "rn": "l2-[vxlan-333437]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333437",
										  "rtt": "route-target:as2-nn2:65000:333437"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333437",
										  "rtt": "route-target:as2-nn2:65000:333437"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333460",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33227",
							  "rn": "l2-[vxlan-333460]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333460",
										  "rtt": "route-target:as2-nn2:65000:333460"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333460",
										  "rtt": "route-target:as2-nn2:65000:333460"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333307",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33074",
							  "rn": "l2-[vxlan-333307]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333307",
										  "rtt": "route-target:as2-nn2:65000:333307"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333307",
										  "rtt": "route-target:as2-nn2:65000:333307"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333462",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33229",
							  "rn": "l2-[vxlan-333462]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333462",
										  "rtt": "route-target:as2-nn2:65000:333462"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333462",
										  "rtt": "route-target:as2-nn2:65000:333462"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333463",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33230",
							  "rn": "l2-[vxlan-333463]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333463",
										  "rtt": "route-target:as2-nn2:65000:333463"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333463",
										  "rtt": "route-target:as2-nn2:65000:333463"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333429",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33196",
							  "rn": "l2-[vxlan-333429]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333429",
										  "rtt": "route-target:as2-nn2:65000:333429"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333429",
										  "rtt": "route-target:as2-nn2:65000:333429"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333466",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33233",
							  "rn": "l2-[vxlan-333466]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333466",
										  "rtt": "route-target:as2-nn2:65000:333466"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333466",
										  "rtt": "route-target:as2-nn2:65000:333466"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333468",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33235",
							  "rn": "l2-[vxlan-333468]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333468",
										  "rtt": "route-target:as2-nn2:65000:333468"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333468",
										  "rtt": "route-target:as2-nn2:65000:333468"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333762",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33529",
							  "rn": "l2-[vxlan-333762]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333762",
										  "rtt": "route-target:as2-nn2:65000:333762"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333762",
										  "rtt": "route-target:as2-nn2:65000:333762"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331333",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34100",
							  "rn": "l2-[vxlan-3331333]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331333",
										  "rtt": "route-target:as2-nn2:65000:3331333"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331333",
										  "rtt": "route-target:as2-nn2:65000:3331333"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333795",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33562",
							  "rn": "l2-[vxlan-333795]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333795",
										  "rtt": "route-target:as2-nn2:65000:333795"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333795",
										  "rtt": "route-target:as2-nn2:65000:333795"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333469",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33236",
							  "rn": "l2-[vxlan-333469]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333469",
										  "rtt": "route-target:as2-nn2:65000:333469"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333469",
										  "rtt": "route-target:as2-nn2:65000:333469"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331308",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34075",
							  "rn": "l2-[vxlan-3331308]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331308",
										  "rtt": "route-target:as2-nn2:65000:3331308"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331308",
										  "rtt": "route-target:as2-nn2:65000:3331308"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333772",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33539",
							  "rn": "l2-[vxlan-333772]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333772",
										  "rtt": "route-target:as2-nn2:65000:333772"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333772",
										  "rtt": "route-target:as2-nn2:65000:333772"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333470",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33237",
							  "rn": "l2-[vxlan-333470]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333470",
										  "rtt": "route-target:as2-nn2:65000:333470"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333470",
										  "rtt": "route-target:as2-nn2:65000:333470"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333878",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33645",
							  "rn": "l2-[vxlan-333878]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333878",
										  "rtt": "route-target:as2-nn2:65000:333878"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333878",
										  "rtt": "route-target:as2-nn2:65000:333878"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333471",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33238",
							  "rn": "l2-[vxlan-333471]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333471",
										  "rtt": "route-target:as2-nn2:65000:333471"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333471",
										  "rtt": "route-target:as2-nn2:65000:333471"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333964",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33731",
							  "rn": "l2-[vxlan-333964]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333964",
										  "rtt": "route-target:as2-nn2:65000:333964"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333964",
										  "rtt": "route-target:as2-nn2:65000:333964"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333477",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33244",
							  "rn": "l2-[vxlan-333477]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333477",
										  "rtt": "route-target:as2-nn2:65000:333477"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333477",
										  "rtt": "route-target:as2-nn2:65000:333477"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331439",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34206",
							  "rn": "l2-[vxlan-3331439]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331439",
										  "rtt": "route-target:as2-nn2:65000:3331439"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331439",
										  "rtt": "route-target:as2-nn2:65000:3331439"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333818",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33585",
							  "rn": "l2-[vxlan-333818]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333818",
										  "rtt": "route-target:as2-nn2:65000:333818"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333818",
										  "rtt": "route-target:as2-nn2:65000:333818"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333833",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33600",
							  "rn": "l2-[vxlan-333833]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333833",
										  "rtt": "route-target:as2-nn2:65000:333833"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333833",
										  "rtt": "route-target:as2-nn2:65000:333833"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331339",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34106",
							  "rn": "l2-[vxlan-3331339]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331339",
										  "rtt": "route-target:as2-nn2:65000:3331339"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331339",
										  "rtt": "route-target:as2-nn2:65000:3331339"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333898",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33665",
							  "rn": "l2-[vxlan-333898]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333898",
										  "rtt": "route-target:as2-nn2:65000:333898"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333898",
										  "rtt": "route-target:as2-nn2:65000:333898"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333490",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33257",
							  "rn": "l2-[vxlan-333490]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333490",
										  "rtt": "route-target:as2-nn2:65000:333490"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333490",
										  "rtt": "route-target:as2-nn2:65000:333490"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333853",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33620",
							  "rn": "l2-[vxlan-333853]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333853",
										  "rtt": "route-target:as2-nn2:65000:333853"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333853",
										  "rtt": "route-target:as2-nn2:65000:333853"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333491",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33258",
							  "rn": "l2-[vxlan-333491]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333491",
										  "rtt": "route-target:as2-nn2:65000:333491"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333491",
										  "rtt": "route-target:as2-nn2:65000:333491"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331259",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34026",
							  "rn": "l2-[vxlan-3331259]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331259",
										  "rtt": "route-target:as2-nn2:65000:3331259"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331259",
										  "rtt": "route-target:as2-nn2:65000:3331259"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333928",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33695",
							  "rn": "l2-[vxlan-333928]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333928",
										  "rtt": "route-target:as2-nn2:65000:333928"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333928",
										  "rtt": "route-target:as2-nn2:65000:333928"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333280",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33047",
							  "rn": "l2-[vxlan-333280]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333280",
										  "rtt": "route-target:as2-nn2:65000:333280"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333280",
										  "rtt": "route-target:as2-nn2:65000:333280"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333493",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33260",
							  "rn": "l2-[vxlan-333493]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333493",
										  "rtt": "route-target:as2-nn2:65000:333493"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333493",
										  "rtt": "route-target:as2-nn2:65000:333493"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333496",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33263",
							  "rn": "l2-[vxlan-333496]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333496",
										  "rtt": "route-target:as2-nn2:65000:333496"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333496",
										  "rtt": "route-target:as2-nn2:65000:333496"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333302",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33069",
							  "rn": "l2-[vxlan-333302]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333302",
										  "rtt": "route-target:as2-nn2:65000:333302"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333302",
										  "rtt": "route-target:as2-nn2:65000:333302"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333488",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33255",
							  "rn": "l2-[vxlan-333488]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333488",
										  "rtt": "route-target:as2-nn2:65000:333488"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333488",
										  "rtt": "route-target:as2-nn2:65000:333488"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333497",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33264",
							  "rn": "l2-[vxlan-333497]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333497",
										  "rtt": "route-target:as2-nn2:65000:333497"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333497",
										  "rtt": "route-target:as2-nn2:65000:333497"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333753",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33520",
							  "rn": "l2-[vxlan-333753]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333753",
										  "rtt": "route-target:as2-nn2:65000:333753"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333753",
										  "rtt": "route-target:as2-nn2:65000:333753"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331434",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34201",
							  "rn": "l2-[vxlan-3331434]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331434",
										  "rtt": "route-target:as2-nn2:65000:3331434"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331434",
										  "rtt": "route-target:as2-nn2:65000:3331434"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333754",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33521",
							  "rn": "l2-[vxlan-333754]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333754",
										  "rtt": "route-target:as2-nn2:65000:333754"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333754",
										  "rtt": "route-target:as2-nn2:65000:333754"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333808",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33575",
							  "rn": "l2-[vxlan-333808]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333808",
										  "rtt": "route-target:as2-nn2:65000:333808"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333808",
										  "rtt": "route-target:as2-nn2:65000:333808"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333994",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33761",
							  "rn": "l2-[vxlan-333994]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333994",
										  "rtt": "route-target:as2-nn2:65000:333994"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333994",
										  "rtt": "route-target:as2-nn2:65000:333994"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333755",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33522",
							  "rn": "l2-[vxlan-333755]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333755",
										  "rtt": "route-target:as2-nn2:65000:333755"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333755",
										  "rtt": "route-target:as2-nn2:65000:333755"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331342",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34109",
							  "rn": "l2-[vxlan-3331342]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331342",
										  "rtt": "route-target:as2-nn2:65000:3331342"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331342",
										  "rtt": "route-target:as2-nn2:65000:3331342"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333759",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33526",
							  "rn": "l2-[vxlan-333759]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333759",
										  "rtt": "route-target:as2-nn2:65000:333759"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333759",
										  "rtt": "route-target:as2-nn2:65000:333759"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333760",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33527",
							  "rn": "l2-[vxlan-333760]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333760",
										  "rtt": "route-target:as2-nn2:65000:333760"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333760",
										  "rtt": "route-target:as2-nn2:65000:333760"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333764",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33531",
							  "rn": "l2-[vxlan-333764]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333764",
										  "rtt": "route-target:as2-nn2:65000:333764"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333764",
										  "rtt": "route-target:as2-nn2:65000:333764"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333766",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33533",
							  "rn": "l2-[vxlan-333766]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333766",
										  "rtt": "route-target:as2-nn2:65000:333766"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333766",
										  "rtt": "route-target:as2-nn2:65000:333766"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331498",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34265",
							  "rn": "l2-[vxlan-3331498]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331498",
										  "rtt": "route-target:as2-nn2:65000:3331498"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331498",
										  "rtt": "route-target:as2-nn2:65000:3331498"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333769",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33536",
							  "rn": "l2-[vxlan-333769]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333769",
										  "rtt": "route-target:as2-nn2:65000:333769"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333769",
										  "rtt": "route-target:as2-nn2:65000:333769"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333773",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33540",
							  "rn": "l2-[vxlan-333773]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333773",
										  "rtt": "route-target:as2-nn2:65000:333773"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333773",
										  "rtt": "route-target:as2-nn2:65000:333773"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333388",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33155",
							  "rn": "l2-[vxlan-333388]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333388",
										  "rtt": "route-target:as2-nn2:65000:333388"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333388",
										  "rtt": "route-target:as2-nn2:65000:333388"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333804",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33571",
							  "rn": "l2-[vxlan-333804]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333804",
										  "rtt": "route-target:as2-nn2:65000:333804"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333804",
										  "rtt": "route-target:as2-nn2:65000:333804"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333775",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33542",
							  "rn": "l2-[vxlan-333775]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333775",
										  "rtt": "route-target:as2-nn2:65000:333775"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333775",
										  "rtt": "route-target:as2-nn2:65000:333775"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333908",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33675",
							  "rn": "l2-[vxlan-333908]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333908",
										  "rtt": "route-target:as2-nn2:65000:333908"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333908",
										  "rtt": "route-target:as2-nn2:65000:333908"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333779",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33546",
							  "rn": "l2-[vxlan-333779]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333779",
										  "rtt": "route-target:as2-nn2:65000:333779"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333779",
										  "rtt": "route-target:as2-nn2:65000:333779"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333921",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33688",
							  "rn": "l2-[vxlan-333921]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333921",
										  "rtt": "route-target:as2-nn2:65000:333921"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333921",
										  "rtt": "route-target:as2-nn2:65000:333921"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333409",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33176",
							  "rn": "l2-[vxlan-333409]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333409",
										  "rtt": "route-target:as2-nn2:65000:333409"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333409",
										  "rtt": "route-target:as2-nn2:65000:333409"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333482",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33249",
							  "rn": "l2-[vxlan-333482]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333482",
										  "rtt": "route-target:as2-nn2:65000:333482"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333482",
										  "rtt": "route-target:as2-nn2:65000:333482"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333784",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33551",
							  "rn": "l2-[vxlan-333784]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333784",
										  "rtt": "route-target:as2-nn2:65000:333784"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333784",
										  "rtt": "route-target:as2-nn2:65000:333784"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331431",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34198",
							  "rn": "l2-[vxlan-3331431]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331431",
										  "rtt": "route-target:as2-nn2:65000:3331431"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331431",
										  "rtt": "route-target:as2-nn2:65000:3331431"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333786",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33553",
							  "rn": "l2-[vxlan-333786]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333786",
										  "rtt": "route-target:as2-nn2:65000:333786"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333786",
										  "rtt": "route-target:as2-nn2:65000:333786"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333961",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33728",
							  "rn": "l2-[vxlan-333961]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333961",
										  "rtt": "route-target:as2-nn2:65000:333961"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333961",
										  "rtt": "route-target:as2-nn2:65000:333961"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333787",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33554",
							  "rn": "l2-[vxlan-333787]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333787",
										  "rtt": "route-target:as2-nn2:65000:333787"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333787",
										  "rtt": "route-target:as2-nn2:65000:333787"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333949",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33716",
							  "rn": "l2-[vxlan-333949]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333949",
										  "rtt": "route-target:as2-nn2:65000:333949"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333949",
										  "rtt": "route-target:as2-nn2:65000:333949"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333790",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33557",
							  "rn": "l2-[vxlan-333790]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333790",
										  "rtt": "route-target:as2-nn2:65000:333790"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333790",
										  "rtt": "route-target:as2-nn2:65000:333790"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333839",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33606",
							  "rn": "l2-[vxlan-333839]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333839",
										  "rtt": "route-target:as2-nn2:65000:333839"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333839",
										  "rtt": "route-target:as2-nn2:65000:333839"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331280",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34047",
							  "rn": "l2-[vxlan-3331280]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331280",
										  "rtt": "route-target:as2-nn2:65000:3331280"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331280",
										  "rtt": "route-target:as2-nn2:65000:3331280"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333398",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33165",
							  "rn": "l2-[vxlan-333398]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333398",
										  "rtt": "route-target:as2-nn2:65000:333398"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333398",
										  "rtt": "route-target:as2-nn2:65000:333398"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333792",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33559",
							  "rn": "l2-[vxlan-333792]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333792",
										  "rtt": "route-target:as2-nn2:65000:333792"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333792",
										  "rtt": "route-target:as2-nn2:65000:333792"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333793",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33560",
							  "rn": "l2-[vxlan-333793]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333793",
										  "rtt": "route-target:as2-nn2:65000:333793"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333793",
										  "rtt": "route-target:as2-nn2:65000:333793"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333796",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33563",
							  "rn": "l2-[vxlan-333796]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333796",
										  "rtt": "route-target:as2-nn2:65000:333796"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333796",
										  "rtt": "route-target:as2-nn2:65000:333796"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333797",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33564",
							  "rn": "l2-[vxlan-333797]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333797",
										  "rtt": "route-target:as2-nn2:65000:333797"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333797",
										  "rtt": "route-target:as2-nn2:65000:333797"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331474",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34241",
							  "rn": "l2-[vxlan-3331474]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331474",
										  "rtt": "route-target:as2-nn2:65000:3331474"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331474",
										  "rtt": "route-target:as2-nn2:65000:3331474"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331471",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34238",
							  "rn": "l2-[vxlan-3331471]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331471",
										  "rtt": "route-target:as2-nn2:65000:3331471"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331471",
										  "rtt": "route-target:as2-nn2:65000:3331471"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333798",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33565",
							  "rn": "l2-[vxlan-333798]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333798",
										  "rtt": "route-target:as2-nn2:65000:333798"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333798",
										  "rtt": "route-target:as2-nn2:65000:333798"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333801",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33568",
							  "rn": "l2-[vxlan-333801]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333801",
										  "rtt": "route-target:as2-nn2:65000:333801"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333801",
										  "rtt": "route-target:as2-nn2:65000:333801"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333257",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33024",
							  "rn": "l2-[vxlan-333257]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333257",
										  "rtt": "route-target:as2-nn2:65000:333257"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333257",
										  "rtt": "route-target:as2-nn2:65000:333257"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333805",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33572",
							  "rn": "l2-[vxlan-333805]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333805",
										  "rtt": "route-target:as2-nn2:65000:333805"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333805",
										  "rtt": "route-target:as2-nn2:65000:333805"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331389",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34156",
							  "rn": "l2-[vxlan-3331389]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331389",
										  "rtt": "route-target:as2-nn2:65000:3331389"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331389",
										  "rtt": "route-target:as2-nn2:65000:3331389"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333356",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33123",
							  "rn": "l2-[vxlan-333356]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333356",
										  "rtt": "route-target:as2-nn2:65000:333356"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333356",
										  "rtt": "route-target:as2-nn2:65000:333356"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333306",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33073",
							  "rn": "l2-[vxlan-333306]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333306",
										  "rtt": "route-target:as2-nn2:65000:333306"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333306",
										  "rtt": "route-target:as2-nn2:65000:333306"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333802",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33569",
							  "rn": "l2-[vxlan-333802]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333802",
										  "rtt": "route-target:as2-nn2:65000:333802"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333802",
										  "rtt": "route-target:as2-nn2:65000:333802"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333879",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33646",
							  "rn": "l2-[vxlan-333879]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333879",
										  "rtt": "route-target:as2-nn2:65000:333879"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333879",
										  "rtt": "route-target:as2-nn2:65000:333879"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333803",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33570",
							  "rn": "l2-[vxlan-333803]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333803",
										  "rtt": "route-target:as2-nn2:65000:333803"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333803",
										  "rtt": "route-target:as2-nn2:65000:333803"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333806",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33573",
							  "rn": "l2-[vxlan-333806]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333806",
										  "rtt": "route-target:as2-nn2:65000:333806"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333806",
										  "rtt": "route-target:as2-nn2:65000:333806"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333324",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33091",
							  "rn": "l2-[vxlan-333324]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333324",
										  "rtt": "route-target:as2-nn2:65000:333324"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333324",
										  "rtt": "route-target:as2-nn2:65000:333324"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333829",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33596",
							  "rn": "l2-[vxlan-333829]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333829",
										  "rtt": "route-target:as2-nn2:65000:333829"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333829",
										  "rtt": "route-target:as2-nn2:65000:333829"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331438",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34205",
							  "rn": "l2-[vxlan-3331438]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331438",
										  "rtt": "route-target:as2-nn2:65000:3331438"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331438",
										  "rtt": "route-target:as2-nn2:65000:3331438"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331427",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34194",
							  "rn": "l2-[vxlan-3331427]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331427",
										  "rtt": "route-target:as2-nn2:65000:3331427"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331427",
										  "rtt": "route-target:as2-nn2:65000:3331427"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333438",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33205",
							  "rn": "l2-[vxlan-333438]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333438",
										  "rtt": "route-target:as2-nn2:65000:333438"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333438",
										  "rtt": "route-target:as2-nn2:65000:333438"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333311",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33078",
							  "rn": "l2-[vxlan-333311]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333311",
										  "rtt": "route-target:as2-nn2:65000:333311"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333311",
										  "rtt": "route-target:as2-nn2:65000:333311"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333809",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33576",
							  "rn": "l2-[vxlan-333809]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333809",
										  "rtt": "route-target:as2-nn2:65000:333809"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333809",
										  "rtt": "route-target:as2-nn2:65000:333809"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331466",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34233",
							  "rn": "l2-[vxlan-3331466]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331466",
										  "rtt": "route-target:as2-nn2:65000:3331466"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331466",
										  "rtt": "route-target:as2-nn2:65000:3331466"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333811",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33578",
							  "rn": "l2-[vxlan-333811]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333811",
										  "rtt": "route-target:as2-nn2:65000:333811"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333811",
										  "rtt": "route-target:as2-nn2:65000:333811"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331448",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34215",
							  "rn": "l2-[vxlan-3331448]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331448",
										  "rtt": "route-target:as2-nn2:65000:3331448"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331448",
										  "rtt": "route-target:as2-nn2:65000:3331448"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333871",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33638",
							  "rn": "l2-[vxlan-333871]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333871",
										  "rtt": "route-target:as2-nn2:65000:333871"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333871",
										  "rtt": "route-target:as2-nn2:65000:333871"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331355",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34122",
							  "rn": "l2-[vxlan-3331355]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331355",
										  "rtt": "route-target:as2-nn2:65000:3331355"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331355",
										  "rtt": "route-target:as2-nn2:65000:3331355"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333812",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33579",
							  "rn": "l2-[vxlan-333812]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333812",
										  "rtt": "route-target:as2-nn2:65000:333812"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333812",
										  "rtt": "route-target:as2-nn2:65000:333812"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333947",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33714",
							  "rn": "l2-[vxlan-333947]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333947",
										  "rtt": "route-target:as2-nn2:65000:333947"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333947",
										  "rtt": "route-target:as2-nn2:65000:333947"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333828",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33595",
							  "rn": "l2-[vxlan-333828]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333828",
										  "rtt": "route-target:as2-nn2:65000:333828"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333828",
										  "rtt": "route-target:as2-nn2:65000:333828"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331253",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34020",
							  "rn": "l2-[vxlan-3331253]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331253",
										  "rtt": "route-target:as2-nn2:65000:3331253"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331253",
										  "rtt": "route-target:as2-nn2:65000:3331253"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333813",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33580",
							  "rn": "l2-[vxlan-333813]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333813",
										  "rtt": "route-target:as2-nn2:65000:333813"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333813",
										  "rtt": "route-target:as2-nn2:65000:333813"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333272",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33039",
							  "rn": "l2-[vxlan-333272]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333272",
										  "rtt": "route-target:as2-nn2:65000:333272"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333272",
										  "rtt": "route-target:as2-nn2:65000:333272"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333881",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33648",
							  "rn": "l2-[vxlan-333881]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333881",
										  "rtt": "route-target:as2-nn2:65000:333881"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333881",
										  "rtt": "route-target:as2-nn2:65000:333881"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333815",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33582",
							  "rn": "l2-[vxlan-333815]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333815",
										  "rtt": "route-target:as2-nn2:65000:333815"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333815",
										  "rtt": "route-target:as2-nn2:65000:333815"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333816",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33583",
							  "rn": "l2-[vxlan-333816]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333816",
										  "rtt": "route-target:as2-nn2:65000:333816"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333816",
										  "rtt": "route-target:as2-nn2:65000:333816"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331332",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34099",
							  "rn": "l2-[vxlan-3331332]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331332",
										  "rtt": "route-target:as2-nn2:65000:3331332"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331332",
										  "rtt": "route-target:as2-nn2:65000:3331332"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331254",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34021",
							  "rn": "l2-[vxlan-3331254]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331254",
										  "rtt": "route-target:as2-nn2:65000:3331254"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331254",
										  "rtt": "route-target:as2-nn2:65000:3331254"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333817",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33584",
							  "rn": "l2-[vxlan-333817]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333817",
										  "rtt": "route-target:as2-nn2:65000:333817"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333817",
										  "rtt": "route-target:as2-nn2:65000:333817"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331478",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34245",
							  "rn": "l2-[vxlan-3331478]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331478",
										  "rtt": "route-target:as2-nn2:65000:3331478"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331478",
										  "rtt": "route-target:as2-nn2:65000:3331478"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333819",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33586",
							  "rn": "l2-[vxlan-333819]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333819",
										  "rtt": "route-target:as2-nn2:65000:333819"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333819",
										  "rtt": "route-target:as2-nn2:65000:333819"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333989",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33756",
							  "rn": "l2-[vxlan-333989]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333989",
										  "rtt": "route-target:as2-nn2:65000:333989"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333989",
										  "rtt": "route-target:as2-nn2:65000:333989"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333821",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33588",
							  "rn": "l2-[vxlan-333821]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333821",
										  "rtt": "route-target:as2-nn2:65000:333821"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333821",
										  "rtt": "route-target:as2-nn2:65000:333821"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333822",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33589",
							  "rn": "l2-[vxlan-333822]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333822",
										  "rtt": "route-target:as2-nn2:65000:333822"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333822",
										  "rtt": "route-target:as2-nn2:65000:333822"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333823",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33590",
							  "rn": "l2-[vxlan-333823]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333823",
										  "rtt": "route-target:as2-nn2:65000:333823"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333823",
										  "rtt": "route-target:as2-nn2:65000:333823"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331319",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34086",
							  "rn": "l2-[vxlan-3331319]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331319",
										  "rtt": "route-target:as2-nn2:65000:3331319"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331319",
										  "rtt": "route-target:as2-nn2:65000:3331319"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333845",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33612",
							  "rn": "l2-[vxlan-333845]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333845",
										  "rtt": "route-target:as2-nn2:65000:333845"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333845",
										  "rtt": "route-target:as2-nn2:65000:333845"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333861",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33628",
							  "rn": "l2-[vxlan-333861]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333861",
										  "rtt": "route-target:as2-nn2:65000:333861"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333861",
										  "rtt": "route-target:as2-nn2:65000:333861"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331464",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34231",
							  "rn": "l2-[vxlan-3331464]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331464",
										  "rtt": "route-target:as2-nn2:65000:3331464"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331464",
										  "rtt": "route-target:as2-nn2:65000:3331464"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333824",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33591",
							  "rn": "l2-[vxlan-333824]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333824",
										  "rtt": "route-target:as2-nn2:65000:333824"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333824",
										  "rtt": "route-target:as2-nn2:65000:333824"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333854",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33621",
							  "rn": "l2-[vxlan-333854]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333854",
										  "rtt": "route-target:as2-nn2:65000:333854"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333854",
										  "rtt": "route-target:as2-nn2:65000:333854"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333825",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33592",
							  "rn": "l2-[vxlan-333825]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333825",
										  "rtt": "route-target:as2-nn2:65000:333825"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333825",
										  "rtt": "route-target:as2-nn2:65000:333825"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331317",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34084",
							  "rn": "l2-[vxlan-3331317]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331317",
										  "rtt": "route-target:as2-nn2:65000:3331317"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331317",
										  "rtt": "route-target:as2-nn2:65000:3331317"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333826",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33593",
							  "rn": "l2-[vxlan-333826]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333826",
										  "rtt": "route-target:as2-nn2:65000:333826"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333826",
										  "rtt": "route-target:as2-nn2:65000:333826"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333318",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33085",
							  "rn": "l2-[vxlan-333318]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333318",
										  "rtt": "route-target:as2-nn2:65000:333318"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333318",
										  "rtt": "route-target:as2-nn2:65000:333318"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333834",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33601",
							  "rn": "l2-[vxlan-333834]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333834",
										  "rtt": "route-target:as2-nn2:65000:333834"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333834",
										  "rtt": "route-target:as2-nn2:65000:333834"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331382",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34149",
							  "rn": "l2-[vxlan-3331382]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331382",
										  "rtt": "route-target:as2-nn2:65000:3331382"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331382",
										  "rtt": "route-target:as2-nn2:65000:3331382"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333884",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33651",
							  "rn": "l2-[vxlan-333884]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333884",
										  "rtt": "route-target:as2-nn2:65000:333884"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333884",
										  "rtt": "route-target:as2-nn2:65000:333884"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333777",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33544",
							  "rn": "l2-[vxlan-333777]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333777",
										  "rtt": "route-target:as2-nn2:65000:333777"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333777",
										  "rtt": "route-target:as2-nn2:65000:333777"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333835",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33602",
							  "rn": "l2-[vxlan-333835]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333835",
										  "rtt": "route-target:as2-nn2:65000:333835"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333835",
										  "rtt": "route-target:as2-nn2:65000:333835"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333890",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33657",
							  "rn": "l2-[vxlan-333890]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333890",
										  "rtt": "route-target:as2-nn2:65000:333890"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333890",
										  "rtt": "route-target:as2-nn2:65000:333890"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333395",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33162",
							  "rn": "l2-[vxlan-333395]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333395",
										  "rtt": "route-target:as2-nn2:65000:333395"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333395",
										  "rtt": "route-target:as2-nn2:65000:333395"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333327",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33094",
							  "rn": "l2-[vxlan-333327]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333327",
										  "rtt": "route-target:as2-nn2:65000:333327"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333327",
										  "rtt": "route-target:as2-nn2:65000:333327"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333274",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33041",
							  "rn": "l2-[vxlan-333274]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333274",
										  "rtt": "route-target:as2-nn2:65000:333274"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333274",
										  "rtt": "route-target:as2-nn2:65000:333274"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333836",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33603",
							  "rn": "l2-[vxlan-333836]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333836",
										  "rtt": "route-target:as2-nn2:65000:333836"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333836",
										  "rtt": "route-target:as2-nn2:65000:333836"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333837",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33604",
							  "rn": "l2-[vxlan-333837]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333837",
										  "rtt": "route-target:as2-nn2:65000:333837"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333837",
										  "rtt": "route-target:as2-nn2:65000:333837"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333838",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33605",
							  "rn": "l2-[vxlan-333838]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333838",
										  "rtt": "route-target:as2-nn2:65000:333838"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333838",
										  "rtt": "route-target:as2-nn2:65000:333838"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333840",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33607",
							  "rn": "l2-[vxlan-333840]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333840",
										  "rtt": "route-target:as2-nn2:65000:333840"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333840",
										  "rtt": "route-target:as2-nn2:65000:333840"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333281",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33048",
							  "rn": "l2-[vxlan-333281]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333281",
										  "rtt": "route-target:as2-nn2:65000:333281"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333281",
										  "rtt": "route-target:as2-nn2:65000:333281"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333874",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33641",
							  "rn": "l2-[vxlan-333874]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333874",
										  "rtt": "route-target:as2-nn2:65000:333874"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333874",
										  "rtt": "route-target:as2-nn2:65000:333874"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333841",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33608",
							  "rn": "l2-[vxlan-333841]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333841",
										  "rtt": "route-target:as2-nn2:65000:333841"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333841",
										  "rtt": "route-target:as2-nn2:65000:333841"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333887",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33654",
							  "rn": "l2-[vxlan-333887]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333887",
										  "rtt": "route-target:as2-nn2:65000:333887"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333887",
										  "rtt": "route-target:as2-nn2:65000:333887"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333847",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33614",
							  "rn": "l2-[vxlan-333847]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333847",
										  "rtt": "route-target:as2-nn2:65000:333847"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333847",
										  "rtt": "route-target:as2-nn2:65000:333847"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333848",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33615",
							  "rn": "l2-[vxlan-333848]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333848",
										  "rtt": "route-target:as2-nn2:65000:333848"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333848",
										  "rtt": "route-target:as2-nn2:65000:333848"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333856",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33623",
							  "rn": "l2-[vxlan-333856]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333856",
										  "rtt": "route-target:as2-nn2:65000:333856"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333856",
										  "rtt": "route-target:as2-nn2:65000:333856"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333464",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33231",
							  "rn": "l2-[vxlan-333464]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333464",
										  "rtt": "route-target:as2-nn2:65000:333464"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333464",
										  "rtt": "route-target:as2-nn2:65000:333464"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331481",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34248",
							  "rn": "l2-[vxlan-3331481]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331481",
										  "rtt": "route-target:as2-nn2:65000:3331481"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331481",
										  "rtt": "route-target:as2-nn2:65000:3331481"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333858",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33625",
							  "rn": "l2-[vxlan-333858]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333858",
										  "rtt": "route-target:as2-nn2:65000:333858"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333858",
										  "rtt": "route-target:as2-nn2:65000:333858"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333860",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33627",
							  "rn": "l2-[vxlan-333860]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333860",
										  "rtt": "route-target:as2-nn2:65000:333860"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333860",
										  "rtt": "route-target:as2-nn2:65000:333860"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333335",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33102",
							  "rn": "l2-[vxlan-333335]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333335",
										  "rtt": "route-target:as2-nn2:65000:333335"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333335",
										  "rtt": "route-target:as2-nn2:65000:333335"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333862",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33629",
							  "rn": "l2-[vxlan-333862]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333862",
										  "rtt": "route-target:as2-nn2:65000:333862"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333862",
										  "rtt": "route-target:as2-nn2:65000:333862"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333867",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33634",
							  "rn": "l2-[vxlan-333867]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333867",
										  "rtt": "route-target:as2-nn2:65000:333867"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333867",
										  "rtt": "route-target:as2-nn2:65000:333867"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333780",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33547",
							  "rn": "l2-[vxlan-333780]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333780",
										  "rtt": "route-target:as2-nn2:65000:333780"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333780",
										  "rtt": "route-target:as2-nn2:65000:333780"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333877",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33644",
							  "rn": "l2-[vxlan-333877]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333877",
										  "rtt": "route-target:as2-nn2:65000:333877"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333877",
										  "rtt": "route-target:as2-nn2:65000:333877"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333319",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33086",
							  "rn": "l2-[vxlan-333319]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333319",
										  "rtt": "route-target:as2-nn2:65000:333319"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333319",
										  "rtt": "route-target:as2-nn2:65000:333319"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333870",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33637",
							  "rn": "l2-[vxlan-333870]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333870",
										  "rtt": "route-target:as2-nn2:65000:333870"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333870",
										  "rtt": "route-target:as2-nn2:65000:333870"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-3331346",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:34113",
							  "rn": "l2-[vxlan-3331346]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331346",
										  "rtt": "route-target:as2-nn2:65000:3331346"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:3331346",
										  "rtt": "route-target:as2-nn2:65000:3331346"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333763",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33530",
							  "rn": "l2-[vxlan-333763]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333763",
										  "rtt": "route-target:as2-nn2:65000:333763"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333763",
										  "rtt": "route-target:as2-nn2:65000:333763"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333872",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33639",
							  "rn": "l2-[vxlan-333872]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333872",
										  "rtt": "route-target:as2-nn2:65000:333872"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333872",
										  "rtt": "route-target:as2-nn2:65000:333872"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333343",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33110",
							  "rn": "l2-[vxlan-333343]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333343",
										  "rtt": "route-target:as2-nn2:65000:333343"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333343",
										  "rtt": "route-target:as2-nn2:65000:333343"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333873",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33640",
							  "rn": "l2-[vxlan-333873]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333873",
										  "rtt": "route-target:as2-nn2:65000:333873"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333873",
										  "rtt": "route-target:as2-nn2:65000:333873"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpOperRtctrlL2": {
							"attributes": {
							  "encap": "vxlan-333875",
							  "name": "",
							  "rd": "rd:ipv4-nn2:172.16.1.1:33642",
							  "rn": "l2-[vxlan-333875]"
							},
							"children": [
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-import",
									"type": "import"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333875",
										  "rtt": "route-target:as2-nn2:65000:333875"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpOperRttP": {
								  "attributes": {
									"descr": "",
									"name": "",
									"rn": "rttp-export",
									"type": "export"
								  },
								  "children": [
									{
									  "bgpOperRttEntry": {
										"attributes": {
										  "rn": "entry-route-target:as2-nn2:65000:333875",
										  "rtt": "route-target:as2-nn2:65000:333875"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						}
					  ]
					}
				  },
				  {
					"bgpEvtHist": {
					  "attributes": {
						"childAction": "",
						"modTs": "2020-10-01T08:53:57.611+00:00",
						"rn": "evthist-errors",
						"size": "default",
						"status": "",
						"type": "errors"
					  }
					}
				  },
				  {
					"bgpEvtHist": {
					  "attributes": {
						"childAction": "",
						"modTs": "2020-10-01T08:53:57.611+00:00",
						"rn": "evthist-events",
						"size": "default",
						"status": "",
						"type": "events"
					  }
					}
				  },
				  {
					"bgpEvtHist": {
					  "attributes": {
						"childAction": "",
						"modTs": "2020-10-01T08:53:57.611+00:00",
						"rn": "evthist-cli",
						"size": "default",
						"status": "",
						"type": "cli"
					  }
					}
				  },
				  {
					"bgpDom": {
					  "attributes": {
						"allocIndex": "0",
						"always": "disabled",
						"bestPathIntvl": "300",
						"bgpCfgFailedBmp": "",
						"bgpCfgFailedTs": "0",
						"bgpCfgState": "0",
						"childAction": "",
						"clusterId": "",
						"firstPeerUpTs": "2020-11-05T06:38:19.291+00:00",
						"holdIntvl": "180",
						"id": "3",
						"kaIntvl": "60",
						"localAsn": "",
						"maxAsLimit": "0",
						"modTs": "2020-10-01T08:54:00.971+00:00",
						"mode": "fabric",
						"name": "iAZ",
						"numEstPeers": "0",
						"numPeers": "1",
						"numPeersPending": "0",
						"operRtrId": "11.1.1.1",
						"operSt": "up",
						"pfxPeerTimeout": "90",
						"pfxPeerWaitTime": "90",
						"reConnIntvl": "60",
						"rn": "dom-iAZ",
						"routerMac": "70:EA:1A:36:1D:2F",
						"rtrId": "0.0.0.0",
						"status": "",
						"vnid": "2010000",
						"vtepIp": "172.16.1.101",
						"vtepVirtIp": "172.17.1.1"
					  },
					  "children": [
						{
						  "bgpPeer": {
							"attributes": {
							  "activePfxPeers": "0",
							  "addr": "10.1.11.3",
							  "adminSt": "enabled",
							  "affGrp": "0",
							  "asn": "1000",
							  "bgpCfgFailedBmp": "",
							  "bgpCfgFailedTs": "0",
							  "bgpCfgState": "0",
							  "bmpSrvId1St": "disabled",
							  "bmpSrvId2St": "disabled",
							  "capSuppr4ByteAsn": "disabled",
							  "childAction": "",
							  "connMode": "",
							  "ctrl": "",
							  "curPfxPeers": "0",
							  "dscp": "cs6",
							  "dynRtMap": "",
							  "epe": "disabled",
							  "epePeerSet": "",
							  "gShutOperSt": "disabled",
							  "holdIntvl": "180",
							  "inheritContPeerCtrl": "",
							  "kaIntvl": "60",
							  "logNbrChgs": "none",
							  "lowMemExempt": "disabled",
							  "maxCurPeers": "0",
							  "maxPeerCnt": "0",
							  "maxPfxPeers": "0",
							  "modTs": "2020-10-01T08:54:01.058+00:00",
							  "name": "",
							  "passwdType": "LINE",
							  "password": "",
							  "peerImp": "",
							  "peerType": "fabric-internal",
							  "privateASctrl": "none",
							  "rn": "peer-[10.1.11.3]",
							  "sessionContImp": "",
							  "srcIf": "unspecified",
							  "status": "",
							  "totalPfxPeers": "0",
							  "ttl": "0",
							  "ttlScrtyHops": "0"
							},
							"children": [
							  {
								"bgpPeerEntry": {
								  "attributes": {
									"addr": "10.1.11.3",
									"advCap": "",
									"childAction": "",
									"connAttempts": "558",
									"connDrop": "0",
									"connEst": "0",
									"connIf": "unspecified",
									"fd": "4294967295",
									"flags": "cap-neg,gr-enabled",
									"holdIntvl": "180",
									"kaIntvl": "60",
									"lastFlapTs": "2020-11-05T06:40:39.198+00:00",
									"localIp": "0.0.0.0",
									"localPort": "unspecified",
									"maxConnRetryIntvl": "60",
									"modTs": "never",
									"operSt": "idle",
									"passwdSet": "disabled",
									"peerIdx": "3",
									"prevOperSt": "idle",
									"rcvCap": "",
									"remotePort": "unspecified",
									"rn": "ent-[10.1.11.3]",
									"rtrId": "0.0.0.0",
									"shutStQual": "unspecified",
									"stReason": "none",
									"status": "",
									"type": "ebgp",
									"updateElapsedTs": "2020-11-05T06:38:19.291+00:00"
								  },
								  "children": [
									{
									  "bgpPeerEntryStats": {
										"attributes": {
										  "byteInRecvQ": "0",
										  "byteInSendQ": "0",
										  "byteRcvd": "0",
										  "byteSent": "0",
										  "capRcvd": "0",
										  "capSent": "0",
										  "connectRetryTs": "2020-11-05T06:41:07.268+00:00",
										  "kaRcvd": "0",
										  "kaSent": "0",
										  "kaTs": "2020-11-05T06:42:08.268+00:00",
										  "msgRcvd": "0",
										  "msgSent": "0",
										  "notifRcvd": "0",
										  "notifSent": "0",
										  "openRcvd": "0",
										  "openSent": "0",
										  "rn": "peerstats",
										  "routeRefreshRcvd": "0",
										  "routeRefreshSent": "0",
										  "totalDiscardedAttributes": "0",
										  "updateRcvd": "0",
										  "updateSent": "0"
										}
									  }
									},
									{
									  "bgpPeerAfEntry": {
										"attributes": {
										  "acceptedPaths": "0",
										  "deniedPaths": "0",
										  "firstEorRcvdTs": "2020-11-05T06:39:42.664+00:00",
										  "flags": "gr-sent",
										  "lastEorRcvdTs": "2020-11-05T06:40:39.198+00:00",
										  "memAccPaths": "0",
										  "peerTblVer": "0",
										  "pfxFlushed": "0",
										  "pfxSaved": "0",
										  "pfxSent": "0",
										  "rn": "af-ipv4-ucast",
										  "tblSt": "up",
										  "tblVer": "58",
										  "treatAswithDrawnPaths": "0",
										  "type": "ipv4-ucast",
										  "withDrawnPaths": "0"
										}
									  }
									},
									{
									  "bgpPeerEvents": {
										"attributes": {
										  "lastErrDataRsvd": "",
										  "lastErrDataSent": "",
										  "lastErrLenRsvd": "0",
										  "lastErrLenSent": "0",
										  "lastErrValRsvd": "0",
										  "lastErrValSent": "0",
										  "majErrRstRsvd": "none",
										  "majErrRstSent": "none",
										  "minErrRstRsvd": "none",
										  "minErrRstSent": "none",
										  "rn": "ev",
										  "rstRsvdTs": "2020-11-05T06:38:19.291+00:00",
										  "rstSentTs": "2020-11-05T06:38:19.291+00:00"
										}
									  }
									},
									{
									  "bgpGrSt": {
										"attributes": {
										  "grTs": "1970-01-01T03:00:00.000+00:00",
										  "operSt": "na",
										  "restartIntvl": "0",
										  "rn": "gr"
										}
									  }
									}
								  ]
								}
							  },
							  {
								"bgpPeerAf": {
								  "attributes": {
									"advGwIp": "enabled",
									"advIntvl": "0",
									"advLocalLblRt": "enabled",
									"allowedSelfAsCnt": "0",
									"asOverride": "disabled",
									"capAddlPaths": "",
									"childAction": "",
									"ctrl": "",
									"defOrg": "disabled",
									"defOrgRtMap": "",
									"encapMpls": "disabled",
									"inheritContPeerPolicyCtrl": "",
									"modTs": "2020-10-01T08:54:01.077+00:00",
									"name": "bgp-PeerAf",
									"nhSelfAll": "no",
									"nhThirdparty": "enabled",
									"rewriteRtAsn": "disabled",
									"rn": "af-[ipv4-ucast]",
									"sendComExt": "disabled",
									"sendComStd": "disabled",
									"softReconfigBackup": "none",
									"soo": "unknown:unknown:0:0",
									"status": "",
									"type": "ipv4-ucast",
									"unSupprMap": "",
									"wght": ""
								  }
								}
							  }
							]
						  }
						},
						{
						  "bgpDomAf": {
							"attributes": {
							  "advPip": "disabled",
							  "advSysMac": "disabled",
							  "advertL2vpnEvpn": "disabled",
							  "allocLblAll": "disabled",
							  "allocLblOptB": "enabled",
							  "allocLblRtMap": "",
							  "bestPathCmpltTs": "2020-11-05T06:40:51.200+00:00",
							  "bestPathSigTs": "2020-11-05T06:40:51.191+00:00",
							  "childAction": "",
							  "clReflection": "enabled",
							  "critNhTimeout": "crit",
							  "defInfOrigRd": "unknown:unknown:0:0",
							  "defInfOrigRtt": "unknown:unknown:0:0",
							  "defInfOriginate": "disabled",
							  "defMetric": "",
							  "exportGwIp": "disabled",
							  "igpMetric": "600",
							  "lblAllocMod": "disabled",
							  "maxEcmp": "4",
							  "maxExtEcmp": "1",
							  "maxExtIntEcmp": "1",
							  "maxMxdEcmp": "1",
							  "modTs": "2020-10-01T08:54:01.021+00:00",
							  "nhRtMap": "",
							  "nonCritNhTimeout": "noncrit",
							  "numAggregates": "0",
							  "numNetworks": "0",
							  "numPaths": "9",
							  "numPeers": "1",
							  "numPeersActive": "0",
							  "numRoutes": "9",
							  "retainRttAll": "disabled",
							  "retainRttRtMap": "",
							  "rn": "af-[ipv4-ucast]",
							  "status": "",
							  "supprInactive": "disabled",
							  "tblId": "0x3",
							  "tblMap": "",
							  "tblMapFltr": "disabled",
							  "tblSt": "up",
							  "tblVer": "58",
							  "tmrBstpthDfr": "0",
							  "tmrMax": "0",
							  "type": "ipv4-ucast",
							  "vniEthTag": "disabled",
							  "waitIgpConv": "disabled"
							},
							"children": [
							  {
								"bgpInterLeakP": {
								  "attributes": {
									"asn": "none",
									"childAction": "",
									"inst": "none",
									"modTs": "2020-10-01T08:54:01.003+00:00",
									"proto": "direct",
									"rn": "interleak-direct-interleak-none",
									"rtMap": "RM-REDIST-SUBNET-iAZ",
									"scope": "inter",
									"srv6PrefixType": "unspecified",
									"status": ""
								  }
								}
							  }
							]
						  }
						}
					  ]
					}
				  },
				  {
					"bgpDom": {
					  "attributes": {
						"allocIndex": "0",
						"always": "disabled",
						"bestPathIntvl": "300",
						"bgpCfgFailedBmp": "",
						"bgpCfgFailedTs": "0",
						"bgpCfgState": "0",
						"childAction": "",
						"clusterId": "",
						"firstPeerUpTs": "2020-11-05T06:46:14.399+00:00",
						"holdIntvl": "180",
						"id": "1",
						"kaIntvl": "60",
						"localAsn": "",
						"maxAsLimit": "0",
						"modTs": "2020-10-01T08:53:59.770+00:00",
						"mode": "fabric",
						"name": "default",
						"numEstPeers": "2",
						"numPeers": "2",
						"numPeersPending": "0",
						"operRtrId": "172.16.1.1",
						"operSt": "up",
						"pfxPeerTimeout": "90",
						"pfxPeerWaitTime": "90",
						"reConnIntvl": "60",
						"rn": "dom-default",
						"routerMac": "00:00:00:00:00:00",
						"rtrId": "172.16.1.1",
						"status": "",
						"vnid": "0",
						"vtepIp": "0.0.0.0",
						"vtepVirtIp": "0.0.0.0"
					  },
					  "children": [
						{
						  "bgpRtCtrl": {
							"attributes": {
							  "childAction": "",
							  "enforceFirstAs": "enabled",
							  "fibAccelerate": "disabled",
							  "logNeighborChanges": "enabled",
							  "modTs": "2020-10-01T08:54:00.593+00:00",
							  "rn": "rtctrl",
							  "status": "",
							  "supprRt": "enabled"
							}
						  }
						},
						{
						  "bgpPeerCont": {
							"attributes": {
							  "adminSt": "enabled",
							  "affGrp": "0",
							  "asn": "1",
							  "bmpSrvId1St": "disabled",
							  "bmpSrvId2St": "disabled",
							  "capSuppr4ByteAsn": "disabled",
							  "childAction": "",
							  "connMode": "",
							  "ctrl": "",
							  "desc": "",
							  "dscp": "cs6",
							  "epe": "disabled",
							  "epePeerSet": "",
							  "holdIntvl": "180",
							  "inheritContPeerCtrl": "",
							  "kaIntvl": "60",
							  "logNbrChgs": "none",
							  "lowMemExempt": "disabled",
							  "modTs": "2020-10-01T08:54:00.745+00:00",
							  "name": "SPINE-PEERS",
							  "passwdType": "LINE",
							  "password": "",
							  "peerType": "fabric-internal",
							  "privateASctrl": "none",
							  "rn": "peercont-SPINE-PEERS",
							  "sessionContImp": "",
							  "srcIf": "lo0",
							  "status": "",
							  "ttl": "0",
							  "ttlScrtyHops": "0"
							},
							"children": [
							  {
								"bgpPeerAf": {
								  "attributes": {
									"advGwIp": "enabled",
									"advIntvl": "0",
									"advLocalLblRt": "enabled",
									"allowedSelfAsCnt": "0",
									"asOverride": "disabled",
									"capAddlPaths": "",
									"childAction": "",
									"ctrl": "",
									"defOrg": "disabled",
									"defOrgRtMap": "",
									"encapMpls": "disabled",
									"inheritContPeerPolicyCtrl": "",
									"modTs": "2020-10-01T08:54:00.849+00:00",
									"name": "bgp-PeerAf",
									"nhSelfAll": "no",
									"nhThirdparty": "disabled",
									"rewriteRtAsn": "disabled",
									"rn": "af-[l2vpn-evpn]",
									"sendComExt": "enabled",
									"sendComStd": "enabled",
									"softReconfigBackup": "none",
									"soo": "unknown:unknown:0:0",
									"status": "",
									"type": "l2vpn-evpn",
									"unSupprMap": "",
									"wght": ""
								  }
								}
							  }
							]
						  }
						},
						{
						  "bgpPeer": {
							"attributes": {
							  "activePfxPeers": "0",
							  "addr": "172.16.1.12",
							  "adminSt": "enabled",
							  "affGrp": "0",
							  "asn": "",
							  "bgpCfgFailedBmp": "",
							  "bgpCfgFailedTs": "0",
							  "bgpCfgState": "0",
							  "bmpSrvId1St": "disabled",
							  "bmpSrvId2St": "disabled",
							  "capSuppr4ByteAsn": "disabled",
							  "childAction": "",
							  "connMode": "",
							  "ctrl": "",
							  "curPfxPeers": "0",
							  "dscp": "cs6",
							  "dynRtMap": "",
							  "epe": "disabled",
							  "epePeerSet": "",
							  "gShutOperSt": "disabled",
							  "holdIntvl": "180",
							  "inheritContPeerCtrl": "",
							  "kaIntvl": "60",
							  "logNbrChgs": "none",
							  "lowMemExempt": "disabled",
							  "maxCurPeers": "0",
							  "maxPeerCnt": "0",
							  "maxPfxPeers": "0",
							  "modTs": "2020-10-01T08:54:00.955+00:00",
							  "name": "site1-ag2",
							  "passwdType": "LINE",
							  "password": "",
							  "peerImp": "SPINE-PEERS",
							  "peerType": "fabric-internal",
							  "privateASctrl": "none",
							  "rn": "peer-[172.16.1.12]",
							  "sessionContImp": "",
							  "srcIf": "unspecified",
							  "status": "",
							  "totalPfxPeers": "0",
							  "ttl": "0",
							  "ttlScrtyHops": "0"
							},
							"children": [
							  {
								"bgpPeerEntry": {
								  "attributes": {
									"addr": "172.16.1.12",
									"advCap": "as4,dynamic,dynamic-gr,dynamic-mp,dynamic-old,dynamic-refresh,gr,l2vpn-evpn,refresh,refresh-old",
									"childAction": "",
									"connAttempts": "6",
									"connDrop": "0",
									"connEst": "1",
									"connIf": "unspecified",
									"fd": "79",
									"flags": "cap-neg,gr-enabled",
									"holdIntvl": "180",
									"kaIntvl": "60",
									"lastFlapTs": "2020-11-05T06:46:14.399+00:00",
									"localIp": "172.16.1.1",
									"localPort": "179",
									"maxConnRetryIntvl": "60",
									"modTs": "never",
									"operSt": "established",
									"passwdSet": "disabled",
									"peerIdx": "4",
									"prevOperSt": "open-confirm",
									"rcvCap": "as4,cap,dynamic,dynamic-gr,dynamic-mp,dynamic-old,dynamic-refresh,gr,l2vpn-evpn,refresh,refresh-old",
									"remotePort": "16973",
									"rn": "ent-[172.16.1.12]",
									"rtrId": "172.16.1.12",
									"shutStQual": "admin-up",
									"stReason": "none",
									"status": "",
									"type": "ibgp",
									"updateElapsedTs": "2020-11-05T06:47:28.436+00:00"
								  },
								  "children": [
									{
									  "bgpPeerEntryStats": {
										"attributes": {
										  "byteInRecvQ": "0",
										  "byteInSendQ": "0",
										  "byteRcvd": "401626",
										  "byteSent": "14486",
										  "capRcvd": "2",
										  "capSent": "2",
										  "connectRetryTs": "1970-01-01T03:00:00.000+00:00",
										  "kaRcvd": "727",
										  "kaSent": "728",
										  "kaTs": "2020-11-05T18:52:22.796+00:00",
										  "msgRcvd": "3750",
										  "msgSent": "734",
										  "notifRcvd": "0",
										  "notifSent": "0",
										  "openRcvd": "1",
										  "openSent": "1",
										  "rn": "peerstats",
										  "routeRefreshRcvd": "2",
										  "routeRefreshSent": "0",
										  "totalDiscardedAttributes": "0",
										  "updateRcvd": "3018",
										  "updateSent": "5"
										}
									  }
									},
									{
									  "bgpPeerAfEntry": {
										"attributes": {
										  "acceptedPaths": "3012",
										  "deniedPaths": "8",
										  "firstEorRcvdTs": "2020-11-05T06:39:42.664+00:00",
										  "flags": "first-eor-rcvd,gr-cap,gr-sent",
										  "lastEorRcvdTs": "2020-11-05T06:46:15.410+00:00",
										  "memAccPaths": "361440",
										  "peerTblVer": "61763",
										  "pfxFlushed": "0",
										  "pfxSaved": "0",
										  "pfxSent": "5",
										  "rn": "af-l2vpn-evpn",
										  "tblSt": "up",
										  "tblVer": "61763",
										  "treatAswithDrawnPaths": "0",
										  "type": "l2vpn-evpn",
										  "withDrawnPaths": "0"
										}
									  }
									},
									{
									  "bgpPeerEvents": {
										"attributes": {
										  "lastErrDataRsvd": "",
										  "lastErrDataSent": "",
										  "lastErrLenRsvd": "0",
										  "lastErrLenSent": "0",
										  "lastErrValRsvd": "0",
										  "lastErrValSent": "0",
										  "majErrRstRsvd": "none",
										  "majErrRstSent": "none",
										  "minErrRstRsvd": "none",
										  "minErrRstSent": "none",
										  "rn": "ev",
										  "rstRsvdTs": "2020-11-05T06:38:19.291+00:00",
										  "rstSentTs": "2020-11-05T06:38:19.291+00:00"
										}
									  }
									},
									{
									  "bgpGrSt": {
										"attributes": {
										  "grTs": "1970-01-01T03:00:00.000+00:00",
										  "operSt": "na",
										  "restartIntvl": "120",
										  "rn": "gr"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpPeer": {
							"attributes": {
							  "activePfxPeers": "0",
							  "addr": "172.16.1.11",
							  "adminSt": "enabled",
							  "affGrp": "0",
							  "asn": "",
							  "bgpCfgFailedBmp": "",
							  "bgpCfgFailedTs": "0",
							  "bgpCfgState": "0",
							  "bmpSrvId1St": "disabled",
							  "bmpSrvId2St": "disabled",
							  "capSuppr4ByteAsn": "disabled",
							  "childAction": "",
							  "connMode": "",
							  "ctrl": "",
							  "curPfxPeers": "0",
							  "dscp": "cs6",
							  "dynRtMap": "",
							  "epe": "disabled",
							  "epePeerSet": "",
							  "gShutOperSt": "disabled",
							  "holdIntvl": "180",
							  "inheritContPeerCtrl": "",
							  "kaIntvl": "60",
							  "logNbrChgs": "none",
							  "lowMemExempt": "disabled",
							  "maxCurPeers": "0",
							  "maxPeerCnt": "0",
							  "maxPfxPeers": "0",
							  "modTs": "2020-10-01T08:54:00.909+00:00",
							  "name": "site1-ag1",
							  "passwdType": "LINE",
							  "password": "",
							  "peerImp": "SPINE-PEERS",
							  "peerType": "fabric-internal",
							  "privateASctrl": "none",
							  "rn": "peer-[172.16.1.11]",
							  "sessionContImp": "",
							  "srcIf": "unspecified",
							  "status": "",
							  "totalPfxPeers": "0",
							  "ttl": "0",
							  "ttlScrtyHops": "0"
							},
							"children": [
							  {
								"bgpPeerEntry": {
								  "attributes": {
									"addr": "172.16.1.11",
									"advCap": "as4,dynamic,dynamic-gr,dynamic-mp,dynamic-old,dynamic-refresh,gr,l2vpn-evpn,refresh,refresh-old",
									"childAction": "",
									"connAttempts": "7",
									"connDrop": "0",
									"connEst": "1",
									"connIf": "unspecified",
									"fd": "1478",
									"flags": "cap-neg,gr-enabled",
									"holdIntvl": "180",
									"kaIntvl": "60",
									"lastFlapTs": "2020-11-05T06:46:41.418+00:00",
									"localIp": "172.16.1.1",
									"localPort": "20253",
									"maxConnRetryIntvl": "60",
									"modTs": "never",
									"operSt": "established",
									"passwdSet": "disabled",
									"peerIdx": "3",
									"prevOperSt": "open-confirm",
									"rcvCap": "as4,cap,dynamic,dynamic-gr,dynamic-mp,dynamic-old,dynamic-refresh,gr,l2vpn-evpn,refresh,refresh-old",
									"remotePort": "179",
									"rn": "ent-[172.16.1.11]",
									"rtrId": "172.16.1.11",
									"shutStQual": "admin-up",
									"stReason": "none",
									"status": "",
									"type": "ibgp",
									"updateElapsedTs": "2020-11-05T07:03:21.396+00:00"
								  },
								  "children": [
									{
									  "bgpPeerEntryStats": {
										"attributes": {
										  "byteInRecvQ": "0",
										  "byteInSendQ": "0",
										  "byteRcvd": "401555",
										  "byteSent": "14193",
										  "capRcvd": "2",
										  "capSent": "2",
										  "connectRetryTs": "1970-01-01T03:00:00.000+00:00",
										  "kaRcvd": "726",
										  "kaSent": "728",
										  "kaTs": "2020-11-05T18:52:54.687+00:00",
										  "msgRcvd": "3747",
										  "msgSent": "733",
										  "notifRcvd": "0",
										  "notifSent": "0",
										  "openRcvd": "1",
										  "openSent": "1",
										  "rn": "peerstats",
										  "routeRefreshRcvd": "1",
										  "routeRefreshSent": "0",
										  "totalDiscardedAttributes": "0",
										  "updateRcvd": "3017",
										  "updateSent": "3"
										}
									  }
									},
									{
									  "bgpPeerAfEntry": {
										"attributes": {
										  "acceptedPaths": "3012",
										  "deniedPaths": "8",
										  "firstEorRcvdTs": "2020-11-05T06:39:42.664+00:00",
										  "flags": "first-eor-rcvd,gr-cap,gr-sent",
										  "lastEorRcvdTs": "2020-11-05T06:46:42.499+00:00",
										  "memAccPaths": "361440",
										  "peerTblVer": "61763",
										  "pfxFlushed": "0",
										  "pfxSaved": "0",
										  "pfxSent": "5",
										  "rn": "af-l2vpn-evpn",
										  "tblSt": "up",
										  "tblVer": "61763",
										  "treatAswithDrawnPaths": "0",
										  "type": "l2vpn-evpn",
										  "withDrawnPaths": "0"
										}
									  }
									},
									{
									  "bgpPeerEvents": {
										"attributes": {
										  "lastErrDataRsvd": "",
										  "lastErrDataSent": "",
										  "lastErrLenRsvd": "0",
										  "lastErrLenSent": "0",
										  "lastErrValRsvd": "0",
										  "lastErrValSent": "0",
										  "majErrRstRsvd": "none",
										  "majErrRstSent": "none",
										  "minErrRstRsvd": "none",
										  "minErrRstSent": "none",
										  "rn": "ev",
										  "rstRsvdTs": "2020-11-05T06:38:19.291+00:00",
										  "rstSentTs": "2020-11-05T06:38:19.291+00:00"
										}
									  }
									},
									{
									  "bgpGrSt": {
										"attributes": {
										  "grTs": "1970-01-01T03:00:00.000+00:00",
										  "operSt": "na",
										  "restartIntvl": "120",
										  "rn": "gr"
										}
									  }
									}
								  ]
								}
							  }
							]
						  }
						},
						{
						  "bgpDomAf": {
							"attributes": {
							  "advPip": "disabled",
							  "advSysMac": "disabled",
							  "advertL2vpnEvpn": "disabled",
							  "allocLblAll": "disabled",
							  "allocLblOptB": "enabled",
							  "allocLblRtMap": "",
							  "bestPathCmpltTs": "2020-11-05T06:45:47.671+00:00",
							  "bestPathSigTs": "2020-11-05T06:45:47.671+00:00",
							  "childAction": "",
							  "clReflection": "enabled",
							  "critNhTimeout": "1",
							  "defInfOrigRd": "unknown:unknown:0:0",
							  "defInfOrigRtt": "unknown:unknown:0:0",
							  "defInfOriginate": "disabled",
							  "defMetric": "",
							  "exportGwIp": "disabled",
							  "igpMetric": "600",
							  "lblAllocMod": "disabled",
							  "maxEcmp": "1",
							  "maxExtEcmp": "1",
							  "maxExtIntEcmp": "1",
							  "maxMxdEcmp": "1",
							  "modTs": "2020-10-01T08:54:00.641+00:00",
							  "nhRtMap": "",
							  "nonCritNhTimeout": "1",
							  "numAggregates": "0",
							  "numNetworks": "0",
							  "numPaths": "9041",
							  "numPeers": "2",
							  "numPeersActive": "2",
							  "numRoutes": "6029",
							  "retainRttAll": "disabled",
							  "retainRttRtMap": "",
							  "rn": "af-[l2vpn-evpn]",
							  "status": "",
							  "supprInactive": "disabled",
							  "tblId": "0x1",
							  "tblMap": "",
							  "tblMapFltr": "disabled",
							  "tblSt": "up",
							  "tblVer": "61763",
							  "tmrBstpthDfr": "0",
							  "tmrMax": "0",
							  "type": "l2vpn-evpn",
							  "vniEthTag": "disabled",
							  "waitIgpConv": "disabled"
							}
						  }
						}
					  ]
					}
				  }
				]
			  }
			}
		  ]
		}
	  }
`
	src := make(map[string]interface{})
	err := json.Unmarshal([]byte(Raw), &src)
	if err != nil {
		panic(err)
	}

	var path = []string{"bgpEntity", "bgpInst", "bgpDom", "bgpPeer", "bgpPeerEntry"}
	var pathIndex int
	header := make(map[string]interface{})

	flatten(src, path, pathIndex, header)

}
