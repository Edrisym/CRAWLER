package main

import (
	"CrawlerBot/Scrapper"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/scrapper/{drugName}", Scrapper.Scrapper).Methods(http.MethodGet)
	fmt.Println("Listening on http://localhost:8080/api/v1/scrapper/")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Visit error: No data is available", err)
		return
	}
}
