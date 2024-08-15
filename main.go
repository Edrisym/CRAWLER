package main

import (
	"CrawlerBot/Scrapper"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/scrapper", Scrapper.Scrapper).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
