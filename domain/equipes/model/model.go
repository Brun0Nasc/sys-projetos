package equipes

import (
	modelPessoa "github.com/Brun0Nasc/sys-projetos/domain/pessoas/model"
	"time"
)

type ReqEquipe struct {
	ID_Equipe   *uint                    `json:"id_equipe"`
	Nome_Equipe *string                  `json:"nome_equipe,omitempty"`
	Pessoas     []*modelPessoa.ReqPessoa `json:"pessoas,omitempty"`
	CreatedAt   *time.Time               `json:"created_at"`
	UpdatedAt   *time.Time               `json:"updated_at"`
}
