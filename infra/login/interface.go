package login

import (
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/login/model"
	modelUser "github.com/Brun0Nasc/sys-projetos/domain/user/model"
)

type ILogin interface {
	Logar(req *modelApresentacao.ReqMakeLogin) (*modelUser.ReqUser, error)
}