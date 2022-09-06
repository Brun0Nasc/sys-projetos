package login

import (
	"github.com/Brun0Nasc/sys-projetos/config/database"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/login/model"
	"github.com/Brun0Nasc/sys-projetos/infra/login"
	"github.com/Brun0Nasc/sys-projetos/config/services"
	modelUser "github.com/Brun0Nasc/sys-projetos/domain/user/model"
)

func Logar(req *modelApresentacao.ReqMakeLogin) (*modelUser.ReqUser, error) {
	db := database.Conectar()
	defer db.Close()
	loginRepo := login.NovoRepo(db)

	str := services.SHA256Encoder(req.Senha)

	req.Senha = str

	return loginRepo.Logar(req)
}