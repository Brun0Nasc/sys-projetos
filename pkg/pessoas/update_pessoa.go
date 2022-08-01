package pessoas

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

func (h handler) UpdatePessoa(c *gin.Context) {
	id := c.Param("id")
	body := AddPessoaRequestBody{}

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
	if body.EquipeID != 0{
			pessoa.Nome_Pessoa = body.Nome_Pessoa
			pessoa.Funcao_Pessoa = body.Funcao_Pessoa
			pessoa.EquipeID = body.EquipeID

		if result := h.DB.Raw("update pessoas set nome_pessoa = ?, funcao_pessoa = ?, equipe_id = ? where id_pessoa = ?", 
		pessoa.Nome_Pessoa, pessoa.Funcao_Pessoa, pessoa.EquipeID, pessoa.ID_Pessoa).Scan(&pessoa).Find(&pessoa); result.Error != nil {
			c.AbortWithError(http.StatusNotModified, result.Error)
			return
		}
		
	} else {
		pessoa.Nome_Pessoa = body.Nome_Pessoa
		pessoa.Funcao_Pessoa = body.Funcao_Pessoa
		
		if result := h.DB.Raw("update pessoas set nome_pessoa = ?, funcao_pessoa = ?, equipe_id = null where id_pessoa = ?", 
		pessoa.Nome_Pessoa, pessoa.Funcao_Pessoa, pessoa.ID_Pessoa).Scan(&pessoa).Find(&pessoa); result.Error != nil {
			c.AbortWithError(http.StatusNotModified, result.Error)
			return
		}
	}

	c.JSON(http.StatusOK, &pessoa)
}

func (h handler) FavoritarPessoa(c *gin.Context) {
	id := c.Param("id")
	body := models.Pessoa{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pessoa models.Pessoa

	if result := h.DB.Model(&pessoa).Where("id_pessoa = ?", id).Update("favoritar", body.Favoritar); result.Error != nil {
		c.AbortWithError(http.StatusNotModified, result.Error)
		return
	}

	c.JSON(http.StatusOK, &pessoa)
}