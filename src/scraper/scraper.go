package scraper

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func Scrape(c *colly.Collector, ticker string) (float64, error) {

	tickerURL := fmt.Sprintf("https://www.cnbc.com/quotes/%v", strings.ToLower(ticker))

	var value float64

	var err error

	c.AllowURLRevisit = false

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Chrome/91.0.4472.124")
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	valueSet := false

	c.OnHTML("span.QuoteStrip-lastPrice", func(e *colly.HTMLElement) {

		// make sure only the first value is set
		if !valueSet {

			// Modifying text so that a float can be parsed from it
			textString := e.Text
			fmt.Println("Stock Value: " + textString)
			value, err = strconv.ParseFloat(textString, 64)

			valueSet = true

		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	err = c.Visit(tickerURL)

	if err != nil {
		return 0.0, errors.New("Scraper error somehow!")
	}

	return value, err

}
