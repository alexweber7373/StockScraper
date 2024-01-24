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

	tickerURL := fmt.Sprintf("https://search.brave.com/search?q=%v&source=web", ticker)

	var value float64

	var err error

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("h1.desktop-heading-h2", func(e *colly.HTMLElement) {

		// Modifying text so that a float can be parsed from it
		textString := e.Text
		textString = strings.ReplaceAll(textString, ",", ".")
		textString = textString[:len(textString)-4]
		value, err = strconv.ParseFloat(textString, 64)

	})

	fmt.Println("Value: " + fmt.Sprintf("%f", value))

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	err = c.Visit(tickerURL)

	if err != nil {
		return 0.0, errors.New("Scraper error somehow!")
	}

	return value, err

}
