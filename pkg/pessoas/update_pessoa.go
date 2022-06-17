package pessoas

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

type UpdatePessoaRequestBody struct {
	Nome		string `json:"nome"`
	Funcao 		string `json:"funcao"`
	Equipe 		models.Equipe `json:"equipe"`
}

func (h handler) UpdatePessoa(c *gin.Context) {
	id := c.Param("id")
	body := UpdatePessoaRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pessoa models.Pessoa

	if result := h.DB.First(&pessoa, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	pessoa.Nome = body.Nome
	pessoa.Funcao = body.Funcao
	pessoa.Equipe = body.Equipe

	h.DB.Save(&pessoa)

	c.JSON(http.StatusOK, &pessoa)
}