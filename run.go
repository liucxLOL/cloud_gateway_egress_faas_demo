package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func RunFaasCliLoop(addrs map[string]string) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			for name, addr := range addrs {
				_ = name
				postFun := func() {
					AccessOpenApi(addr)
				}
				postFun()
			}

		}
	}
}

func AccessOpenApi(host string) {
	fmt.Printf("start host:%v\n", host)
	//url := "https://developer.toutiao.com/api/apps/qrcode"
	//url := fmt.Sprintf("%v/api/apps/qrcode", host)
	url := fmt.Sprintf("%v/api/apps/qrcode", host)
	method := "POST"
	//payload := strings.NewReader(`{"access_token": "0801121846765a5a4d2f6b385a68307237534d43397a667865513d3d","appname": "douyin"}`)
	payloadWithoutToken := strings.NewReader(`{"appname": "douyin"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payloadWithoutToken)

	if err != nil {
		fmt.Printf("request failed. url:%v, err:%v\n", url, err)
		return
	}
	//req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("net failed. url:%v, err:%v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("body failed. url:%v, err:%v\n", url, err)
		return
	}

	log.Printf("url:%v, raw request:%+v %v %v\n", url, resp.StatusCode, len(body), resp.Request.RemoteAddr)
	//log.Printf("resp from openapi:%+v\n\n", string(body))
	if len(body) < 100 {
		log.Printf("url:%v, resp from openapi:%+v\n", url, string(body))
	} else {
		log.Printf("url:%v, len resp from openapi:%+v\n", url, len(body))
	}
}

func AccessOpenApiUrl(protocol, hostUrl string) string {
	if protocol == "http" {
		hostUrl = "http://"+hostUrl
	} else if protocol == "https" {
		hostUrl = "https://"+hostUrl
	} else {
		str := fmt.Sprintf("invalid protocol. protocol:%v", protocol)
		fmt.Printf("%v\n", str)
		return str
	}

	url := hostUrl
	method := "POST"
	//payload := strings.NewReader(`{"access_token": "0801121846765a5a4d2f6b385a68307237534d43397a667865513d3d","appname": "douyin"}`)
	payloadWithoutToken := strings.NewReader(`{"appname": "douyin"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payloadWithoutToken)

	if err != nil {
		str := fmt.Sprintf("request failed. url:%v, err:%v", url, err)
		fmt.Printf("%v\n", str)
		return str
	}
	//req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		str := fmt.Sprintf("net failed. url:%v, err:%v", url, err)
		fmt.Printf("%v\n", str)
		return str
	}
	defer resp.Body.Close()

	logid := resp.Header.Get("X-Tt-Logid")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		str := fmt.Sprintf("body failed. logid:%v, url:%v, err:%v", logid, url, err)
		fmt.Printf("%v\n", str)
		return str
	}

	if len(body) < 100 {
		str := fmt.Sprintf("logid:%v, url:%v, http_code:%v, body:%v", logid, url, resp.StatusCode, string(body))
		fmt.Printf("%v\n", str)
		return str
	}
	str := fmt.Sprintf("logid:%v, url:%v, http_code:%v, len_body:%v", logid, url, resp.StatusCode, len(body))
	fmt.Printf("%v\n", str)
	return str

}
