package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("content-type", "application/json")
	fmt.Println("start parse file")
	if err := loadHtml("golang23"); err != nil {
		log.Println("loadHTML error:", err)
	}
	// res, err := tushare.Tushare()
	// if err != nil {
	// 	log.Error("error ")
	// }
}

func loadHtml(data string) error {
	parseHTML, err := template.ParseFiles("view/home/home.html")
	if err != nil {
		return err
	}
	if err := parseHTML.Execute(os.Stdout, data); err != nil {
		return err
	}
	return nil
}
