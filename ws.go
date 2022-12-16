package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// newProxy takes target host and creates a reverse proxy
func newProxy(targetHost string) (*httputil.ReverseProxy, error) {
	u, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	return proxy, nil
}

func gatewayWsPush(w http.ResponseWriter, req *http.Request) {
	mockHost := req.Header.Get("MockHost")
	if len(mockHost) > 0{
		req.Host = mockHost
	}
	mockPath := req.Header.Get("MockPath")
	if len(mockPath) > 0{
		req.URL.Path = mockPath
	}
	proxy, err := newProxy("http://developer.toutiao.com")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	proxy.ServeHTTP(w, req)
}

func gatewayWsHandle(w http.ResponseWriter, req *http.Request) {
	logid := req.Header.Get("X-TT-LOGID")
	method := req.Method
	header := req.Header.Clone()
	headerBytes, _ := json.Marshal(header)
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	fmt.Printf("logid:%v, ReqMethod:%v, headers:%+v, body:%s\n", logid, method, header, bodyBytes)
	w.Header().Set("X-TT-LOGID", logid)
	w.Header().Set("ReqMethod", method)
	w.Header().Set("ReqHeader", string(headerBytes))
	w.Write(bodyBytes)
}
