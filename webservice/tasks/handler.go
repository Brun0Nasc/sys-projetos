package tasks

import (
	"fmt"

	"github.com/Brun0Nasc/sys-projetos/domain/tasks"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/tasks/model"

	"github.com/gin-gonic/gin"
)

func novaTask(c *gin.Context) {
	fmt.Println("Tentando adicionar uma nova task.")
	req := modelApresentacao.ReqTask{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
		})
		return
	}

	res, err := tasks.NovaTask(&req)

	if err != nil {
		c.JSON(400, gin.H{
			"err":err.Error(),
		})
		return
	}

	if res == nil && err == nil {
		c.JSON(400, gin.H{
			"message":"Tasks só podem ser atribuídas a membros da equipe responsável pelo projeto.",
		})
		return
	}

	c.JSON(201, res)
}