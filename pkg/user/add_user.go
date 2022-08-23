package user

import (
	"net/http"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/services"
	"github.com/gin-gonic/gin"
)

type AddUserRequestBody struct {
    Nome_User	string		`json:"nome_user"`
	Email		string		`json:"email"`
	Senha		string		`json:"senha"`
}

func (h handler) AddUser(c *gin.Context) {
	//Variável do tipo AddEquipeRequestBody que vai receber as informações passadas através do JSON
	body := AddUserRequestBody{}

	//Verificando o corpo de requisição
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//Variável do tipo equipe, que foi definido no pacote models, que vai criar o registro no banco de dados
	var user models.User

	//Transferência de informação que o body recebeu para o elemento equipe
	user.Nome_User = body.Nome_User
	user.Email = body.Email
	user.Senha = services.SHA256Encoder(body.Senha)

	//Verificação de erros e criação do registro no banco de dados
	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	//Se retornar o status de criação correto, exibe o registro gravado no banco de dados
	c.JSON(http.StatusCreated, &user)
}