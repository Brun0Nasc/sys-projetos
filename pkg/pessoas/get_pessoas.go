package pessoas

import (
	"net/http"
	"time"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type GetPessoasBody struct{
	ID_Pessoa 		uint 			`json:"id_pessoa"`
	Nome_Pessoa 	string 			`json:"nome_pessoa"`
	Funcao_Pessoa 	string 			`json:"funcao_pessoa"`
	DataContratacao time.Time		`json:"data_contratacao"`
	EquipeID 		*int 			`json:"equipe_id"`
	Equipe 			*models.Equipe 	`json:"equipe"`
}

func (h handler) GetPessoas(c *gin.Context) {
	var pessoas []GetPessoasBody


	if result := h.DB.Raw("select * from pessoas").Scan(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for i := 0; i < len(pessoas); i++{
		if pessoas[i].EquipeID == nil{
			pessoas[i].Equipe = nil
		} else{
			var equipe models.Equipe
			if result := h.DB.First(&equipe, pessoas[i].EquipeID); result.Error != nil {
				c.AbortWithError(http.StatusNotFound, result.Error)
				return
			}
			pessoas[i].Equipe = &equipe
		}
	}

	c.JSON(http.StatusOK, &pessoas)
}
