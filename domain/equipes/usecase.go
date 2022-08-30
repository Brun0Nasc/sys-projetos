package equipe

import (
	"github.com/Brun0Nasc/sys-projetos/config/database"
	"github.com/gin-gonic/gin"
)

func NovaEquipe(c *gin.Context) {
	db := database.Conectar()
	equipeRepo := equipe.NovoRepo(db)
	
	str := *req.Nome

	str = "Equipe: " + str

	req.Nome = &str

	equipeRepo.NovaEquipe()
}