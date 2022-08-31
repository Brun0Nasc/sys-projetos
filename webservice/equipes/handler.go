package equipes

import (
	"fmt"
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/domain/equipes"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"

	"github.com/gin-gonic/gin"
)

func novaEquipe(c *gin.Context) {
	fmt.Println("Tentando adicionar nova equipe")
	req := modelApresentacao.ReqEquipe{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly " + err.Error(),
		})
		return
	}

	equipe.NovaEquipe(&req, c)
	c.JSON(http.StatusCreated, gin.H{"OK":"Registro adicionado com Sucesso!"})
}

func listarEquipes(c *gin.Context) {
	fmt.Println("Tentando listar equipes") 
	c.JSON(http.StatusOK, equipe.ListarEquipes(c))
}