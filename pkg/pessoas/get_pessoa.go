package pessoas

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

func (h handler) GetPessoa(c *gin.Context) {
	id := c.Param("id")

	var pessoa models.Pessoa

	if result := h.DB.First(&pessoa, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &pessoa)
}