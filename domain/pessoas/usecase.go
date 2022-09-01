package pessoas

import (
	"github.com/Brun0Nasc/sys-projetos/config/database"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/pessoas/model"
	"github.com/Brun0Nasc/sys-projetos/infra/pessoas"
)

func NovaPessoa(req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error) {
	db := database.Conectar()
	defer db.Close()
	pessoasRepo := pessoas.NovoRepo(db)

	return pessoasRepo.NovaPessoa(req)
}