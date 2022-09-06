package user

import modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/user/model"

type IUser interface {
	NovoUser(req *modelApresentacao.ReqUser) (*modelApresentacao.ReqUser, error)
}