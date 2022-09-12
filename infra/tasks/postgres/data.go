package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/tasks/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/tasks/model"
) 

type DBTasks struct {
	DB *sql.DB
}

func (postgres *DBTasks) NovaTask(req *modelData.Task) (*modelApresentacao.ReqTask, error) {
	var task = &modelApresentacao.ReqTask{}

	sqlStatement := `INSERT INTO tasks
	(descricao_task, nivel, pessoa_id, projeto_id)
	VALUES($1::TEXT, $2::VARCHAR(80), $3::BIGINT, $4::BIGINT)
	RETURNING *`

	row := postgres.DB.QueryRow(sqlStatement, req.Descricao_Task, req.Nivel, req.PessoaID, req.ProjetoID)

	if err := row.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, 
	&task.ProjetoID, &task.Status, &task.Nivel, &task.CreatedAt, &task.UpdatedAt); 
	err != nil {
		return nil, err
	}

	fmt.Println("Task adicionada")
	return task, nil
}

func (postgres *DBTasks) ListarTasks() ([]modelApresentacao.ReqTask, error) {
	sqlStatement := `SELECT tk.*, pe.nome_pessoa FROM tasks as tk
	join pessoas as pe on tk.pessoa_id = pe.id_pessoa
	 ORDER BY id_task;`

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	task := modelApresentacao.ReqTask{}
	res := []modelApresentacao.ReqTask{}

	for rows.Next() {
		if err := rows.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, 
		&task.ProjetoID, &task.Status, &task.Nivel, &task.CreatedAt, &task.UpdatedAt, &task.Nome_Pessoa); err != nil {
			return nil, err
		}
		res = append(res, task)
	}

	fmt.Println("Listando tasks")
	return res, nil
}

func (postgres *DBTasks) BuscarTask(id string) (*modelApresentacao.ReqTask, error) {
	sqlStatement := `SELECT tk.*, pe.nome_pessoa FROM tasks as tk
	join pessoas as pe on tk.pessoa_id = pe.id_pessoa
	WHERE id_task=$1;`
	res := modelApresentacao.ReqTask{}

	row := postgres.DB.QueryRow(sqlStatement, id)

	if err := row.Scan(&res.ID_Task, &res.Descricao_Task, &res.PessoaID, 
		&res.ProjetoID, &res.Status, &res.Nivel, &res.CreatedAt, &res.UpdatedAt, &res.Nome_Pessoa); err != nil {
		return nil, err
	}

	fmt.Println("Task encontrada")
	return &res, nil
}

func (postgres *DBTasks) AtualizarTask(id string, req *modelData.Task) (*modelApresentacao.ReqTask, error) {
	sqlStatement := `UPDATE tasks SET descricao_task = $1::TEXT,
	nivel = $2::VARCHAR(80), pessoa_id = $3::BIGINT, projeto_id = $4::BIGINT WHERE id_task = $5
	RETURNING *`

	res := &modelApresentacao.ReqTask{}

	row := postgres.DB.QueryRow(sqlStatement, req.Descricao_Task, req.Nivel, req.PessoaID,
	req.ProjetoID, id)

	if err := row.Scan(&res.ID_Task, &res.Descricao_Task, &res.PessoaID, 
		&res.ProjetoID, &res.Status, &res.Nivel, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return nil, err
	}

	fmt.Println("Task atualizada")
	return res, nil
}

func (postgres *DBTasks) AtualizarStatus(id string, req *modelData.Task) (*modelApresentacao.ReqTask, error) {
	sqlStatement := `UPDATE tasks
	SET status=$1::VARCHAR(80)
	WHERE id_task=$2
	RETURNING *;`

	res := modelApresentacao.ReqTask{}

	row := postgres.DB.QueryRow(sqlStatement, req.Status, id)

	if err := row.Scan(&res.ID_Task, &res.Descricao_Task, &res.PessoaID, 
		&res.ProjetoID, &res.Status, &res.Nivel, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return nil, err
	}

	fmt.Println("Task atualizada")
	return &res, nil
}

func (postgres *DBTasks) DeletarTask(id string) error {
	if _, err := postgres.BuscarTask(id); err != nil {
		return err
	} else {
		sqlStatement := `DELETE FROM tasks WHERE id_task = $1`

		_, err := postgres.DB.Exec(sqlStatement, id)
		if err != nil {
			return err
		}

		fmt.Println("Task deletada com sucesso!")
		return nil
	}
}

func (postgres *DBTasks) VerificaPessoa(idPessoa int, idProjeto int) (bool, error) {
	sqlStatement := `select count(pr.nome_projeto) from pessoas as pe
	inner join projetos as pr on pe.equipe_id = pr.equipe_id
	where pe.id_pessoa = $1 and pr.id_projeto = $2`

	var ver int

	row := postgres.DB.QueryRow(sqlStatement, idPessoa, idProjeto)

	if err := row.Scan(&ver); err != nil {
		return false, err
	}
	if ver == 0{
		return false, nil
	} else {
		return true, nil
	}
}

func (postgres *DBTasks) TasksPessoa(id string) ([]modelApresentacao.ReqTask, error) {
	sqlStatement := `SELECT * FROM tasks WHERE pessoa_id =$1 ORDER BY id_task;`

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}

	task := modelApresentacao.ReqTask{}
	res := []modelApresentacao.ReqTask{}

	for rows.Next() {
		if err := rows.Scan(&task.ID_Task, &task.Descricao_Task, &task.PessoaID, 
		&task.ProjetoID, &task.Status, &task.Nivel, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		res = append(res, task)
	}

	fmt.Println("Listando tasks de uma pessoa")
	return res, nil
}

func (postgres *DBTasks) CheckStatus(id string) (*int, error) {
	sqlStatement := `SELECT count(id_task) FROM tasks WHERE projeto_id = $1 AND status != 'Concluído'`

	n := 0

	row := postgres.DB.QueryRow(sqlStatement, id)

	if err := row.Scan(&n); err != nil {
		return nil, err
	}

	return &n, nil
}