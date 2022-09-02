package projetos

import (
	"fmt"

	"github.com/Brun0Nasc/sys-projetos/config/database"
	"github.com/Brun0Nasc/sys-projetos/domain/equipes"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/projetos/model"
	"github.com/Brun0Nasc/sys-projetos/infra/projetos"
)

func NovoProjeto(req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjeto, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	if req.EquipeID > 0 {
		membros, err := equipes.ListarMembros(fmt.Sprint(req.EquipeID))
		if err != nil {
			return nil, err
		}
		if len(membros) == 0 {
			fmt.Println("Não dá pra cadastrar projeto em equipe sem membros")
			return nil, nil
		}
	}

	return projetosRepo.NovoProjeto(req)
}

func ListarProjetos() ([]modelApresentacao.ReqProjeto, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	return projetosRepo.ListarProjetos()
}

func BuscarProjeto(id string) (*modelApresentacao.ReqProjeto, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	return projetosRepo.BuscarProjeto(id)
}

func AtualizarProjeto(id string, req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjeto, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	if req.EquipeID > 0 {
		membros, err := equipes.ListarMembros(fmt.Sprint(req.EquipeID))
		if err != nil {
			return nil, err
		}
		if len(membros) == 0 {
			fmt.Println("Não dá pra cadastrar projeto em equipe sem membros")
			return nil, nil
		}
	}

	return projetosRepo.AtualizarProjeto(id, req)
}

func DeletarProjeto(id string) error {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	return projetosRepo.DeletarProjeto(id)
}