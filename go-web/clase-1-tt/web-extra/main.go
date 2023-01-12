package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// --------------------------------------------------------
// service
type WebSite struct {
	URL			string	`json:"url"`
	Host 		string	`json:"name"`
	Category 	string	`json:"category"`
	Protected 	bool	`json:"protected"`
}
var websites = []WebSite{
	{URL: "www.google.com", Host: "google", Category: "search", Protected: false},
	{URL: "www.duckgo.com", Host: "duck", Category: "search", Protected: true},
	{URL: "www.mercadolibre.com", Host: "MeLi", Category: "e-commerce", Protected: false},
}
func GetWebsites() []WebSite {
	return websites
}

type Query struct {
	Category 	string	`form:"category"`
	Protected	*bool	`form:"protected"`
}
func GetWebsitesQuery(query Query) []WebSite {
	websites := GetWebsites()

	var filtered []WebSite
	for _, w := range websites {
		if query.Category != "" && (query.Category!=w.Category) {
			continue
		}
		if query.Protected != nil && (*query.Protected!=w.Protected) {
			continue
		}

		filtered = append(filtered, w)
	}
	
	return filtered
}

// --------------------------------------------------------
// pkg
type Response struct {
	Message	string		`json:"message"`
	Data	interface{}	`json:"data"`
}
func New(message string, data interface{}) *Response {
	return &Response{Message: message, Data: data}
}

// --------------------------------------------------------
// server
func main() {
	// router
	r := gin.Default()

	// read
	r.GET("/ping", func(ctx *gin.Context) {
		// request
		// ...

		// process
		// ...

		// response
		// ctx.String(200, "pong")
		ctx.JSON(200, gin.H{"message": "pong", "data": nil})
	})
	r.GET("/next/:id", func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, New("failed to parse int", nil))
			return
		}

		if id < 0 {
			ctx.JSON(http.StatusUnprocessableEntity, New("invalid id. Can not be negative", nil))
			return
		}

		// process
		id++

		// response
		ctx.JSON(http.StatusOK, New("succeed to get next id", id))
	})

	r.GET("/websites", func(ctx *gin.Context) {
		// request
		var query Query
		if err := ctx.BindQuery(&query); err != nil {
			ctx.JSON(http.StatusBadRequest, New(err.Error(), nil))
			return
		}

		// process
		websites := GetWebsitesQuery(query)

		// response
		ctx.JSON(http.StatusOK, New("succeed to get websites", websites))
	})

	// init server
	r.Run(":8080")
}