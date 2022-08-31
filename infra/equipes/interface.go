package equipes

import (
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	"github.com/gin-gonic/gin"
)

type IEquipe interface {
	NovaEquipe(req *modelApresentacao.ReqEquipe, c *gin.Context)
	ListarEquipes(c *gin.Context) []modelApresentacao.ReqEquipe
}