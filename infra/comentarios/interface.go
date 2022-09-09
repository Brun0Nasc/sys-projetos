package comentarios

import (
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/comentarios/model"
)

type IComentario interface {
	NovoComentario(req *modelApresentacao.ReqComentario) (*modelApresentacao.ReqComentario, error)
	ListarComentarios(id string) ([]*modelApresentacao.ReqComentario, error)
	//AtualizarComentario(id string, req *modelApresentacao.ReqComentario) (*modelApresentacao.ReqComentario, error)
	//DeletarComentario(id string) error
}