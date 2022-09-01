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