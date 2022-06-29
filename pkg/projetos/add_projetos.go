package projetos

import (
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddProjetoRequestBody struct {
	ID_Projeto	 uint	`json:"id_projeto"`
	Nome_Projeto string `json:"nome_projeto"`
	EquipeID 	 int 	`json:"equipe_id"`
}

func (h handler) AddProjeto(c *gin.Context) {
	body := AddProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var projeto models.Projeto

	projeto.ID_Projeto = body.ID_Projeto
	projeto.Nome_Projeto = body.Nome_Projeto
	projeto.EquipeID = body.EquipeID

	if result := h.DB.Create(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &projeto)
}