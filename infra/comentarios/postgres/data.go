package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/comentarios/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/comentarios/model"
)

type DBComentarios struct {
	DB *sql.DB
}

func (postgres *DBComentarios) NovoComentario(id *uint, req *modelData.Comentario) (*modelApresentacao.ReqComentario, error) {
	sqlStatement := `INSERT INTO comentarios (task_id, comentario)
	VALUES($1::BIGINT, $2::TEXT)
	RETURNING *`

	var res = &modelApresentacao.ReqComentario{}

	row := postgres.DB.QueryRow(sqlStatement, id, req.Comentario)
	if err := row.Scan(&res.ID_Comentario, &res.Task_ID, &res.Comentario,
		&res.CreatedAt, &res.UpdatedAt); err != nil {
		return res, err
	}

	fmt.Println("Comentário adicionado.")
	return res, nil
}

func (postgres *DBComentarios) ListarComentarios(id string) ([]modelApresentacao.ReqComentario, error) {
	sqlStatement := `SELECT * FROM comentarios WHERE task_id = $1 ORDER BY id_comentario;`
	var res = []modelApresentacao.ReqComentario{}
	var comentario = modelApresentacao.ReqComentario{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		if err := rows.Scan(&comentario.ID_Comentario, &comentario.Task_ID, &comentario.Comentario,
			&comentario.CreatedAt, &comentario.UpdatedAt); err != nil {
			return res, err
		}
		res = append(res, comentario)
	}

	fmt.Println("Listando comentários")
	return res, nil
}

func (postgres *DBComentarios) AtualizarComentario(id string, req *modelData.Comentario) (*modelApresentacao.ReqComentario, error) {
	sqlStatement := `UPDATE comentarios SET comentario=$1::TEXT WHERE id_comentario=$2 RETURNING *;`
	var res = modelApresentacao.ReqComentario{}

	row := postgres.DB.QueryRow(sqlStatement, req.Comentario, id)

	if err := row.Scan(&res.ID_Comentario, &res.Task_ID, &res.Comentario, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return &res, err
	}
	fmt.Println("Comentário atualizado")
	return &res, nil
}

func (postgres *DBComentarios) DeletarComentario(id string) error {
	comentario := modelApresentacao.ReqComentario{}
	row := postgres.DB.QueryRow(`SELECT * FROM comentarios WHERE id_comentario=$1`, id)
	err := row.Scan(&comentario.ID_Comentario, &comentario.Task_ID, &comentario.Comentario, &comentario.CreatedAt,
	&comentario.UpdatedAt)
	if err != nil {
		return err
	} else {
		sqlStatement := `DELETE FROM comentarios WHERE id_comentario=$1;`

		_, err := postgres.DB.Exec(sqlStatement, id)
		if err != nil {
			return err
		}
		fmt.Println("Comentário deletado")
		return nil
	}
}
