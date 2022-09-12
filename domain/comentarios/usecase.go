package comentarios

import(
	"github.com/Brun0Nasc/sys-projetos/config/database"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/comentarios/model"
	"github.com/Brun0Nasc/sys-projetos/infra/comentarios"
)

func NovoComentario(id *uint, req *modelApresentacao.ReqComentario)(*modelApresentacao.ReqComentario, error){
	db := database.Conectar()
	defer db.Close()
	comentariosRepo := comentarios.NovoRepo(db)
	return comentariosRepo.NovoComentario(id, req)
}

func ListarComentarios(id string) ([]modelApresentacao.ReqComentario, error) {
	db := database.Conectar()
	defer db.Close()
	comentariosRepo := comentarios.NovoRepo(db)
	return comentariosRepo.ListarComentarios(id)
}

func AtualizarComentario(id string, req *modelApresentacao.ReqComentario)(*modelApresentacao.ReqComentario, error){
	db := database.Conectar()
	defer db.Close()
	comentariosRepo := comentarios.NovoRepo(db)
	return comentariosRepo.AtualizarComentario(id, req)
}

func DeletarComentario(id string) error {
	db := database.Conectar()
	defer db.Close()
	comentariosRepo := comentarios.NovoRepo(db)
	return comentariosRepo.DeletarComentario(id)
}