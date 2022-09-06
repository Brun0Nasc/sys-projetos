package login

import (
	"database/sql"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/login/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/login/model"
	modelUser "github.com/Brun0Nasc/sys-projetos/domain/user/model"
	"github.com/Brun0Nasc/sys-projetos/infra/login/postgres"
)

type repositorio struct {
	Data *postgres.DBLogin
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBLogin{DB: novoDB},
	}
}

func (r *repositorio) Logar(req *modelApresentacao.ReqMakeLogin) (*modelUser.ReqUser, error) {
	return r.Data.Logar(&modelData.MakeLogin{
		Email: req.Email,
		Senha: req.Senha,
	})
}