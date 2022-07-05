package pessoas

import (
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetPessoas(c *gin.Context) {
	var pessoas []models.Pessoa

	if result := h.DB.Raw("select * from pessoas order by id_pessoa").Scan(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &pessoas)
}
