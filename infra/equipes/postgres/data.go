package postgres

import (
	"database/sql"
	"fmt"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/equipes/model"
)

type DBEquipes struct {
	DB *sql.DB
}

func (postgres *DBEquipes) NovaEquipe(req *modelData.Equipes) {
	sqlStatement := `INSERT INTO equipes
	(nome_equipe)
	VALUES($1::TEXT)`
	_, err := postgres.DB.Exec(sqlStatement, req.Nome_Equipe)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("deu tudo certo")
}