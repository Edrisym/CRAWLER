package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
)

func main() {
	c := colly.NewCollector()

	type Result struct {
		Title     string
		PriceText string
		//Description string
	}

	var results []Result

	c.OnHTML(".row", func(e *colly.HTMLElement) {
		fmt.Println("Page found:", e.Request.URL)
		title := e.ChildText(".txtSearch1")   // Adjust selector as needed
		priceText := e.ChildText(".priceTxt") // Adjust selector as needed

		results = append(results, Result{
			Title:     title,
			PriceText: priceText,
			//Description: description,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit("https://irc.fda.gov.ir/nfi/Search?Term=baricitinib&PageNumber=1&PageSize=10&Count=0")
	if err != nil {
		log.Fatal(err)
	}

	// Print out the results
	for _, result := range results {
		fmt.Printf("Title: %s\nDescription: %s\n\n", result.Title, result.PriceText)
	}

}
