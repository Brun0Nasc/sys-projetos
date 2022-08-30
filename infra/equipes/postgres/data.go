package postgres

import (
	"database/sql"
	"fmt"
	"net/http"

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

func (postgres *DBEquipes) ListarEquipes(w http.ResponseWriter) error{
	sqlStatement := `SELECT * FROM equipes ORDER BY id_equipe`
	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return err
	}
	
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)

	for i := range values {
		scanArgs[i] = &values[i]
	}

	masterData := make(map[string][]interface{})

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return err
		}
	}
}

// https://stackoverflow.com/questions/42774467/how-to-convert-sql-rows-to-typed-json-in-golang