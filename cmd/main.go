package main

import (
	"BetterContent/internal/handlers"
	"BetterContent/internal/validators"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	urlValidator := validators.NewURLValidator()
	contentHandler := handlers.NewContentHandler(urlValidator)

	router.POST("/addContentLink", func(context *gin.Context) {
		contentHandler.HandleAddContentLink(context)
	})
	err := router.Run()

	if err != nil {
		return
	}
}
