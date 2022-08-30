package equipes

import modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"

type IEquipe interface {
	NovaEquipe(req *modelApresentacao.ReqEquipes)
	ListarEquipes()
}