package projetos

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type BodyGetProjetos struct {
	ID_Projeto 		uint 	`json:"id_projeto"`
	Nome_Projeto 	string 	`json:"nome_projeto"`
	EquipeID 		int 	`json:"equipe_id"`
	Nome_Equipe 	string 	`json:"nome_equipe"`
	Status 			string 	`json:"status"`
	DataInicio 		string 	`json:"data_inicio"`
	DataConclusao 	string 	`default:"" json:"data_conclusao"`
}

func (h handler) GetProjetos(c *gin.Context) {
	var projetos []BodyGetProjetos
	sql := "select pr.id_projeto, pr.nome_projeto, pr.equipe_id, eq.nome_equipe, pr.status, pr.data_inicio, pr.data_conclusao from projetos as pr inner join equipes as eq on pr.equipe_id = eq.id_equipe order by id_projeto"

	if projetos := h.DB.Raw(sql).Scan(&projetos); projetos.Error != nil {
		c.AbortWithError(http.StatusNotFound, projetos.Error)
		return
	}
	
	c.JSON(http.StatusOK, &projetos)
}
