package main

import (
	"example.com/stock-scraper/db"
	"example.com/stock-scraper/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// start the database (might crash entire program due to use of panic)
	db.InitDB()

	// Establish web scraper and gin server
	app := gin.Default()

	routes.RegisterRoutes(app)

	app.Run("127.0.0.1:8080")

}
