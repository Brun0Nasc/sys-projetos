package postgres

import (
	"database/sql"
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