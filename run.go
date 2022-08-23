package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/volcengine/vefaas-golang-runtime/events"
	"github.com/volcengine/vefaas-golang-runtime/vefaas"
	"github.com/volcengine/vefaas-golang-runtime/vefaascontext"
)

func main() {
	// Start your vefaas function =D.
	go RunFaasCliLoop(map[string]string{
		"http_plb":      "http://dev.douyincloud.gateway.egress.ivolces.com",
		"https_plb":     "https://dev.douyincloud.gateway.egress.ivolces.com",
		"http_openapi":  "http://developer.toutiao.com",
		"https_openapi": "https://developer.toutiao.com",
	})
	select {}
	vefaas.Start(handler)
}

// Define your handler function.
func handler(ctx context.Context, r *events.HTTPRequest) (*events.EventResponse, error) {
	fmt.Printf("received new request: %s %s, request id: %s\n", r.HTTPMethod, r.Path, vefaascontext.RequestIdFromContext(ctx))
	fmt.Printf("debug request: header:%v, body:%s\n", r.Headers, r.Body)
	ret := make(map[string]interface{})
	ret["content"] = "Hello From QA YTR Test For gateway!"
	ret["http_method"] = r.HTTPMethod
	ret["http_path"] = r.Path
	query := make(map[string]interface{})
	header := make(map[string]interface{})
	for k, v := range r.QueryStringParameters {
		query[k] = v
	}
	for k, v := range r.Headers {
		header[k] = v
	}
	ret["http_query"] = query
	ret["http_header"] = header
	ret["http_body"] = string(r.Body)
	retBody, _ := json.Marshal(ret)
	return &events.EventResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: retBody,
	}, nil
}

func RunFaasCliLoop(addrs map[string]string) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			for name, addr := range addrs {
				_ = name
				postFun := func() {
					log.Printf("name:%v, addr:%v\n", name, addr)
					AccessOpenApi(addr)
				}
				postFun()
			}

		}
	}
}

//// Define your handler function.
//func handler(ctx context.Context, r *events.HTTPRequest) (*events.EventResponse, error) {
//	fmt.Printf("received new request: %s %s, request id: %s\n", r.HTTPMethod, r.Path, vefaascontext.RequestIdFromContext(ctx))
//
//	return &events.EventResponse{
//		Headers: map[string]string{
//			"Content-Type": "application/json",
//		},
//		Body: []byte("Hello veFaaS!"),
//	}, nil
//}

func AccessOpenApi(host string) {
	fmt.Printf("start host:%v\n", host)
	//url := "https://developer.toutiao.com/api/apps/qrcode"
	//url := fmt.Sprintf("%v/api/apps/qrcode", host)
	url := fmt.Sprintf("%v/api/apps/qrcode_test", host)
	method := "POST"
	//payload := strings.NewReader(`{"access_token": "0801121846765a5a4d2f6b385a68307237534d43397a667865513d3d","appname": "douyin"}`)
	payloadWithoutToken := strings.NewReader(`{"appname": "douyin"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payloadWithoutToken)

	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	//req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("net err:%v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	log.Printf("raw request:%+v %v %v\n", resp.StatusCode, len(body), resp.Request.RemoteAddr)
	//log.Printf("resp from openapi:%+v\n\n", string(body))
	if len(body) < 100 {
		log.Printf("resp from openapi:%+v\n", string(body))
	}
	fmt.Printf("end\n\n")
}
