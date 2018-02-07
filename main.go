package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//tushare.Tushare()
	fmt.Println("listen on server => 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Print("start server error :", err)
	}
}
