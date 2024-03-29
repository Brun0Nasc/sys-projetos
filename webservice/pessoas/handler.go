package pessoas

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/domain/pessoas"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/pessoas/model"

	"github.com/gin-gonic/gin"
)

func novaPessoa(c *gin.Context) {
	fmt.Println("Tentando adicionar nova pessoa")

	req := modelApresentacao.ReqPessoa{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
		})
		return
	}

	if res, err := pessoas.NovaPessoa(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func listarPessoas(c *gin.Context) {
	fmt.Println("Tentando listar pessoas")

	if res, err := pessoas.ListarPessoas(); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Não há dados cadastrados no banco.", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func buscarPessoa(c *gin.Context) {
	fmt.Println("Tentando buscar pessoa")

	id := c.Param("id")

	res, err := pessoas.BuscarPessoa(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"message":"Cadastro não encontrado.", "err":err.Error()})
		} else {
			c.JSON(404, gin.H{"err":err.Error()})
		}
		return
	}

	c.JSON(200, res)
}

func atualizarPessoa(c *gin.Context) {
	fmt.Println("Tentando atualizar pessoa")

	id := c.Param("id")
	req := modelApresentacao.ReqPessoa{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
		})
		return
	}

	res, err := pessoas.AtualizarPessoa(id, &req)
	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{"message":"Algo deu errado", "err":err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(200, res)
}

func deletarPessoa(c *gin.Context) {
	fmt.Println("Tentando deletar pessoa.")
	id := c.Param("id")

	err := pessoas.DeletarPessoa(id)
	if err != nil {
		c.JSON(404, gin.H{"err":err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"message":"Cadastro deletado com sucesso!"})
}

func favoritar(c *gin.Context) {
	fmt.Println("Tentando favoritar.")
	id := c.Param("id")

	err := pessoas.Favoritar(id)
	if err != nil {
		c.JSON(404, gin.H{"err":err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"message":"Pessoa favoritada!"})
}