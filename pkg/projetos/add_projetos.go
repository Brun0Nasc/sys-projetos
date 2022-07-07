package projetos

import (
	"net/http"
	"time"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddProjetoRequestBody struct {
	Nome_Projeto string `json:"nome_projeto"`
	EquipeID 	 int 	`json:"equipe_id"`
}

func (h handler) AddProjeto(c *gin.Context) {
	body := AddProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var projeto models.Projeto
	dt := time.Now()

	projeto.Nome_Projeto = body.Nome_Projeto
	projeto.EquipeID = body.EquipeID
	projeto.Status = "Em planejamento"
	projeto.DataInicio = dt.Format("02-01-2006")

	var check int
	sql1 := "select count(id_equipe) from equipes where id_equipe = ?"
	sql := "select count(id_pessoa) from pessoas where equipe_id = ?"

	if result := h.DB.Raw(sql1, projeto.EquipeID).Scan(&check); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if check == 0{
		if result := h.DB.Create(&projeto); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		c.JSON(http.StatusCreated, &projeto)
	} else{
		if result := h.DB.Raw(sql, projeto.EquipeID).Scan(&check); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	
		if(check > 0){
			if result := h.DB.Create(&projeto); result.Error != nil {
				c.AbortWithError(http.StatusNotFound, result.Error)
				return
			}
			c.JSON(http.StatusCreated, &projeto)
		} else{
			c.JSON(http.StatusOK, gin.H{"Message":"Projetos só podem ser cadastrados em equipes que possuam membros."})
		}
	}
}