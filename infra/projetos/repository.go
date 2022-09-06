package projetos

import (
	"database/sql"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/projetos/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/projetos/model"
	"github.com/Brun0Nasc/sys-projetos/infra/projetos/postgres"
)

type repositorio struct {
	Data *postgres.DBProjetos
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBProjetos{DB: novoDB},
	}
}

func (r *repositorio) NovoProjeto(req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjeto, error) {
	return r.Data.NovoProjeto(&modelData.Projeto{
		Nome_Projeto:      req.Nome_Projeto,
		Descricao_Projeto: req.Descricao_Projeto,
		EquipeID:          req.EquipeID,
	})
}
func (r *repositorio) ListarProjetos() ([]modelApresentacao.ReqProjeto, error) {
	return r.Data.ListarProjetos()
}
func (r *repositorio) BuscarProjeto(id string) (*modelApresentacao.ReqProjeto, error) {
	return r.Data.BuscarProjeto(id)
}
func (r *repositorio) AtualizarProjeto(id string, req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjeto, error) {
	return r.Data.AtualizarProjeto(id, &modelData.Projeto{
		Nome_Projeto:      req.Nome_Projeto,
		Descricao_Projeto: req.Descricao_Projeto,
		EquipeID:          req.EquipeID,
	})
}
func (r *repositorio) AtualizarStatus(id string, req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjeto, error) {
	return r.Data.AtualizarStatus(id, &modelData.Projeto{Status: req.Status})
}
func (r *repositorio) DeletarProjeto(id string) error {
	return r.Data.DeletarProjeto(id)
}
