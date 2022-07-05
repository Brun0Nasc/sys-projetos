package pessoas

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetPessoaID struct {
	ID_Pessoa 		uint 	`json:"id_pessoa"`
	Nome_Pessoa 	string 	`json:"nome_pessoa"`
	Funcao_Pessoa	string 	`json:"funcao_pessoa"`
	EquipeID 		int 	`json:"equipe_id"`
	Nome_Equipe 	string 	`json:"nome_equipe"`
	DataContratacao string 	`json:"data_contratacao"`
}

func (h handler) GetPessoa(c *gin.Context) {
	id := c.Param("id")

	var pessoa GetPessoaID

	if pessoa := h.DB.Raw("select pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, pe.equipe_id, eq.nome_equipe, pe.data_contratacao from pessoas as pe inner join equipes as eq on pe.equipe_id = eq.id_equipe where id_pessoa = ?", id).Scan(&pessoa); pessoa.Error != nil {
		c.AbortWithError(http.StatusNotFound, pessoa.Error)
		return
	}

	c.JSON(http.StatusOK, &pessoa)
}