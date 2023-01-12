package main

import (
	"bcgow7/cmd/handlers"
	"bcgow7/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear un router en gin.
	router := gin.Default()

	// Crear un storage local.
	var storage products.Storage = &products.LocalStorage{
		Products: []products.Product{
			{
				ID:    "p123",
				Name:  "Galletitas Boreo 500g",
				Price: 120.0,
			},
		},
	}

	// Crear un handler de productos.
	productsHandler := handlers.NewProductsHandler(storage)

	// Crear un grupo de rutas para /products.
	productsRoutes := router.Group("/products")
	{
		productsRoutes.GET("/", productsHandler.GetAll())
		productsRoutes.POST("/", productsHandler.Store())
		productsRoutes.PUT("/:id", productsHandler.Update())
		productsRoutes.PATCH("/:id", productsHandler.PartialUpdate())
		productsRoutes.DELETE("/:id", productsHandler.Delete())
	}

	// Iniciar el servidor.
	router.Run(":8080")
}
