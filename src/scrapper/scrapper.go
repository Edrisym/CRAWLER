package Scrapper

import (
	"CrawlerBot/Excelizing"
	"CrawlerBot/StreamFile"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type ProductDetails struct {
	PriceUnit string `json:"PriceUnit"`
	Date      string `json:"Date"`
}

type Product struct {
	PersianName    string         `json:"persian_name"`
	EnglishName    string         `json:"english_name"`
	BrandOwner     string         `json:"brand_owner"`
	LicenseHolder  string         `json:"license_holder"`
	Price          string         `json:"price"`
	Packaging      string         `json:"packaging"`
	ProductCode    string         `json:"product_code"`
	GenericCode    string         `json:"generic_code"`
	ProductDetails ProductDetails `json:"product_details"`
}

func Scrapper(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	var products []Product

	vars := mux.Vars(r)
	drugName := vars["drugName"]

	urlTemplate := "https://irc.fda.gov.ir/nfi/Search?Term={drugName}&PageNumber={page}&PageSize={size}&Count={count}"
	detailUrl := "https://irc.fda.gov.ir"

	url := strings.Replace(urlTemplate, "{drugName}", drugName, -1)
	url = strings.Replace(url, "{page}", "1", -1)
	url = strings.Replace(url, "{size}", "1000", -1)

	c.OnHTML(".RowSearchSty", func(e *colly.HTMLElement) {
		product := Product{}

		product.PersianName = e.ChildText(".titleSearch-Link-RtlAlter a")
		link := e.ChildAttr(".titleSearch-Link-RtlAlter a", "href")
		detail := detailUrl + link

		c.OnHTML(".row", func(h *colly.HTMLElement) {
			h.Request.Visit(detail)
			s := h.ChildText(".txtSearch:contains('تاریخ اعتبار پروانه') + .txtAlignLTRFa")
			product.ProductDetails.Date = s
			fmt.Println(s)
		})
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
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Visit error: No data is available", err)
		return
	}

	jsonData, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Visit error: No data is available", err)
		return
	}

	StreamFile.TextOut(jsonData, drugName)

	Excelizing.ToExcel(jsonData, drugName)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
