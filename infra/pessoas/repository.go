package pessoas

import (
	"database/sql"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/pessoas/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/pessoas/model"
	"github.com/Brun0Nasc/sys-projetos/infra/pessoas/postgres"
)

type repositorio struct {
	Data *postgres.DBPessoas
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBPessoas{DB: novoDB},
	}
}

func (r *repositorio) NovaPessoa(req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error) {
	return r.Data.NovaPessoa(&modelData.Pessoa{
		Nome_Pessoa: req.Nome_Pessoa,
		Funcao_Pessoa: req.Funcao_Pessoa,
		EquipeID: req.EquipeID,
	})
}
func (r *repositorio) ListarPessoas() ([]modelApresentacao.ReqPessoa, error){
	return r.Data.ListarPessoas()
}
func (r *repositorio) BuscarPessoa(id string) (*modelApresentacao.ReqPessoa, error){
	return nil, nil
}
func (r *repositorio) DeletarPessoa(id string) error{
	return nil
}
func (r *repositorio) AtualizarPessoa(id string, req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error) {
	return nil, nil
}
func (r *repositorio) FavoritarPessoa(id string) error {
	return nil
}