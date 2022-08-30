package equipes

import (
	"database/sql"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/equipes/model"
	"github.com/Brun0Nasc/sys-projetos/infra/equipes/postgres"
)

type respositorio struct {
	Data *postgres.DBEquipes
}

func novoRepo(novoDB *sql.DB) *respositorio {
	return &respositorio{
		Data: &postgres.DBEquipes{DB: novoDB},
	}
}

func (r *respositorio) NovaEquipe(req *modelApresentacao.ReqEquipes) {
	r.Data.NovaEquipe(&modelData.Equipes{Nome_Equipe: req.Nome_Equipe})
}
func (r *respositorio) ListarEquipes() {

}