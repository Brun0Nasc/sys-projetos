package projetos

import (
	"net/http"
	"time"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type Projeto struct {
	ID_Projeto 		uint 			`gorm:"primary_key" json:"id_projeto"`
	Nome_Projeto 	string 			`json:"nome_projeto"`
	EquipeID 		int 			`json:"equipe_id"`
	Status			string			`json:"status"`
	DataInicio 		*time.Time 		`json:"data_inicio"`
	DataConclusao 	*time.Time		`json:"data_conclusao"`
}

type BodyGetProjetos struct {
	ID_Projeto 		uint 			`json:"id_projeto"`
	Nome_Projeto 	string 			`json:"nome_projeto"`
	Status 			string 			`json:"status"`
	DataInicio 		*time.Time 		`json:"data_inicio"`
	DataConclusao 	*time.Time		`json:"data_conclusao"`
	Equipe			models.Equipe	`json:"equipe"`
	Tasks			*[]models.Task 	`json:"tasks"`
}

func (h handler) GetProjetos(c *gin.Context) {
	var projetos []Projeto
	var pr []BodyGetProjetos

	if result := h.DB.Find(&projetos); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for i := 0; i < len(projetos); i++ {
		var equipe models.Equipe
		var tasks []models.Task

		if result := h.DB.Where("id_equipe = ?", projetos[i].EquipeID).Find(&equipe); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		if result := h.DB.Where("projeto_id = ?", projetos[i].ID_Projeto).Find(&tasks); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		p := BodyGetProjetos {
			ID_Projeto: projetos[i].ID_Projeto,
			Nome_Projeto: projetos[i].Nome_Projeto,
			Status: projetos[i].Status,
			DataInicio: projetos[i].DataInicio,
			DataConclusao: projetos[i].DataConclusao,
			Equipe: equipe,
			Tasks: &tasks,
		}

		pr = append(pr, p)
	}
	
	c.JSON(http.StatusOK, &pr)
}
