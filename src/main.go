package main

import (
	"fmt"
	"strings"

	"example.com/stock-scraper/scraper"
	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector()

	var ticker string

	fmt.Println("Please enter your ticker symbol you would like the stock for: ")

	_, err := fmt.Scan(&ticker)

	if err != nil {
		fmt.Println("Error reading input. Exiting.")
		return
	}

	stockValue, err := scraper.Scrape(c, ticker)

	if stockValue == 0.0 {
		fmt.Println("Error: Scraper did not return a value.")
		return
	}

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	fmt.Printf("STOCK VALUE FOR %v: $%.2f\n", strings.ToUpper(ticker), stockValue)

}
