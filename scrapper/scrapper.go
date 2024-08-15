package Scrapper

import (
	"CrawlerBot/StreaFile"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/http"
	"strings"
)

// Define the Product struct with JSON tags for output
type Product struct {
	PersianName   string `json:"persian_name"`
	EnglishName   string `json:"english_name"`
	BrandOwner    string `json:"brand_owner"`
	LicenseHolder string `json:"license_holder"`
	Price         string `json:"price"`
	Packaging     string `json:"packaging"`
	ProductCode   string `json:"product_code"`
	GenericCode   string `json:"generic_code"`
}

// Scrapper function to handle HTTP requests and respond with scraped data
func Scrapper(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	var products []Product

	// Base URL with placeholders
	//urlTemplate := "https://irc.fda.gov.ir/nfi/Search?Term={term}&PageNumber={page}&PageSize={size}&Count={count}"

	//url := strings.Replace(urlTemplate, "{term}", "baricitinib", -1)
	//url = strings.Replace(url, "{page}", "1", -1)
	//url = strings.Replace(url, "{size}", "1", -1)
	//url = strings.Replace(url, "{count}", "0", -1)

	url := "https://irc.fda.gov.ir/nfi/Search?Term=baricitinib"
	//&PageNumber=1&PageSize=1&Count=0
	// Set up a callback function for when a visited HTML page is parsed
	c.OnHTML(".RowSearchSty", func(e *colly.HTMLElement) {
		product := Product{}

		product.PersianName = e.ChildText(".titleSearch-Link-RtlAlter a")

		product.EnglishName = e.ChildText(".titleSearch-Link-ltrAlter a")

		product.BrandOwner = e.ChildText(".txtSearch:contains('صاحب برند') + .txtSearch1")

		product.LicenseHolder = e.ChildText(".txtSearch:contains('صاحب پروانه') + .txtSearch1")

		price := e.ChildText(".priceTxt") // This extracts "14,893,200"
		priceUnit := "ریال"
		product.Price = fmt.Sprintf("%s %s", strings.TrimSpace(price), strings.TrimSpace(priceUnit))

		product.Packaging = e.ChildText(".txtSearch:contains('بسته بندی') + bdo")
		product.ProductCode = e.ChildText(".txtSearch:contains('کد فرآورده') + .txtSearch1")
		product.GenericCode = e.ChildText(".txtSearch:contains('کد ژنریک') + .txtSearch1")
		products = append(products, product)
	})

	err := c.Visit(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	StreaFile.PdfOut(jsonData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
