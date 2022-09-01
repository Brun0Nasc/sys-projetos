package pessoas

import "time"

type ReqPessoa struct {
	ID_Pessoa       *uint      `json:"id_pessoa"`
	Nome_Pessoa     *string    `json:"nome_pessoa"`
	Funcao_Pessoa   *string    `json:"funcao_pessoa"`
	EquipeID        *uint      `json:"equipe_id"`
	Favoritar       *int       `json:"favoritar"`
	DataContratacao *time.Time `json:"data_contratacao"`
	UpdatedAt       *time.Time `json:"updated_at"`
}