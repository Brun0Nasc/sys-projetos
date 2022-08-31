package equipes

import (
	"database/sql"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/equipes/model"
	"github.com/Brun0Nasc/sys-projetos/infra/equipes/postgres"
)

type repositorio struct {
	Data *postgres.DBEquipes
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBEquipes{DB: novoDB},
	}
}

func (r *repositorio) NovaEquipe(req *modelApresentacao.ReqEquipe) (*modelApresentacao.ReqEquipe, error){
	return r.Data.NovaEquipe(&modelData.Equipe{Nome_Equipe: req.Nome_Equipe})
}
func (r *repositorio) ListarEquipes() ([]modelApresentacao.ReqEquipe, error) {
	return r.Data.ListarEquipes()
}
func (r *repositorio) BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error) {
	return r.Data.BuscarEquipe(id)
}
func (r *repositorio) DeletarEquipe(id string) error {
	return r.Data.DeletarEquipe(id)
}