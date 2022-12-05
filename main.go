package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	fmt.Fprintf(w, "NowTime: %s", time.Now().Format("2006-01-02 15:04:05"))
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "TZ:%s", os.Getenv("TZ"))
	fmt.Fprintf(w, "NowTime: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

func gatewayTest(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(401)
	fmt.Fprintf(w, "gateway_test, req_url:%v, ts:%s", req.URL, time.Now().Format("2006-01-02 15:04:05"))
}

func gatewayNetworkTest(w http.ResponseWriter, req *http.Request) {
	protocol := "http"
	protocols := req.URL.Query()["protocol"]
	if len(protocols) > 0 && protocols[0] == "https"{
		protocol = "https"
	}
	urls := req.URL.Query()["url"]
	if len(urls) == 0{
		fmt.Fprintf(w, "non url param")
		return
	}

	fmt.Fprintf(w, AccessOpenApiUrl(protocol, urls[0]))
}


func main() {

	//go RunFaasCliLoop(map[string]string{
	//	"http_plb":      "http://dev.douyincloud.gateway.egress.ivolces.com",
	//	"https_plb":     "https://dev.douyincloud.gateway.egress.ivolces.com",
	//	"http_openapi":  "http://developer.toutiao.com",
	//	"https_openapi": "https://developer.toutiao.com",
	//	"http_openapi2":  "http://open.douyin.com",
	//	"https_openapi2": "https://open.douyin.com",
	//})

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/v1/ping", ping)
	http.HandleFunc("/gateway_test", gatewayTest)
	http.HandleFunc("/gateway_network_test", gatewayNetworkTest)
	http.HandleFunc("/gateway_ws_push", gatewayWsPush)
	http.HandleFunc("/gateway_ws_handle", gatewayWsHandle)

	http.ListenAndServe(":8000", nil)
}