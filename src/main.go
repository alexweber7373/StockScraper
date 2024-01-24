package main

import (
	"fmt"

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
	}

	stockValue, err := scraper.Scrape(c, ticker)

	if err != nil {
		fmt.Println("Program exiting!")
		return
	}

	fmt.Printf("STOCK VALUE FOR %v: $%.2f\n", ticker, stockValue)

}
