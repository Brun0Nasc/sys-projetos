package projetos

import (
	//"database/sql"
	"fmt"
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/domain/projetos"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/projetos/model"

	"github.com/gin-gonic/gin"
)

func novoProjeto(c *gin.Context) {
	fmt.Println("Tentando adicionar novo projeto")

	req := modelApresentacao.ReqProjeto{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
		})
		return
	}

	if res, err := projetos.NovoProjeto(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	} else if err == nil && res == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Impossível cadastrar um projeto em uma equipe sem membros.",
		})
	} else {
		c.JSON(http.StatusCreated, res)
	}
}