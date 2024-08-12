package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Run the tasks
	var res string
	err := chromedp.Run(ctx,
		// Navigate to the page
		chromedp.Navigate("https://irc.fda.gov.ir/nfi/Search"),
		// Wait for the search input to be visible
		chromedp.WaitVisible(`#Term`),
		// Fill the search input field
		chromedp.SendKeys(`#Term`, "baricitinib"),
		// Click the submit button
		chromedp.Click(`#submit`),
		// Wait for navigation or page load
		chromedp.WaitVisible(`#result-id`),
		// Extract the result
		chromedp.Text(`#result-id`, &res),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print the result
	fmt.Println("Search result:", res)
}
