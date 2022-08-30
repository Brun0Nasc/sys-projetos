package equipes

import (
	"fmt"
	"github.com/Brun0Nasc/sys-projetos/domain/equipes"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"

	"github.com/gin-gonic/gin"
)

func novaEquipe(c *gin.Context) {
	fmt.Println("Tentando adicionar nova equipe")
	req := modelApresentacao.ReqEquipes{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly " + err.Error(),
		})
		return
	}

	equipe.NovaEquipe(c, &req)
}