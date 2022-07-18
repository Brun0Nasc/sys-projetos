package equipes

import (
	"net/http"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type EquipesGetBody struct {
	ID_Equipe 	uint 				`json:"id_equipe"`
	Nome_Equipe string 				`json:"nome_equipe"`
	Pessoas 	[]models.Pessoa 	`json:"pessoas"`
}

func (h handler) GetEquipes(c *gin.Context) {
	var equipes []models.Equipe
	var eq []EquipesGetBody

	if result := h.DB.Find(&equipes); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for r:=0; r<len(equipes); r++{
		var pessoas []models.Pessoa
		if result := h.DB.Raw("SELECT * FROM pessoas WHERE equipe_id = ?", equipes[r].ID_Equipe).Scan(&pessoas); result.Error != nil{
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		e := &EquipesGetBody{
			ID_Equipe: equipes[r].ID_Equipe,
			Nome_Equipe: equipes[r].Nome_Equipe,
			Pessoas: pessoas,
		}
		eq = append(eq, *e)
	}

	c.JSON(http.StatusOK, &eq)
}