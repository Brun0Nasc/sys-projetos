package tasks

import (
	"database/sql"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/tasks/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/tasks/model"
	"github.com/Brun0Nasc/sys-projetos/infra/tasks/postgres"
)

type repositorio struct {
	Data *postgres.DBTasks
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBTasks{DB: novoDB},
	}
}

func (r *repositorio) NovaTask(req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error) {
	return r.Data.NovaTask(&modelData.Task{
		Descricao_Task: req.Descricao_Task,
		Nivel:          req.Nivel,
		PessoaID:       req.PessoaID,
		ProjetoID:      req.ProjetoID,
	})
}

func (r *repositorio) ListarTasks() ([]modelApresentacao.ReqTask, error) {
	return r.Data.ListarTasks()
}

func (r *repositorio) BuscarTask(id string) (*modelApresentacao.ReqTask, error) {
	return r.Data.BuscarTask(id)
}

func (r *repositorio) AtualizarTask(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error) {
	return r.Data.AtualizarTask(id, &modelData.Task{
		Descricao_Task: req.Descricao_Task,
		Nivel:          req.Nivel,
		PessoaID:       req.PessoaID,
		ProjetoID:      req.ProjetoID,
	})
}

func (r *repositorio) AtualizarStatus(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error) {
	return r.Data.AtualizarStatus(id, &modelData.Task{Status: req.Status})
}

func (r *repositorio) DeletarTask(id string) error {
	return r.Data.DeletarTask(id)
}

func (r *repositorio) VerificaPessoa(idPessoa int, idProjeto int) (bool, error) {
	return r.Data.VerificaPessoa(idPessoa, idProjeto)
}

func (r *repositorio) TasksPessoa(id string) ([]modelApresentacao.ReqTask, error) {
	return r.Data.TasksPessoa(id)
}

func (r *repositorio) CheckStatus(id string) (*int, error) {
	return r.Data.CheckStatus(id)
}