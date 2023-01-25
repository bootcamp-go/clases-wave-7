package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// rutas de la api con los servicios que va a consumir
func main() {
	router := gin.Default()
	
	
	router.GET("/albums",middlewareUno() ,getAlbums)
	router.POST("/albums",middlewareLista(postAlbum)... )
	
	router.Use(middlewareUno())
	router.GET("/album/:id", getIdAlbum)
	router.Run("localhost:8080")
}

// estructura de las request
type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// base de datos del repositorio
var albums = []album{
	{ID: "1", Title: "cualquiera", Artist: "cualquiera", Year: 2002},
}

// servicio que se conecta al repo
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {

	var newAlbum album
	c.BindJSON(&newAlbum)
	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func getIdAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, estructura := range albums {
		if estructura.ID == id {
			c.IndentedJSON(http.StatusOK, estructura)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No se encontro el album"})
}

//ejemplos de middlewares

func middlewareUno() gin.HandlerFunc {
	log.Println("Este es el primer middleware")

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "El token no es el mismo"})
			return
		}
		c.Next()
	}
}

func middlewareDos() gin.HandlerFunc {
	log.Println("Este es el segundo endpoint")
	return func(c *gin.Context) {
		var newAlbum album

		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}

		if newAlbum.Title == "" {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": "Falta el campo titulo"})
			return
		}

		c.Next()
	}

}

func middlewareLista(f gin.HandlerFunc) []gin.HandlerFunc {

	list := []gin.HandlerFunc{
		middlewareUno(),
		middlewareDos(),
	}
	list = append(list, f)
	return list

}
