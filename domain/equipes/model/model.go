package equipes

import (
	"time"
	modelPessoa "github.com/Brun0Nasc/sys-projetos/infra/pessoas/model"
)

type ReqEquipe struct {
	ID_Equipe   *uint      `json:"id_equipe"`
	Nome_Equipe *string    `json:"nome_equipe,omitempty"`
	Pessoas 	[]modelPessoa.
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}