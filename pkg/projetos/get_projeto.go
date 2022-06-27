package projetos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetProjeto(c *gin.Context) {
	id := c.Param("id")

	var projeto BodyGetProjetos
	

	if projeto := h.DB.Raw("select pr.id_projeto, pr.nome_projeto, pr.equipe_id, eq.nome_equipe from projetos as pr inner join equipes as eq on pr.equipe_id = eq.id_equipe where id_projeto = ?", id).Scan(&projeto); projeto.Error != nil {
		c.AbortWithError(http.StatusNotFound, projeto.Error)
		return
	}

	c.JSON(http.StatusOK, &projeto)
}