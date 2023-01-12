package handlers

import (
	"rest/internal/website"

	"github.com/gin-gonic/gin"
)

type Website struct {
	sv website.Service
}
func NewWebsite(sv website.Service) *Website {
	return &Website{sv: sv}
}

func (w *Website) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request

		// process
		ws, err := w.sv.Get()
		if err != nil {
			ctx.JSON(500, nil)
			return
		}

		// response
		ctx.JSON(200, ws)
	}
}
func (w *Website) Create() gin.HandlerFunc {
	type request struct {
		URL       string	`json:"url" validate:"required, url"`
		Host      string	`json:"host" validate:"required"`
		Category  string	`json:"category"`
		Protected bool		`json:"protected"`
	}
	
	return func(ctx *gin.Context) {
		// request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, nil)
			return
		}

		// process
		ws, err := w.sv.Create(req.URL, req.Host, req.Category, req.Protected)
		if err != nil {
			ctx.JSON(500, nil)
			return
		}

		// response
		ctx.JSON(200, ws)
	}
}