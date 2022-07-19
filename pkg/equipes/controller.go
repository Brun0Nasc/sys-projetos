package equipes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	//Registro das rotas de equipes
	
	h := &handler{
		DB: db,
	}

	//
	routes := r.Group("/equipes")
	routes.GET("/", h.GetEquipes)
	routes.GET("/:id", h.GetEquipe)
	routes.POST("/", h.AddEquipe)
	routes.GET("/:id/projetos", h.GetEquipeProjeto)
	routes.PUT("/:id", h.UpdateEquipe)
	routes.DELETE("/:id", h.DeleteEquipe)
}

