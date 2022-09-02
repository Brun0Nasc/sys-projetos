package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/pessoas/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/pessoas/model"
)

type DBPessoas struct {
	DB *sql.DB
}

func (postgres *DBPessoas) NovaPessoa(req *modelData.Pessoa) (*modelApresentacao.ReqPessoa, error) {
	sqlStatement := `INSERT INTO pessoas
	(nome_pessoa, funcao_pessoa, equipe_id)
	VALUES($1::VARCHAR(80), $2::VARCHAR(80), $3::BIGINT)
	RETURNING *`

	var pessoa = &modelApresentacao.ReqPessoa{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Pessoa, req.Funcao_Pessoa, req.EquipeID)

	if err := row.Scan(&pessoa.ID_Pessoa, &pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa, &pessoa.EquipeID,
	&pessoa.Favoritar, &pessoa.DataContratacao, &pessoa.UpdatedAt); err != nil {
		return nil, err
	}

	return pessoa, nil
}

func (postgres *DBPessoas) ListarPessoas() ([]modelApresentacao.ReqPessoa, error) {
	sqlStatement := `SELECT * FROM pessoas ORDER BY id_pessoa;`

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	pessoa := modelApresentacao.ReqPessoa{}
	res := []modelApresentacao.ReqPessoa{}

	for rows.Next() {
		if err := rows.Scan(&pessoa.ID_Pessoa, &pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa,
		&pessoa.EquipeID, &pessoa.Favoritar, &pessoa.DataContratacao, &pessoa.UpdatedAt); err != nil {
			return nil, err
		}
		res = append(res, pessoa)
	}

	return res, nil
}

func (postgres *DBPessoas) BuscarPessoa(id string) (*modelApresentacao.ReqPessoa, error) {
	sqlStatement := `SELECT * FROM pessoas WHERE id_pessoa = $1`
	res := modelApresentacao.ReqPessoa{}

	row := postgres.DB.QueryRow(sqlStatement, id)

	if err := row.Scan(&res.ID_Pessoa, &res.Nome_Pessoa, &res.Funcao_Pessoa,
	&res.EquipeID, &res.Favoritar, &res.DataContratacao, &res.UpdatedAt); err != nil {
		return nil, err
	}

	return &res, nil
}

func (postgres *DBPessoas) AtualizarPessoa(id string, req *modelData.Pessoa) (*modelApresentacao.ReqPessoa, error) {
	sqlStatement := `UPDATE pessoas 
	SET nome_pessoa = $1::VARCHAR(80), funcao_pessoa = $2::VARCHAR(80), equipe_id = $3::BIGINT 
	WHERE id_pessoa = $4
	RETURNING *;`

	res := &modelApresentacao.ReqPessoa{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Pessoa, req.Funcao_Pessoa, req.EquipeID, id)

	if err := row.Scan(&res.ID_Pessoa, &res.Nome_Pessoa, &res.Funcao_Pessoa, &res.EquipeID,
	&res.Favoritar, &res.DataContratacao, &res.UpdatedAt); err != nil {
		return nil, err
	}

	return res, nil
}

func (postgres *DBPessoas) DeletarPessoa(id string) error {
	if _, err := postgres.BuscarPessoa(id); err != nil {
		return err
	} else {
		sqlStatement := `DELETE FROM pessoas WHERE id_pessoa = $1`

		_, err := postgres.DB.Exec(sqlStatement, id)
		if err != nil {
			return err
		}

		fmt.Println("Deletado com sucesso!")
		return nil
	}
}