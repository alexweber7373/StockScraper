package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// c.OnHTML("a", func(e *colly.HTMLElement) {
	// 	// printing all URLs associated with the a links in the page
	// 	fmt.Printf("%v", e.Attr("href"))
	// })

	c.OnHTML("h1.desktop-heading-h2", func(e *colly.HTMLElement) {
		fmt.Printf("%v", e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	c.Visit("https://search.brave.com/search?q=NFLX&source=web")

}
