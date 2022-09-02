package pessoas

import (
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/pessoas/model"
)

type IPessoa interface {
	NovaPessoa(req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error)
	ListarPessoas() ([]modelApresentacao.ReqPessoa, error)
	BuscarPessoa(id string) (*modelApresentacao.ReqPessoa, error)
	DeletarPessoa(id string) error
	AtualizarPessoa(id string, req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error)
	Favoritar(id string) error
}