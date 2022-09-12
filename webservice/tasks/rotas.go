package tasks

import (

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.POST("/", novaTask)
	r.POST("/:id/comentarios", novoComentario)
	r.GET("/", listarTasks)
	r.GET("/:id", buscarTask)
	r.GET("/:id/comentarios", listarComentarios)
	r.PUT("/:id", atualizarTask)
	r.PUT("/:id/status", atualizarStatus)
	r.PUT("/:id/comentarios/:idc", atualizarComentario)
	r.DELETE("/:id", deletarTask)
	r.DELETE("/:id/comentarios/:idc", deletarComentario)
}
