package pessoas

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

	routes := r.Group("/pessoas")
	routes.POST("/", h.AddPessoa)
	routes.GET("/", h.GetPessoas)
	routes.GET("/:id", h.GetPessoa)
	routes.GET("/:id/tasks", h.GetTaskPessoa)
	routes.POST("/:id/tasks", h.AddTaskPessoa)
	routes.PUT("/:id", h.UpdatePessoa)
	routes.PUT("/:id/favoritar", h.FavoritarPessoa)
	routes.DELETE("/:id", h.DeletePessoa)
}