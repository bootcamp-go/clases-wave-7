package web

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, response{Data: data})
}
