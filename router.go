package main

import (
	"CrawlerBot/Scrapper"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	json := Scrapper.Scrapper
	r.HandleFunc("/api/v1/scrapper", json).Methods(http.MethodGet)

	return r
}
