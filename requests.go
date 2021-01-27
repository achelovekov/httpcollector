package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type NXAPILoginRequestData struct {
	AaaUser struct {
		Attributes struct {
			Name string `json:"name"`
			Pwd  string `json:"pwd"`
		} `json:"attributes"`
	} `json:"aaaUser"`
}

type NXAPILoginResponseData struct {
	Imdata []struct {
		AaaLogin struct {
			Attributes struct {
				Token string `json:"token"`
			} `json:"attributes"`
		} `json:"aaaLogin"`
	} `json:"imdata"`
}

func NXAPILogin(user string, password string) string {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}

	loginRequest := NXAPILoginRequestData{}
	loginRequest.AaaUser.Attributes.Name = user
	loginRequest.AaaUser.Attributes.Pwd = password
	requestBody, err := json.MarshalIndent(loginRequest, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Post("https://10.62.130.40/api/mo/aaaLogin.json", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	var loginResponse NXAPILoginResponseData

	err = json.Unmarshal([]byte(body), &loginResponse)

	token := "APIC-cookie=" + loginResponse.Imdata[0].AaaLogin.Attributes.Token

	return token
}

func NXAPIbaseRequest(token string, url []string) []byte {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}

	for {
		for _, v := range url {
			req, _ := http.NewRequest("GET", v, io.Reader(nil))
			req.Header.Set("Cookie", token)
			resp, err := client.Do(req)
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("error = %s \n", err)
			}
			src := make(map[string]interface{})
			err = json.Unmarshal(data, &src)

			srcJSON, err := json.MarshalIndent(src, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Printf("MarshalIndent function output %s\n", string(srcJSON))
			time.Sleep(5 * time.Second)
		}
	}
}

func main() {
	token := NXAPILogin("admin", "cisco!123")
	NXAPIbaseRequest(token, []string{"https://10.62.130.40/api/mo/sys/bgp.json?rsp-subtree=full",
		"https://10.62.130.40/api/mo/sys/ospf.json?rsp-subtree=full",
		"https://10.62.130.40/api/mo/sys/intf.json?rsp-subtree=full",
		"https://10.62.130.40/api/mo/sys/procsys.json?rsp-subtree=full",
		"https://10.62.130.40/api/mo/sys/bd.json?rsp-subtree=full",
		"https://10.62.130.40/api/mo/sys/eps.json?rsp-subtree=full",
		"https://10.62.130.40/api/mo/sys/ch.json?rsp-subtree=full"})
}
