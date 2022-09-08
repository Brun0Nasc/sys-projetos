package tasks

import (
	"database/sql"
	"fmt"

	"github.com/Brun0Nasc/sys-projetos/domain/tasks"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/tasks/model"

	"github.com/gin-gonic/gin"
)

func pegaJSON(c *gin.Context) *modelApresentacao.ReqTask {
	req := modelApresentacao.ReqTask{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
		})
		return nil
	}
	return &req
}

func novaTask(c *gin.Context) {
	fmt.Println("Tentando adicionar uma nova task.")
	req := pegaJSON(c)
	res, err := tasks.NovaTask(req)

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

func listarTasks(c *gin.Context) {
	fmt.Println("Tentando listar tasks")
	
	res, err := tasks.ListarTasks()
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Nenhum cadastro encontrado", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	}

	c.JSON(200, res)
}

func buscarTask(c *gin.Context) {
	fmt.Println("Tentando buscar task")
	id := c.Param("id")

	res, err := tasks.BuscarTask(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Nenhum cadastro encontrado", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	}

	c.JSON(200, res)
}

func atualizarTask(c *gin.Context) {
	fmt.Println("Tentando atualizar task")
	id := c.Param("id")
	req := pegaJSON(c)

	res, err := tasks.AtualizarTask(id, req)
	if err != nil {
		if err == sql.ErrNoRows{
			c.JSON(200, gin.H{"message":"Esse cadastro não existe", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	}

	c.JSON(200, res)
}

func atualizarStatus(c *gin.Context) {
	fmt.Println("Tentando atualizar status")
	id := c.Param("id")
	req := pegaJSON(c)

	res, err := tasks.AtualizarStatus(id, req)
	if err != nil {
		if err == sql.ErrNoRows{
			c.JSON(200, gin.H{"message":"Esse cadastro não existe", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	}

	c.JSON(200, res)
}

func deletarTask(c *gin.Context) {
	fmt.Println("Tentando deletar task")
	id := c.Param("id")

	err := tasks.DeletarTask(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Esse cadastro não existe"})
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message":"Cadastro deletado com sucesso"})
}