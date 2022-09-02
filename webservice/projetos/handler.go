package projetos

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/domain/projetos"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/projetos/model"

	"github.com/gin-gonic/gin"
)

func novoProjeto(c *gin.Context) {
	fmt.Println("Tentando adicionar novo projeto")

	req := modelApresentacao.ReqProjeto{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
		})
		return
	}

	if res, err := projetos.NovoProjeto(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	} else if err == nil && res == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Impossível cadastrar um projeto em uma equipe sem membros.",
		})
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func listarProjetos(c *gin.Context) {
	fmt.Println("Tentando listar projetos")

	res, err := projetos.ListarProjetos()
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Não há cadastros registrados", "err":err.Error()})	
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	}
	c.JSON(200, res)
}

func buscarProjeto(c *gin.Context){
	fmt.Println("Tentando buscar projeto")
	id := c.Param("id")

	res, err := projetos.BuscarProjeto(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Esse cadastro não existe", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	}
	c.JSON(200, res)
}

func atualizarProjeto(c *gin.Context){
	fmt.Println("Tentando atualizar projeto")
	id := c.Param("id")
	req := modelApresentacao.ReqProjeto{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message":"Could not update. Parameters were not passed correctly", "err":err.Error()})
	}

	res, err := projetos.AtualizarProjeto(id, &req)
	if err != nil {
		c.JSON(304, gin.H{"err":err.Error()})
		return
	}
	c.JSON(200, res)
}

func atualizarStatus(c *gin.Context){
	fmt.Println("Tentando atualizar status do projeto")
	id := c.Param("id")
	req := modelApresentacao.ReqProjeto{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message":"Could not update. Parameters were not passed correctly", "err":err.Error()})
	}

	res, err := projetos.AtualizarStatus(id, &req)
	if err != nil {
		if err == sql.ErrNoRows{
			c.JSON(200, gin.H{"message":"Esse cadastro não existe", "err":err.Error()})
		} else {
			c.JSON(304, gin.H{"err":err.Error()})
		}
		return
	}
	c.JSON(200, res)
}

func deletarProjeto(c *gin.Context){
	fmt.Println("Tentando deletar projeto")
	id := c.Param("id")

	err := projetos.DeletarProjeto(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Esse cadastro não existe", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message":"Cadastro deletado com sucesso!"})
}