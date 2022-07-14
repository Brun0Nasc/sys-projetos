package pessoas

import (
	"net/http"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)


func (h handler) GetPessoa(c *gin.Context) {
	id := c.Param("id")

	var pessoas GetPessoasBody

	if result := h.DB.Raw("select * from pessoas where id_pessoa = ?", id).Scan(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	if pessoas.EquipeID == nil{
		pessoas.Equipe = nil
	} else{
		var equipe models.Equipe
		if result := h.DB.First(&equipe, pessoas.EquipeID); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		pessoas.Equipe = &equipe
	}

	c.JSON(http.StatusOK, &pessoas)
}