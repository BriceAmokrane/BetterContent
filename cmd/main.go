package main

import (
	"BetterContent/internal/handlers"
	"BetterContent/internal/validators"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	router := gin.Default()

	urlValidator := validators.NewURLValidator()
	contentHandler := handlers.NewContentHandler(urlValidator)

	router.POST("/addContentLink", func(context *gin.Context) {
		contentHandler.HandleAddContentLink(context)
	})
	err = router.Run()

	if err != nil {
		return
	}
}
