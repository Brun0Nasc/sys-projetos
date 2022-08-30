package equipe

import (
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/config/database"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	"github.com/Brun0Nasc/sys-projetos/infra/equipes"

	"github.com/gin-gonic/gin"
)

func NovaEquipe(c *gin.Context, req *modelApresentacao.ReqEquipes) {
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)

	str := *req.Nome_Equipe

	req.Nome_Equipe = &str

	equipesRepo.NovaEquipe(req)

	c.JSON(http.StatusCreated, req)
}

func ListarEquipes() {
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)
	equipesRepo.ListarEquipes()

	
}