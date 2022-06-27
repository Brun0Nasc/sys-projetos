package pessoas

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetPessoa struct {
	ID_Pessoa       uint   `json:"id_pessoa"`
    Nome_Pessoa		string `json:"nome_pessoa"`
	Funcao_Pessoa	string `json:"funcao_pessoa"`
	EquipeID		int    `json:"equipeId"`
	Nome_Equipe		string `json:"nome_equipe"`
}

func (h handler) GetPessoa(c *gin.Context) {
	id := c.Param("id")

	var pessoa GetPessoa

	if result := h.DB.Raw("SELECT pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, pe.equipeId, eq.nome_equipe FROM pessoas AS pe INNER JOIN equipe AS eq on pe.equipeId = eq.id_equipe WHERE pe.id_pessoa = ?", id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &pessoa)
}