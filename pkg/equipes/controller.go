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
	//routes.GET("/projetos", h.GetEquipesProjetos)
	routes.POST("/", h.AddEquipe)
	routes.GET("/:id", h.GetEquipe)
	routes.GET("/:id/projetos", h.GetEquipeProjeto)
	routes.GET("/:id/pessoas", h.GetEquipeMembros)
	routes.PUT("/:id", h.UpdateEquipe)
	routes.DELETE("/:id", h.DeleteEquipe)
}

