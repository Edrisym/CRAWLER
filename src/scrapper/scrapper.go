package Scrapper

import (
	"CrawlerBot/Excelizing"
	"CrawlerBot/Product"
	"CrawlerBot/ProductDetail"
	"CrawlerBot/StreamFile"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func Scrapper(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	var products []Product.Product

	vars := mux.Vars(r)
	drugName := vars["drugName"]

	urlTemplate := "https://irc.fda.gov.ir/nfi/Search?Term={drugName}&PageNumber={page}&PageSize={size}&Count={count}"
	detailUrl := "https://irc.fda.gov.ir"

	url := strings.Replace(urlTemplate, "{drugName}", drugName, -1)
	url = strings.Replace(url, "{page}", "1", -1)
	url = strings.Replace(url, "{size}", "1000", -1)

	product := Product.Product{}
	c.OnHTML(".RowSearchSty", func(e *colly.HTMLElement) {

		product.PersianName = e.ChildText(".titleSearch-Link-RtlAlter a")
		link := e.ChildAttr(".titleSearch-Link-RtlAlter a", "href")
		detail := detailUrl + link

		err := c.Visit(detail)
		if err != nil {
			fmt.Println("Error visiting detail page:", err)
		}

		product.EnglishName = e.ChildText(".titleSearch-Link-ltrAlter a")

		product.BrandOwner = e.ChildText(".txtSearch:contains('صاحب برند') + .txtSearch1")

		product.LicenseHolder = e.ChildText(".txtSearch:contains('صاحب پروانه') + .txtSearch1")

		price := e.ChildText(".priceTxt")
		priceUnit := "ریال"
		product.Price = fmt.Sprintf("%s %s", strings.TrimSpace(price), strings.TrimSpace(priceUnit))

		product.Packaging = e.ChildText(".txtSearch:contains('بسته بندی') + bdo")
		product.ProductCode = e.ChildText(".txtSearch:contains('کد فرآورده') + .txtSearch1")
		product.GenericCode = e.ChildText(".txtSearch:contains('کد ژنریک') + .txtSearch1")

		products = append(products, product)
	})

	c.OnHTML(".row", func(e *colly.HTMLElement) {
		date := e.ChildText(".txtSearch:contains('تاریخ اعتبار پروانه') + .txtAlignLTRFa")
		if date != "" {
			pd := ProductDetail.ProductDetails{
				LicenceDate: date,
			}
			product.ProductDetails = pd
		}
	})

	err := c.Visit(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Visit error: No data is available", err)
		return
	}

	jsonData, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Visit error: No data is available", err)
		return
	}

	//marshaledData := json.Unmarshal(jsonData, &products)
	//fmt.Println(marshaledData)
	//if marshaledData == nil {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusNotFound)
	//	w.Write([]byte("No data is available"))
	//	return
	//}

	StreamFile.TextOut(jsonData, drugName)

	Excelizing.ToExcel(jsonData, drugName)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
