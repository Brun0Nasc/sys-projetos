package comentarios

import "database/sql"

func NovoRepo(DB *sql.DB) IComentario {
	return novoRepo(DB)
}