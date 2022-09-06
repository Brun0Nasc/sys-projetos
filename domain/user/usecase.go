package user

import (
	"github.com/Brun0Nasc/sys-projetos/config/database"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/user/model"
	"github.com/Brun0Nasc/sys-projetos/infra/user"
	"github.com/Brun0Nasc/sys-projetos/config/services"
)

func NovoUser(req *modelApresentacao.ReqUser) (*modelApresentacao.ReqUser, error) {
	db := database.Conectar()
	defer db.Close()
	userRepo := user.NovoRepo(db)

	req.Senha = services.SHA256Encoder(req.Senha)

	return userRepo.NovoUser(req)
}