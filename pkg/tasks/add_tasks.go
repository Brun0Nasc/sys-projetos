package tasks

import (
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	Descricao_Task  string 	`json:"descricao_task"`
	PessoaID		int 	`json:"pessoa_id"`
	ProjetoID		int 	`json:"projeto_id"`
}

func (h handler) AddTask(c *gin.Context) {
	/*select count(pe.id_pessoa) from pessoas as pe inner join
equipes as eq on pe.equipe_id = eq.id_equipe 
inner join projetos as pr on pr.equipe_id = eq.id_equipe
where id_pessoa = 4 and pr.status = 'Em desenvolvimento'*/
	
	body := AddTaskRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task
   
	task.Descricao_Task = body.Descricao_Task
	task.PessoaID = body.PessoaID
	task.ProjetoID = body.ProjetoID
	task.Status = "A fazer"

	var check int

	if result := h.DB.Raw("select count(id_projeto) from projetos where id_projeto = ? and status = 'Em desenvolvimento' and equipe_id is not null", task.ProjetoID).Scan(&check);
	result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if(check > 0){
		if result := h.DB.Create(&task); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	
		c.JSON(http.StatusCreated, &task)
	} else {
		c.JSON(http.StatusOK, gin.H{"Message":"Tasks só podem ser cadastradas em projetos que estão 'Em desenvolvimento' e estão em alguma equipe."})
	}
}