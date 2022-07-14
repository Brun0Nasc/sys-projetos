package pessoas

import (
	"net/http"
	"strconv"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddPessoaRequestBody struct {
	Nome_Pessoa		string  `json:"nome_pessoa"`
	Funcao_Pessoa	string  `json:"funcao_pessoa"`
	EquipeID		string  `json:"equipe_id"`
}

func (h handler) AddPessoa(c *gin.Context) {
	body := AddPessoaRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pessoa models.Pessoa
	if eqId, err := strconv.Atoi(body.EquipeID); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	} else {
		pessoa.Nome_Pessoa = body.Nome_Pessoa
		pessoa.Funcao_Pessoa = body.Funcao_Pessoa
		*pessoa.EquipeID = eqId
	}

	if result := h.DB.Create(&pessoa); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}


	c.JSON(http.StatusCreated, &pessoa)
}