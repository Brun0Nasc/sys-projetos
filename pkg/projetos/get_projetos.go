package projetos

import (
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type BodyGetProjetos struct {
	ID_Projeto uint `json:"id_projeto"`
	Nome_Projeto string `json:"nome_projeto"`
	EquipeID int `json:"equipeId"`
}

func (h handler) GetProjetos(c *gin.Context) {
	var projetos []models.Projeto
	body := BodyGetProjetos{}
	var exibe []BodyGetProjetos

	if result := h.DB.Find(&projetos); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for _, p := range projetos {
		body.ID_Projeto = p.ID_Projeto
		body.Nome_Projeto = p.Nome_Projeto
		body.EquipeID = p.EquipeID

		exibe = append(exibe, body)
	}
	
	c.JSON(http.StatusOK, &exibe)
}
