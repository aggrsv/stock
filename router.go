package main

import (
	"stock/handler"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	router.HandleFunc("/", handler.HomeHandler).Methods("GET")
}
