package pessoas

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetPessoas(c *gin.Context) {
	var pessoas []GetPessoasRequestBody

	if result := h.DB.Raw("select pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, pe.equipe_id, eq.nome_equipe, pe.data_contratacao from pessoas as pe inner join equipes as eq on pe.equipe_id = eq.id_equipe order by pe.id_pessoa").Scan(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &pessoas)
}
