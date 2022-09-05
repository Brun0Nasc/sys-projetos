package tasks

import (
	"github.com/Brun0Nasc/sys-projetos/config/database"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/tasks/model"
	"github.com/Brun0Nasc/sys-projetos/infra/tasks"
)

func NovaTask(req *modelApresentacao.ReqTask) (*modelApresentacao.ReqTask, error) {
	db := database.Conectar()
	defer db.Close()
	tasksRepo := tasks.NovoRepo(db)

	ver, err := tasksRepo.VerificaPessoa(req.PessoaID, req.ProjetoID);
	if err != nil {
		return nil, err
	}

	if !ver{
		return nil, nil
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