package user

import (
	"database/sql"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/user/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/user/model"
	"github.com/Brun0Nasc/sys-projetos/infra/user/postgres"
)

type repositorio struct {
	Data *postgres.DBUser
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBUser{DB: novoDB},
	}
}

func (r *repositorio) NovoUser(req *modelApresentacao.ReqUser) (*modelApresentacao.ReqUser, error) {
	return r.Data.NovoUser(&modelData.User{
		Nome_User: req.Nome_User,
		Email:     req.Email,
		Senha:     req.Senha,
	})
}
