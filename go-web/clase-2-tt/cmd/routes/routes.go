package routes

import (
	"rest/cmd/handlers"
	"rest/internal/domain"
	"rest/internal/website"

	"github.com/gin-gonic/gin"
)

type Router struct {
	db *[]domain.Website
	en *gin.Engine
}
func NewRouter(en *gin.Engine, db *[]domain.Website) *Router {
	return &Router{en: en, db: db}
}

func (r *Router) SetRoutes() {
	r.SetWebsite()
}
// website
func (r *Router) SetWebsite() {
	// instances
	rp := website.NewRepository(r.db, 3)
	sv := website.NewService(rp)
	h := handlers.NewWebsite(sv)

	ws := r.en.Group("/websites")
	ws.GET("", h.Get())
	ws.POST("", h.Create())
}