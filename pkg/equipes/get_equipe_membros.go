package equipes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type membros struct {
	Nome_Equipe 	string `json:"nome_equipe"`
	Nome_Pessoa 	string `json:"nome_pessoa"`
	Funcao_Pessoa 	string `json:"funcao_pessoa"`
}

func (h handler) GetEquipeMembros (c *gin.Context) {
	id := c.Param("id")

	var membros []membros

	if membros := h.DB.Raw("select eq.nome_equipe, pe.nome_pessoa, pe.funcao_pessoa from equipes as eq inner join pessoas as pe on eq.id_equipe = pe.equipe_id where eq.id_equipe = ?", id).Scan(&membros); membros.Error != nil {
		c.AbortWithError(http.StatusNotFound, membros.Error)
		return
	}

	c.JSON(http.StatusOK, &membros)
}