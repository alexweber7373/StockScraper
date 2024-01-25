package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {

	app.GET("/stock/:ticker", getStock)

}
