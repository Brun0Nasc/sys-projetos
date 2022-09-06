package equipes

import (
	"github.com/Brun0Nasc/sys-projetos/config/database"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	modelPessoa "github.com/Brun0Nasc/sys-projetos/domain/pessoas/model"
	"github.com/Brun0Nasc/sys-projetos/infra/equipes"
)

func NovaEquipe(req *modelApresentacao.ReqEquipe)(*modelApresentacao.ReqEquipe, error){
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)

	str := *req.Nome_Equipe

	req.Nome_Equipe = &str

	return equipesRepo.NovaEquipe(req)
}

func ListarEquipes() ([]modelApresentacao.ReqEquipe, error){
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)
	return equipesRepo.ListarEquipes()
}

func ListarMembros(id string) ([]modelPessoa.ReqPessoa, error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)
	return equipesRepo.ListarMembros(id)
}

func BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)
	return equipesRepo.BuscarEquipe(id)
}

func DeletarEquipe(id string) error {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)
	return equipesRepo.DeletarEquipe(id)
}

func AtualizarEquipe(id string, req *modelApresentacao.ReqEquipe) (*modelApresentacao.ReqEquipe, error){
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)

	str := *req.Nome_Equipe

	req.Nome_Equipe = &str

	return equipesRepo.AtualizarEquipe(id, req)
}