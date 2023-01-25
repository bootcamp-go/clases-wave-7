package handlers

import (
	"bcgow7/internal/products"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestForProductUpdate struct {
	// Name is the name of the product.
	Name string

	// Price is the price of the product.
	Price float64
}

type RequestForProductPartialUpdate struct {
	// Name is the name of the product.
	Name *string

	// Price is the price of the product.
	Price *float64
}

// ProductsHandler is a HTTP Handler for /products resource.
type ProductsHandler struct {
	Storage products.Storage
}

// NewProductsHandler creates a new ProductsHandler instance.
func NewProductsHandler(storage products.Storage) *ProductsHandler {
	return &ProductsHandler{
		Storage: storage,
	}
}

func (handler *ProductsHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener el producto del body.
		var product products.Product
		if err := ctx.ShouldBindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		// Guardar el producto.
		err := handler.Storage.Store(&product)
		if err != nil {
			// TODO: Manejar el error.
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		// Retornar el producto creado.
		ctx.JSON(http.StatusCreated, product)
	}
}

func (handler *ProductsHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener todos los productos.
		products, err := handler.Storage.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		// Retornar los productos.
		ctx.JSON(http.StatusOK, products)
	}
}

func (handler *ProductsHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener el identificador del producto a actualizar.
		id := ctx.Param("id")

		// Obtener el producto del body.
		var updateRequest RequestForProductUpdate
		if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		var updatedProduct = products.Product{
			ID:    id,
			Name:  updateRequest.Name,
			Price: updateRequest.Price,
		}

		// Actualizar el producto.
		err := handler.Storage.Update(id, &updatedProduct)
		if err != nil {
			switch err {
			case products.ErrProductNotFound:
				ctx.JSON(http.StatusNotFound, nil)
			default:
				ctx.JSON(http.StatusInternalServerError, nil)
			}
			return
		}

		// Retornar el producto actualizado.
		ctx.JSON(http.StatusOK, updatedProduct)
	}
}

func (handler *ProductsHandler) PartialUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener el identificador del producto a actualizar.
		id := ctx.Param("id")

		// Obtener el producto a actualizar.
		product, err := handler.Storage.GetByID(id)
		if err != nil {
			switch err {
			case products.ErrProductNotFound:
				ctx.JSON(http.StatusNotFound, nil)
			default:
				ctx.JSON(http.StatusInternalServerError, nil)
			}
			return
		}

		err = json.NewDecoder(ctx.Request.Body).Decode(&product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		err = handler.Storage.Update(id, product)
		if err != nil {
			switch err {
			case products.ErrProductNotFound:
				ctx.JSON(http.StatusNotFound, nil)
			default:
				ctx.JSON(http.StatusInternalServerError, nil)
			}
			return
		}

		// Retornar el producto actualizado.
		ctx.JSON(http.StatusOK, product)
	}
}

func (handler *ProductsHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener el identificador del producto a eliminar.
		id := ctx.Param("id")

		// Eliminar el producto.
		err := handler.Storage.Delete(id)
		if err != nil {
			switch err {
			case products.ErrProductNotFound:
				ctx.JSON(http.StatusNotFound, nil)
			default:
				ctx.JSON(http.StatusInternalServerError, nil)
			}
			return
		}

		// Retornar el producto eliminado.
		ctx.JSON(http.StatusNoContent, nil)
	}
}
