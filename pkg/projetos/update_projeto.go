package projetos

import (
    "net/http"
	"time"
    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"strconv"
)

type UpdateProjetoRequestBody struct {
	ID_Projeto	 uint	`json:"id_projeto"`
	Nome_Projeto string `json:"nome_projeto"`
	EquipeID 	 string	`json:"equipe_id"`
}

type UpdateStatusRequestBody struct {
	Status string `json:"status"`
}

func (h handler) UpdateProjeto(c *gin.Context) {
	id := c.Param("id")
	body := UpdateProjetoRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var projeto models.Projeto

	if result := h.DB.First(&projeto, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if eqId, err := strconv.Atoi(body.EquipeID); err == nil {
		projeto.Nome_Projeto = body.Nome_Projeto
		projeto.EquipeID = eqId
	}

	if result := h.DB.Raw(`update projetos set nome_projeto = ?, equipe_id = ? where id_projeto = ?`, 
	projeto.Nome_Projeto, projeto.EquipeID, projeto.ID_Projeto).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotModified, result.Error)
		return
	}

	c.JSON(http.StatusOK, &projeto)
}

func (h handler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	body := UpdateStatusRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var projeto models.Projeto

	if result := h.DB.First(&projeto, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	
	projeto.Status = body.Status
	check := 0

	if projeto.Status == "Em desenvolvimento" {
		sql := "select count(id_projeto) from projetos where status = 'Em desenvolvimento' and equipe_id = ?"
		if result := h.DB.Raw(sql, projeto.EquipeID).Scan(&check); result.Error != nil {
			c.AbortWithError(http.StatusNotModified, result.Error)
			return
		}
		if check == 0 {
			if result := h.DB.Raw(`update projetos set status = ? where id_projeto = ?`, 
			projeto.Status, id).Scan(&projeto); result.Error != nil {
				c.AbortWithError(http.StatusNotModified, result.Error)
				return
			}
	
			c.JSON(http.StatusOK, &projeto)
		} else{
			c.JSON(http.StatusOK, gin.H{"Message":"Essa equipe j?? tem um projeto Em desenvolvimento"})
		}
	} else if projeto.Status == "Conclu??do" {
		dt := time.Now()
		sql := "UPDATE projetos SET status = ?, data_conclusao = ? WHERE id_projeto = ?"
		if result := h.DB.Raw(sql, projeto.Status, dt, id).Scan(&projeto); result.Error != nil {
			c.AbortWithError(http.StatusNotModified, result.Error)
			return
		}

		c.JSON(http.StatusOK, &projeto)
	}

	
	
}