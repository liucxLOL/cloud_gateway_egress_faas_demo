package main

import (
	"net/http"
)

func gatewayDnsTest(w http.ResponseWriter, req *http.Request) {
	mockHost := req.Header.Get("MockHost")
	if len(mockHost) > 0{
		req.Host = mockHost
	}
	mockPath := req.Header.Get("MockPath")
	if len(mockPath) > 0{
		req.URL.Path = mockPath
	}
	proxy, err := newProxy("http://"+req.Host)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	proxy.ServeHTTP(w, req)
}
