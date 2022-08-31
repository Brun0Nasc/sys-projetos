package postgres

import (
	"database/sql"
	"fmt"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/equipes/model"
)

type DBEquipes struct {
	DB *sql.DB
}

func (postgres *DBEquipes) NovaEquipe(req *modelData.Equipes) *modelData.Equipes{
	sqlStatement := `INSERT INTO equipes
	(nome_equipe)
	VALUES($1::TEXT)`
	_, err := postgres.DB.Exec(sqlStatement, req.Nome_Equipe)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("deu tudo certo")

	return req
}

func (postgres *DBEquipes) ListarEquipes() *modelData.Equipes{
	sqlStatement := `SELECT * FROM equipes ORDER BY id_equipe`
	rows, err := postgres.DB.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("deu tudo certo")

	return rows
}