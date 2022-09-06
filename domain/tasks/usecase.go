package tasks

import (
	"fmt"

	"github.com/Brun0Nasc/sys-projetos/config/database"
	modelProjeto "github.com/Brun0Nasc/sys-projetos/domain/projetos/model"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/tasks/model"
	"github.com/Brun0Nasc/sys-projetos/infra/projetos"
	"github.com/Brun0Nasc/sys-projetos/infra/tasks"
)

func NovaTask(req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)
	projetosRepo := projetos.NovoRepo(db)

	ver, err := tasksRepo.VerificaPessoa(req.PessoaID, req.ProjetoID)
	if err != nil {
		return nil, err
	}

	if !ver {
		return nil, nil
	}

	proj, err := projetosRepo.BuscarProjeto(fmt.Sprint(req.ProjetoID))
	if err != nil {
		return nil, err
	}

	if proj.Status == "Concluído" {
		modelProj := modelProjeto.ReqProjeto{
			Status: "Em desenvolvimento",
		}
		projetosRepo.AtualizarStatus(fmt.Sprint(req.ProjetoID), &modelProj)
	}

	return tasksRepo.NovaTask(req)
}

func ListarTasks() ([]modelApresentacao.ReqTask, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.ListarTasks()
}

func BuscarTask(id string) (*modelApresentacao.ReqTask, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.BuscarTask(id)
}

func AtualizarTask(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.AtualizarTask(id, req)
}

func AtualizarStatus(id string, req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)
	projetosRepo := projetos.NovoRepo(db)

	task, err := BuscarTask(id)
	if err != nil {
		return nil, err
	}

	if task.Status == "Concluído" {
		pId := fmt.Sprint(task.ProjetoID)
		projeto, err := projetosRepo.BuscarProjeto(pId)
		if err != nil {
			return nil, err
		}
		if projeto.Status == "Concluído" {
			modelProj := modelProjeto.ReqProjeto{
				Status: "Em desenvolvimento",
			}
			projetosRepo.AtualizarStatus(pId, &modelProj)
		}
	}

	return tasksRepo.AtualizarStatus(id, req)
}

func DeletarTask(id string) error {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.DeletarTask(id)
}

func TasksPessoa(id string) ([]modelApresentacao.ReqTask, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.TasksPessoa(id)
}

func CheckStatus(id string) (*int, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	return tasksRepo.CheckStatus(id)
}
