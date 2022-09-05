package tasks

import modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/tasks/model"

type ITask interface {
	NovaTask(req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error)
	ListarTasks() ([]modelApresentacao.ReqTask, error)
	BuscarTask(id string) (*modelApresentacao.ReqTask, error)
	AtualizarTask(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error)
	AtualizarStatus(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error)
	DeletarTask(id string) error
	VerificaPessoa(idPessoa int, idProjeto int) (bool, error)
}
