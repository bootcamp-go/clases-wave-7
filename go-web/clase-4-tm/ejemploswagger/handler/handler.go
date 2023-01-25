package handler

import (
	"ejemploswagger/repo"
	"ejemploswagger/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//funcion que trae todos los albums
//comentario propio que no afecta a la documentacion
// @Summary List Albums
// @Tags Albums
// @Description Gets all albums without filter
// @Produce json
// @Param token header string true "token"
// @Sucess 200 {object}	web.reponse
// @Failure 401 {object} web.response
// @Router /albums [GET]
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		albums := service.GetAlbums()

		if albums == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no hay ningun album"})
			return
		}

		c.IndentedJSON(http.StatusOK, albums)
	}
}

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var album repo.Album
		err := c.BindJSON(&album)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no se pudo enviar la request"})
			return
		}
		exito := service.Create(album)
		c.IndentedJSON(http.StatusOK, exito)
	}
}

