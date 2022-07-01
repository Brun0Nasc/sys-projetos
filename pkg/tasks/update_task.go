package tasks

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Brun0Nasc/sys-projetos/pkg/common/models"
)

type UpdateTaskRequestBody struct {
	ID_Task			uint	`json:"id_task"`
	Descricao_Task  string 	`json:"descricao_task"`
	PessoaID		int 	`json:"pessoa_id"`
	ProjetoID		int 	`json:"projeto_id"`
}

type UpdateStatus struct {
	Status string `json:"status"`
}

func (h handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	body := UpdateTaskRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	task.Descricao_Task = body.Descricao_Task
	task.PessoaID = body.PessoaID
	task.ProjetoID = body.ProjetoID

	h.DB.Save(&task)

	c.JSON(http.StatusOK, &task)
}

func (h handler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	body := UpdateStatus{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	task.Status = body.Status

	if result := h.DB.Raw("update tasks set status = ? where id = ?", task.Status, id).Scan(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotModified, result.Error)
		return
	}

	c.JSON(http.StatusOK, &task)
}