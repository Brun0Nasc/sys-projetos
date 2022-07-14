package pessoas

import (
    "net/http"
	"strconv"
    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

type UpdatePessoaRequestBody struct {
	ID_Pessoa		uint   `json:"id_pessoa"`
	Nome_Pessoa		string `json:"nome_pessoa"`
	Funcao_Pessoa	string `json:"funcao_pessoa"`
	EquipeID		*string `json:"equipe_id"`
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
	if body.EquipeID != nil{
		if eqId, err := strconv.Atoi(*body.EquipeID); err == nil {
			c.AbortWithError(http.StatusBadRequest, err)
			pessoa.Nome_Pessoa = body.Nome_Pessoa
			pessoa.Funcao_Pessoa = body.Funcao_Pessoa
			pessoa.EquipeID = eqId
		}
	} else {
		pessoa.Nome_Pessoa = body.Nome_Pessoa
		pessoa.Funcao_Pessoa = body.Funcao_Pessoa
		pessoa.EquipeID = 0
	}
	

	if result := h.DB.Raw("update pessoas set nome_pessoa = ?, funcao_pessoa = ?, equipe_id = ? where id_pessoa = ?", 
	pessoa.Nome_Pessoa, pessoa.Funcao_Pessoa, pessoa.EquipeID, pessoa.ID_Pessoa).Scan(&pessoa); result.Error != nil {
		c.AbortWithError(http.StatusNotModified, result.Error)
		return
	}

	c.JSON(http.StatusOK, &pessoa)
}