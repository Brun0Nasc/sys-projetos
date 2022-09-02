package projetos

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", novoProjeto)
	r.GET("/", listarProjetos)
	r.GET("/:id", buscarProjeto)
	r.PUT("/:id", atualizarProjeto)
	r.PUT("/:id/status", atualizarStatus)
	r.DELETE("/:id", deletarProjeto)
}