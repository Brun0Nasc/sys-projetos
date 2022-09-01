package pessoas

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", novaPessoa)
	r.GET("/", listarPessoas)
	r.GET("/:id", buscarPessoa)
}