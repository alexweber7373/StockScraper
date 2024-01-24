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

	c.OnHTML("span.QuoteStrip-lastPrice", func(e *colly.HTMLElement) {

		// Modifying text so that a float can be parsed from it
		fmt.Println("e.Text: " + e.Text)
		textString := e.Text
		// if e.Text != "" {
		// 	textString = textString[1:]
		// }
		fmt.Println(textString)
		value, err = strconv.ParseFloat(textString, 64)

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
