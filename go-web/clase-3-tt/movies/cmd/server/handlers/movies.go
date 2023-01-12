package handlers

import (
	"errors"
	"net/http"
	"os"
	"rest/internal/movies"
	"rest/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller
type Movies struct {
	sv movies.Service
	// auth service
	// auth Authenticator
}

func NewMovies(sv movies.Service) *Movies {
	return &Movies{sv: sv}
}

// read
func (mv *Movies) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, response.Err("invalid token"))
			return
		}

		// request

		// process
		ms, err := mv.sv.Get()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Err("failed to get movies"))
			return
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok("succeed to get movies", ms))
	}
}
func (mv *Movies) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, response.Err("invalid token"))
			return
		}

		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err("invalid id"))
			return
		}

		// process
		m, err := mv.sv.GetByID(id)
		if err != nil {
			if errors.Is(err, movies.ErrNotFound) {
				ctx.JSON(http.StatusNotFound, response.Err("movie not found"))
				return
			}
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok("succeed to get movie", m))
	}
}