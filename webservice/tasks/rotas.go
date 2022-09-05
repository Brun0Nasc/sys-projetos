package tasks

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", novaTask)
	r.GET("/", listarTasks)
	r.GET("/:id", buscarTask)
	r.PUT("/:id", atualizarTask)
	r.PUT("/:id/status", atualizarStatus)
	r.DELETE("/:id", deletarTask)
}