package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"rest/cmd/handlers"
)

func main() {
	// server
	sv := gin.Default()

	// router
	websites := sv.Group("/websites")
	websites.GET("", handlers.Get)
	websites.POST("", handlers.Create)

	// start
	if err := sv.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}