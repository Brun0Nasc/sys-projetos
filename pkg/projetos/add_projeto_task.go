package projetos

import (
	"net/http"
	"strconv"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	ID_Task			uint	`json:"id_task"`
	Descricao_Task  string 	`json:"descricao_task"`
	PessoaID		int 	`json:"pessoa_id"`
}

func (h handler) AddTaskEquipe(c *gin.Context) {
	body := AddTaskRequestBody{}
	id := c.Param("id")

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task
	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	task.ID_Task = body.ID_Task
	task.Descricao_Task = body.Descricao_Task
	task.PessoaID = body.PessoaID
	task.ProjetoID = i
	task.Status = "A fazer"

	if result := h.DB.Create(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &task)
}