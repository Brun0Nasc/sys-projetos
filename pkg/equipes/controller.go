package equipes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/equipes")
	routes.GET("/", h.GetEquipes)
	routes.POST("/", h.AddEquipe)
	routes.GET("/:id", h.GetEquipe)
	routes.PUT("/:id", h.UpdateEquipe)
	routes.DELETE("/:id", h.DeleteEquipe)
}