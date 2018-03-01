package main

import (
	nh "net/http"
	"stock/comm/http"
	h "stock/handler"
)

const (
	prefix = "/stock/v1"
)

var (
	mux = http.Route(prefix)
)

func init() {
	mux.Group("", func(mux *http.Mux) {
		mux.Get("/basicstock", h.HomeHandler)
		mux.Get("/profit", h.ProfitHandler)
	})
}

func main() {
	//http.ListenAndServeEtc("stock", mux)
	nh.ListenAndServe(":8080", mux)
}
