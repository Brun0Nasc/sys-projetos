package tasks

import (
	"github.com/Brun0Nasc/sys-projetos/pkg/common/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.Use(middlewares.CORSMiddleware())

	h := &handler{
		DB: db,
	}

	routes := r.Group("/tasks", middlewares.Auth())
	routes.POST("/", h.AddTask)
	routes.GET("/", h.GetTasks)
	routes.GET("/:id", h.GetTask)
	routes.PUT("/:id", h.UpdateTask)
	routes.PUT("/:id/status", h.UpdateStatus)
	routes.DELETE("/:id", h.DeleteTask)
}