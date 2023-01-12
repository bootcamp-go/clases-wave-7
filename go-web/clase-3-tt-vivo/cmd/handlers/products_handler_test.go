package handlers

import (
	"bcgow7/internal/products"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerForProductsHandlerTest(storage products.Storage) *gin.Engine {
	// Create a router.
	router := gin.Default()

	// Create the handler.
	productsHandler := NewProductsHandler(storage)

	// Define the routes.
	productsRoutes := router.Group("/products")
	{
		productsRoutes.GET("/", productsHandler.GetAll())
		productsRoutes.POST("/", productsHandler.Store())
		productsRoutes.PUT("/:id", productsHandler.Update())
		productsRoutes.PATCH("/:id", productsHandler.PartialUpdate())
		productsRoutes.DELETE("/:id", productsHandler.Delete())
	}

	// Let's go!
	return router
}

func Test_GetAll(t *testing.T) {
	// Arrange.
	storage := &products.LocalStorage{
		Products: []products.Product{
			{
				ID:    "p123",
				Name:  "Galletitas Boreo 500g",
				Price: 120.0,
			},
		},
	}
	server := createServerForProductsHandlerTest(storage)

	request := httptest.NewRequest(http.MethodGet, "/products/", nil)
	response := httptest.NewRecorder()

	expectedResponse := []products.Product{
		{
			ID:    "p123",
			Name:  "Galletitas Boreo 500g",
			Price: 120.0,
		},
	}

	// Act.
	server.ServeHTTP(response, request)

	var obtainedBody []products.Product
	errOnBodyUnmarshall := json.Unmarshal(response.Body.Bytes(), &obtainedBody)

	// Assert.
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Nil(t, errOnBodyUnmarshall, "received body is not a valid json")
	assert.Equal(t, expectedResponse, obtainedBody)
}

func Test_GetAll_InternalServerError(t *testing.T) {
	// Arrange.
	storage := &products.DummyStorage{
		ErrOnGetAllMethod: errors.New("i'm an error :("),
	}
	server := createServerForProductsHandlerTest(storage)

	request := httptest.NewRequest(http.MethodGet, "/products/", nil)
	response := httptest.NewRecorder()

	// Act.
	server.ServeHTTP(response, request)

	obtainedBody, errOnObtainingBody := io.ReadAll(response.Body)

	// Assert.
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Nil(t, errOnObtainingBody)
	assert.True(t, len(obtainedBody) == 0)
}

func Test_Store(t *testing.T) {
	// Arrange.
	storage := &products.DummyStorage{
		ProductOnStoreMethod: &products.Product{
			ID:    "p123",
			Name:  "Galletitas Boreo 500g",
			Price: 120.0,
		},
		ErrOnStoreMethod: nil,
	}
	server := createServerForProductsHandlerTest(storage)

	request := httptest.NewRequest(http.MethodPost, "/products/", bytes.NewBuffer(
		[]byte(`
			{
				"name": "Galletitas Boreo 500g",
				"price": 120.0
			}
		`),
	))

	response := httptest.NewRecorder()
	expectedResponse := products.Product{
		ID:    "p123",
		Name:  "Galletitas Boreo 500g",
		Price: 120.0,
	}

	// Act.
	server.ServeHTTP(response, request)

	var obtainedBody products.Product
	errOnBodyUnmarshall := json.Unmarshal(response.Body.Bytes(), &obtainedBody)

	// Assert.
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Nil(t, errOnBodyUnmarshall, "received body is not a valid json")
	assert.Equal(t, response.Header().Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, expectedResponse, obtainedBody)
}

func Test_Store_ErrBadRequest(t *testing.T) {
	// Arrange.
	storage := &products.LocalStorage{}
	server := createServerForProductsHandlerTest(storage)

	request := httptest.NewRequest(http.MethodPost, "/products/", bytes.NewBuffer(
		[]byte("{+8a7987sdasd2"),
	))
	response := httptest.NewRecorder()

	// Act.
	server.ServeHTTP(response, request)

	// Assert.
	assert.Equal(t, http.StatusBadRequest, response.Code)
}
