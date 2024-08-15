package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()
	url := "https://irc.fda.gov.ir/nfi/Search?Term=baricitinib&PageNumber=1&PageSize=1&Count=0"

	// Set up a callback function for when a visited HTML page is parsed
	c.OnHTML(".RowSearchSty", func(e *colly.HTMLElement) {
		// Extract image source
		imageSrc := e.ChildAttr(".BoxImgSearch img", "src")
		fmt.Println("Image Source:", imageSrc)

		// Extract RTL title
		rtlTitle := e.ChildText(".titleSearch-Link-RtlAlter a")
		rtlLink := e.ChildAttr(".titleSearch-Link-RtlAlter a", "href")
		fmt.Println("Title (RTL):", rtlTitle)
		fmt.Println("Title (RTL Link):", rtlLink)

		// Extract LTR title
		ltrTitle := e.ChildText(".titleSearch-Link-ltrAlter a")
		ltrLink := e.ChildAttr(".titleSearch-Link-ltrAlter a", "href")
		fmt.Println("Title (LTR):", ltrTitle)
		fmt.Println("Title (LTR Link):", ltrLink)

		// Extract Brand Owner
		brandOwner := e.ChildText(".txtSearch:contains('صاحب برند') + .txtSearch1")
		fmt.Println("Brand Owner:", brandOwner)

		// Extract License Holder
		licenseHolder := e.ChildText(".txtSearch:contains('صاحب پروانه') + .txtSearch1")
		fmt.Println("License Holder:", licenseHolder)

		// Extract Price
		price := e.ChildText(".priceTxt")
		priceUnit := e.ChildText(".txtAlignLTR")
		fmt.Println("Price:", price)
		fmt.Println("Price Unit:", priceUnit)

		// Extract Packaging
		packaging := e.ChildText(".txtSearch:contains('بسته بندی') + bdo")
		fmt.Println("Packaging:", packaging)

		// Extract Product Code
		productCode := e.ChildText(".txtSearch:contains('کد فرآورده') + .txtSearch1")
		fmt.Println("Product Code:", productCode)

		// Extract Generic Code
		genericCode := e.ChildText(".txtSearch:contains('کد ژنریک') + .txtSearch1")
		fmt.Println("Generic Code:", genericCode)

		fmt.Println("--------------------------------------------------")
	})

	// Start the scraping process
	err := c.Visit(url) // Replace with the actual URL
	if err != nil {
		log.Fatal(err)
	}
}
