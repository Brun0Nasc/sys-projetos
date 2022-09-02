package equipes

import (
	modelPessoa "github.com/Brun0Nasc/sys-projetos/domain/pessoas/model"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
)

type IEquipe interface {
	NovaEquipe(req *modelApresentacao.ReqEquipe) (*modelApresentacao.ReqEquipe, error)
	ListarEquipes() ([]modelApresentacao.ReqEquipe, error)
	ListarMembros(id string) ([]modelPessoa.ReqPessoa, error)
	BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error)
	DeletarEquipe(id string) error
	AtualizarEquipe(id string, req *modelApresentacao.ReqEquipe) (*modelApresentacao.ReqEquipe, error)
}