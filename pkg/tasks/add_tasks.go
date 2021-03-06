package tasks

import (
	"net/http"
	"strconv"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	Descricao_Task  string 	`json:"descricao_task"`
	PessoaID		string 	`json:"pessoa_id"`
	ProjetoID		string 	`json:"projeto_id"`
}

func (h handler) AddTask(c *gin.Context) {
	
	body := AddTaskRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task
   
	prId, err := strconv.Atoi(body.ProjetoID)
	peId, err2 := strconv.Atoi(body.PessoaID)

	if err == nil && err2 == nil{
		task.Descricao_Task = body.Descricao_Task
		task.PessoaID = peId
		task.ProjetoID = prId
		task.Status = "A fazer"
	} else{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var equipe int
	if result := h.DB.Raw("select equipe_id from projetos where id_projeto = ?", task.ProjetoID).Scan(&equipe); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var checkE int
	verifica_equipe := "select count(pe.id_pessoa) from pessoas as pe inner join equipes as eq on pe.equipe_id = eq.id_equipe inner join projetos as pr on pr.equipe_id = eq.id_equipe where id_pessoa = ? and pr.status = 'Em desenvolvimento' and pr.equipe_id = ?"
	var checkS int
	verifica_status := "select count(id_projeto) from projetos where id_projeto = ? and status = 'Em desenvolvimento' and equipe_id is not null"

	if result := h.DB.Raw(verifica_equipe, task.PessoaID, equipe).Scan(&checkE); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if(checkE > 0){
		if result := h.DB.Raw(verifica_status, task.ProjetoID).Scan(&checkS); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	
		if(checkS > 0){
			if result := h.DB.Create(&task); result.Error != nil {
				c.AbortWithError(http.StatusNotFound, result.Error)
				return
			}
		
			c.JSON(http.StatusCreated, &task)
		} else {
			c.JSON(http.StatusOK, gin.H{"Message":"Tasks s?? podem ser cadastradas em projetos que est??o 'Em desenvolvimento' e est??o em alguma equipe."})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"Message":"Tasks s?? podem ser atribuidas a pessoas que est??o na equipe respons??vel pelo projeto."})
	}
	
}