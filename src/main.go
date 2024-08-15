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
	fmt.Println("Listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
