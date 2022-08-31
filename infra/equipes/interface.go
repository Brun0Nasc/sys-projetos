package equipes

import (
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
)

type IEquipe interface {
	NovaEquipe(req *modelApresentacao.ReqEquipe) error
	ListarEquipes() ([]modelApresentacao.ReqEquipe, error)
	BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error)
}