package projetos

import (
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/projetos/model"
)

type IProjeto interface {
	NovoProjeto(req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjeto, error)
	ListarProjetos() ([]modelApresentacao.ReqProjeto, error)
	BuscarProjeto(id string) (*modelApresentacao.ReqProjeto, error)
	AtualizarProjeto(id string, req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjeto, error)
	//AtualizarStatus(id string, req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjeto, error)
	DeletarProjeto(id string) error
}