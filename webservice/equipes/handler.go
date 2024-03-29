package equipes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/domain/equipes"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"

	"github.com/gin-gonic/gin"
)

func novaEquipe(c *gin.Context) {
	fmt.Println("Tentando adicionar nova equipe")
	req := modelApresentacao.ReqEquipe{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly ", "err":err.Error(),
		})
		return
	}

	res, err := equipes.NovaEquipe(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func listarEquipes(c *gin.Context) {
	fmt.Println("Tentando listar equipes") 
	if res, err := equipes.ListarEquipes(); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusOK, res)
		return
	}
}

func buscarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando encontrar equipe")
	if res, err := equipes.BuscarEquipe(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusOK, res)
		return
	}
}

func deletarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar equipe")

	err := equipes.DeletarEquipe(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"Cadastro deletado com sucesso!"})
}

func atualizarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando atualizar equipe")

	req := modelApresentacao.ReqEquipe{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not update. Parameters were not passed correctly.", "err":err.Error(),
		})
		return
	}

	res, err := equipes.AtualizarEquipe(id, &req)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func projetosEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando listar projetos de uma equipe")
	
	res, err := equipes.ProjetosEquipe(id)
	if err != nil {
		if err == sql.ErrNoRows{
			c.JSON(200, gin.H{"message":"Essa equipe ainda não possui projetos", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	}
	c.JSON(200, res)
}