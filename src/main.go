package main

import (
	"example.com/stock-scraper/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Establish web scraper and gin server
	app := gin.Default()

	routes.RegisterRoutes(app)

	app.Run("127.0.0.1:8080")

	// if err != nil {
	// 	fmt.Println("Error reading input. Exiting.")
	// 	return
	// }

}
