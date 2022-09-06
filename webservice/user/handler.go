package user

import (
	"fmt"

	"github.com/Brun0Nasc/sys-projetos/domain/user"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/user/model"

	"github.com/gin-gonic/gin"
)

func pegaJSON(c *gin.Context) *modelApresentacao.ReqUser {
	req := modelApresentacao.ReqUser{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
		})
		return nil
	}
	return &req
}

func novoUser(c *gin.Context) {
	fmt.Println("Tentando adicionar um novo usuário.")
	req := pegaJSON(c)
	res, err := user.NovoUser(req)

	if err != nil {
		c.JSON(400, gin.H{
			"err":err.Error(),
		})
		return
	}

	c.JSON(201, res)
}