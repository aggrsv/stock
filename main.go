package main

import (
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
		mux.Get("/listbyname", h.CodeNameHandler)
		mux.Get("/basicstock", h.HomeHandler)
		mux.Get("/profit", h.ProfitHandler)
		mux.Get("/growth", h.GrowingCompareHandler)
		mux.Get("/integrate", h.IntegrateHandler)
		mux.Get("/prices", h.RealTimePriceHandler)
		mux.Get("/	", h.RealTimePrice2Handler)
	})
}

func main() {
	http.ListenAndServeEtc("stock", mux)
}
