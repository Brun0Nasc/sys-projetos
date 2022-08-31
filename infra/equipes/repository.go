package equipes

import (
	"database/sql"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/equipes/model"
	"github.com/Brun0Nasc/sys-projetos/infra/equipes/postgres"
	"github.com/gin-gonic/gin"
)

type repositorio struct {
	Data *postgres.DBEquipes
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBEquipes{DB: novoDB},
	}
}

func (r *repositorio) NovaEquipe(req *modelApresentacao.ReqEquipe, c *gin.Context) {
	r.Data.NovaEquipe(&modelData.Equipe{Nome_Equipe: req.Nome_Equipe}, c)
}
func (r *repositorio) ListarEquipes(c *gin.Context) []modelApresentacao.ReqEquipe {
	return r.Data.ListarEquipes(c)
}
func (r *repositorio) BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error) {
	return r.Data.BuscarEquipe(id)
}