package projetos

import (
	"net/http"
	"time"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddProjetoRequestBody struct {
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
	dt := time.Now()

	projeto.Nome_Projeto = body.Nome_Projeto
	projeto.EquipeID = body.EquipeID
	projeto.DataInicio = dt.Format("02-01-2006")

	if result := h.DB.Create(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &projeto)
}