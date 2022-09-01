package pessoas

import (
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