package projetos

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

	routes := r.Group("/projetos")
	routes.POST("/", h.AddProjeto)
	routes.GET("/", h.GetProjetos)
	routes.GET("/:id", h.GetProjeto)
	routes.PUT("/:id", h.UpdateProjeto)
	routes.DELETE("/:id", h.DeleteProjeto)
}