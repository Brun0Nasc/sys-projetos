package pessoas

import (
	"net/http"

    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

type AddPessoaRequestBody struct {
    Nome_Pessoa		string 	`json:"nome_pessoa"`
	Funcao_Pessoa	string 	`json:"funcao_pessoa"`
	EquipeID		int 	`json:"equipeId"`
}

func (h handler) AddPessoa(c *gin.Context) {
	body := AddPessoaRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pessoa models.Pessoa

	pessoa.Nome_Pessoa = body.Nome_Pessoa
	pessoa.Funcao_Pessoa = body.Funcao_Pessoa
	pessoa.EquipeID = body.EquipeID

	sql := "INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES(?,?,?)"

	if result := h.DB.Raw(sql, pessoa.Nome_Pessoa, pessoa.Funcao_Pessoa, pessoa.EquipeID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &pessoa)
}