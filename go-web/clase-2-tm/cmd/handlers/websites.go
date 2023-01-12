package handlers

import (
	"errors"
	"net/http"
	"rest/pkg/response"
	"rest/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const SECRET_KEY = "ABC"

var (
	ErrUnauthorized = errors.New("error: invalid token")
)

func Get(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != SECRET_KEY {
		ctx.JSON(http.StatusUnauthorized, response.Err(ErrUnauthorized))
		return
	}

	// request

	// process
	websites := services.Get()

	// response
	ctx.JSON(http.StatusOK, response.Ok("succeed to get websites", websites))
}

type request struct {
	URL       string `json:"url" validate:"required,url"`
	Host      string `json:"host" validate:"required"`
	Category  string `json:"category"`
	Protected bool   `json:"protected"`
}
func Create(ctx *gin.Context) {
	// request
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err(err))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
		return
	}

	// process
	website, err := services.Create(req.URL, req.Host, req.Category, req.Protected)
	if err != nil {
		if errors.Is(err, services.ErrAlreadyExist) {
			ctx.JSON(http.StatusConflict, response.Err(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	// response
	ctx.JSON(http.StatusCreated, response.Ok("suceed to create website", website))
}