package routes

import (
	"net/http"

	"example.com/stock-scraper/models"
	"example.com/stock-scraper/scraper"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func getStock(context *gin.Context) {

	var stock models.Stock

	c := colly.NewCollector()

	// get ticker out of GET url
	ticker := context.Param("ticker")
	stock.Ticker = ticker

	stockValue, err := scraper.Scrape(c, ticker)

	if stockValue == 0.0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find price for given ticker."})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Scraper threw an error"})
		return
	}

	stock.Value = stockValue

	err = stock.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save stock in database."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Stock found!", "stock_price": stockValue})

}
