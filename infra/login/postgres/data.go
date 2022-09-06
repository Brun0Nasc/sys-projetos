package postgres

import (
	"database/sql"

	modelData "github.com/Brun0Nasc/sys-projetos/infra/login/model"
	modelUser "github.com/Brun0Nasc/sys-projetos/domain/user/model"
)

type DBLogin struct {
	DB *sql.DB
}

func (postgres *DBLogin) Logar(req *modelData.MakeLogin) (*modelUser.ReqUser, error) {
	sqlStatement := `SELECT * FROM "user" WHERE email = $1::VARCHAR(80);`
	user := modelUser.ReqUser{}
	row := postgres.DB.QueryRow(sqlStatement, req.Email)

	if err := row.Scan(&user.ID_User, &user.Nome_User, &user.Email, &user.Senha,
		&user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	if user.Senha != req.Senha {
		return nil, nil
	}

	return &user, nil
}