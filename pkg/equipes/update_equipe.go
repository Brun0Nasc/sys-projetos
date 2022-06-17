package equipes

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

type UpdateEquipeRequestBody struct {
	Nome string `json:"nome"`
}

func (h handler) UpdateEquipe(c *gin.Context) {
	id := c.Param("id")
	body := UpdateEquipeRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var equipe models.Equipe

	if result := h.DB.First(&equipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	equipe.Nome = body.Nome

	h.DB.Save(&equipe)

	c.JSON(http.StatusOK, &equipe)
}