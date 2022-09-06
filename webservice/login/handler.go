package login

import (
	"database/sql"

	"github.com/Brun0Nasc/sys-projetos/config/services"
	"github.com/Brun0Nasc/sys-projetos/domain/login"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/login/model"
	"github.com/gin-gonic/gin"
)

func logar(c *gin.Context) {
	req := modelApresentacao.ReqMakeLogin{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
		})
		return
	}

	user, err := login.Logar(&req)

	if err != nil {
		if err == sql.ErrNoRows{
			c.JSON(400, gin.H{"message":"Cadastro inexistente"})
		} else {
			c.JSON(400, gin.H{"err":err.Error()})
		}
		return
	}

	if user == nil && err == nil{
		c.JSON(400, gin.H{"err":"Senha incorreta!"})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID_User)

	if err != nil {
		c.JSON(500, gin.H{"err":err.Error()})
		return
	}

	c.JSON(200, gin.H{"token":token})
}