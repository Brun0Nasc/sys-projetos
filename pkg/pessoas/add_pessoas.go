package pessoas

import (
	"net/http"

	"gorm.io/gorm"
    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

type AddPessoaRequestBody struct {
    gorm.Model  // adds ID, created_at etc.
    Nome		string `gorm:"type: varchar(30) not null" json:"nome"`
	Funcao		string `gorm:"type: varchar(20) not null" json:"funcao"`
	Equipe		models.Equipe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"equipe"`
}

func (h handler) AddPessoa(c *gin.Context) {
	body := AddPessoaRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pessoa models.Pessoa
	var equipe models.Equipe = body.Equipe

	

	pessoa.Nome = body.Nome
	pessoa.Funcao = body.Funcao
	pessoa.Equipe = equipe

	if result := h.DB.Create(&pessoa); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &pessoa)
}