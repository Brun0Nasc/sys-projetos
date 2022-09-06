package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/projetos/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/projetos/model"
)

type DBProjetos struct {
	DB *sql.DB
}

func (postgres *DBProjetos) NovoProjeto(req *modelData.Projeto) (*modelApresentacao.ReqProjeto, error) {
	var projeto = &modelApresentacao.ReqProjeto{}

	sqlStatement := `INSERT INTO projetos
	(nome_projeto, descricao_projeto, equipe_id)
	VALUES($1::VARCHAR(80), $2::TEXT, $3::BIGINT)
	RETURNING *`

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Projeto, req.Descricao_Projeto, req.EquipeID)

	if err := row.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Descricao_Projeto,
		&projeto.EquipeID, &projeto.Status, &projeto.DataInicio, &projeto.UpdatedAt, &projeto.DataConclusao); err != nil {
		return nil, err
	}

	fmt.Println("Projeto adicionado")
	return projeto, nil
}

func (postgres *DBProjetos) ListarProjetos() ([]modelApresentacao.ReqProjeto, error) {
	sqlStatement := `SELECT * FROM projetos ORDER BY id_projeto;`

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	projeto := modelApresentacao.ReqProjeto{}
	res := []modelApresentacao.ReqProjeto{}

	for rows.Next() {
		if err := rows.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Descricao_Projeto,
			&projeto.EquipeID, &projeto.Status, &projeto.DataInicio, &projeto.UpdatedAt, &projeto.DataConclusao); err != nil {
			return nil, err
		}
		res = append(res, projeto)
	}

	return res, nil
}

func (postgres *DBProjetos) BuscarProjeto(id string) (*modelApresentacao.ReqProjeto, error) {
	sqlStatement := `SELECT * FROM projetos WHERE id_projeto = $1`
	res := modelApresentacao.ReqProjeto{}

	row := postgres.DB.QueryRow(sqlStatement, id)

	if err := row.Scan(&res.ID_Projeto, &res.Nome_Projeto, &res.Descricao_Projeto,
		&res.EquipeID, &res.Status, &res.DataInicio, &res.UpdatedAt, &res.DataConclusao); err != nil {
		return nil, err
	}

	return &res, nil
}

func (postgres *DBProjetos) AtualizarProjeto(id string, req *modelData.Projeto) (*modelApresentacao.ReqProjeto, error) {
	sqlStatement := `UPDATE projetos 
	SET nome_projeto = $1::VARCHAR(80), descricao_projeto = $2::TEXT, 
	equipe_id = $3::BIGINT 
	WHERE id_projeto = $4
	RETURNING *;`

	res := &modelApresentacao.ReqProjeto{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Projeto, req.Descricao_Projeto, req.EquipeID, id)

	if err := row.Scan(&res.ID_Projeto, &res.Nome_Projeto, &res.Descricao_Projeto,
		&res.EquipeID, &res.Status, &res.DataInicio, &res.UpdatedAt, &res.DataConclusao); err != nil {
		return nil, err
	}

	return res, nil
}

func (postgres *DBProjetos) DeletarProjeto(id string) error {
	if _, err := postgres.BuscarProjeto(id); err != nil {
		return err
	} else {
		sqlStatement := `DELETE FROM projetos WHERE id_projeto = $1`

		_, err := postgres.DB.Exec(sqlStatement, id)
		if err != nil {
			return err
		}

		fmt.Println("Deletado com sucesso!")
		return nil
	}
}

func (postgres *DBProjetos) AtualizarStatus(id string, req *modelData.Projeto) (*modelApresentacao.ReqProjeto, error) {
	sqlStatement := `UPDATE projetos
	SET status=$1::VARCHAR(80)
	WHERE id_projeto=$2
	RETURNING *;`

	if req.Status == "Concluído" {
		sqlStatement = `UPDATE projetos
		SET status=$1::VARCHAR(80), data_conclusao = now()
		WHERE id_projeto=$2
		RETURNING *;`
	}

	res := modelApresentacao.ReqProjeto{}

	row := postgres.DB.QueryRow(sqlStatement, req.Status, id)

	if err := row.Scan(&res.ID_Projeto, &res.Nome_Projeto, &res.Descricao_Projeto,
		&res.EquipeID, &res.Status, &res.DataInicio, &res.UpdatedAt, &res.DataConclusao); err != nil {
		return nil, err
	}

	return &res, nil
}
