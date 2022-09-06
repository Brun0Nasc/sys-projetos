package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/user/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/user/model"
)

type DBUser struct {
	DB *sql.DB
}

func (postgres *DBUser) NovoUser(req *modelData.User) (*modelApresentacao.ReqUser, error) {
	var user = &modelApresentacao.ReqUser{}

	sqlStatement := `INSERT INTO user
	(nome_user, email, senha) VALUES($1::VARCHAR(80), $2::VARCHAR(80), &3::VARCHAR(80))
	RETURNING *`

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_User, req.Email, req.Senha)

	if err := row.Scan(&user.ID_User, &user.Nome_User, &user.Email,
	&user.Senha, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	fmt.Println("Usuário registrado com sucesso!")
	return user, nil
}