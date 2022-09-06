package tasks

import "time"

type ReqTask struct {
	ID_Task        uint      `json:"id_task"`
	Descricao_Task string    `json:"descricao_task"`
	Comentario     string    `json:"comentario"`
	Nivel          string    `json:"nivel"`
	PessoaID       int       `json:"pessoa_id"`
	ProjetoID      int       `json:"projeto_id"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
