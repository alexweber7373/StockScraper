package routes

import (
	"net/http"

	"example.com/stock-scraper/scraper"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func getStock(context *gin.Context) {

	c := colly.NewCollector()

	// get ticker out of GET url
	ticker := context.Param("ticker")

	stockValue, err := scraper.Scrape(c, ticker)

	if stockValue == 0.0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find price for given ticker."})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Scraper threw an error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Stock found!", "stock_price": stockValue})

}
