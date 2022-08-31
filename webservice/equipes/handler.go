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

	if equipe, err := equipe.NovaEquipe(&req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
		return
	} else{
		c.JSON(http.StatusCreated, equipe)
	}
}

func listarEquipes(c *gin.Context) {
	fmt.Println("Tentando listar equipes") 
	if equipes, err := equipe.ListarEquipes(); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusOK, equipes)
		return
	}
}

func buscarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando encontrar equipe")
	if equipe, err := equipe.BuscarEquipe(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusOK, equipe)
		return
	}
}

func deletarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando deletar equipe")
	if err := equipe.DeletarEquipe(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message":"Cadastro deletado com sucesso!"})
		return
	}
}