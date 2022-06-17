package equipes

import (
	"net/http"

    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

func (h handler) DeleteEquipe(c *gin.Context) {
	id := c.Param("id")

	var equipe models.Equipe

	if result := h.DB.First(&equipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&equipe)

	c.Status(http.StatusOK)
}