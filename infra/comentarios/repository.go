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

func (r *repositorio) NovoComentario(id *uint, req *modelApresentacao.ReqComentario) (*modelApresentacao.ReqComentario, error) {
	return r.Data.NovoComentario(id, &modelData.Comentario{
		Comentario: req.Comentario,
	})
}

func (r *repositorio) ListarComentarios(id string) ([]modelApresentacao.ReqComentario, error){
	return r.Data.ListarComentarios(id)
}

func (r *repositorio) AtualizarComentario(id string, req *modelApresentacao.ReqComentario) (*modelApresentacao.ReqComentario, error) {
	return r.Data.AtualizarComentario(id, &modelData.Comentario{
		Comentario: req.Comentario,
	})
}

func (r *repositorio) DeletarComentario(id string) error {
	return r.Data.DeletarComentario(id)
}