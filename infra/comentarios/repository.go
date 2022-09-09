package comentarios

import (
	"database/sql"

	"github.com/Brun0Nasc/sys-projetos/infra/comentarios/postgres"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/comentarios/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/comentarios/model"
)

type repositorio struct {
	Data *postgres.DBComentarios
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBComentarios{DB: novoDB},
	}
}

func (r *repositorio) NovoComentario(req *modelApresentacao.ReqComentario) (*modelApresentacao.ReqComentario, error) {
	return r.Data.NovoComentario(&modelData.Comentario{
		Task_ID: req.Task_ID,
		Comentario: req.Comentario,
	})
}

func (r *repositorio) ListarComentarios(id string) ([]*modelApresentacao.ReqComentario, error){
	return r.Data.ListarComentarios(id)
}