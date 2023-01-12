package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// -----------------------------------------------------------------------------------------
// service
type WebSite struct {
	URL			string
	Host 		string
	Category	string
	Protected 	bool
}

var websites = []WebSite{
	{URL: "www.google.com", Host: "google", Category: "search", Protected: false},
	{URL: "www.yahoo.com", Host: "yahoo", Category: "search", Protected: true},
	{URL: "www.mercadolibre.com", Host: "meli", Category: "ecommerce", Protected: false},
}

func Get() []WebSite {
	return websites
}
func GetQuery(category string) []WebSite {
	websites := Get()
	
	var filtered []WebSite
	for _, w := range websites {
		if category != "" && (w.Category != category) {
			continue
		}
		filtered = append(filtered, w)
	}
	
	return filtered
}

// -----------------------------------------------------------------------------------------
// controller
func Ping(ctx *gin.Context) {
	// request
	// ...

	// process
	// ...

	// response
	ctx.String(http.StatusOK, "pong")
} 

// -----------------------------------------------------------------------------------------
// server
func main() {
	sv := gin.Default()

	// router
	sv.GET("/ping", Ping)
	sv.GET("/next/:number", func(ctx *gin.Context) {
		// request
		n, err := strconv.Atoi(ctx.Param("number"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse number", "data": nil})
			return
		}

		// process
		n++

		// response
		ctx.JSON(http.StatusOK, gin.H{"message": "succeed to increment number", "data": n})
	})
	sv.GET("/websites", func(ctx *gin.Context) {
		// request
		queryCategory := ctx.Query("category")

		// process
		websites = GetQuery(queryCategory)

		// response
		ctx.JSON(200, gin.H{"message": "succeed to get all websites", "data": websites})
	})

	sv.Run(":8080")
}