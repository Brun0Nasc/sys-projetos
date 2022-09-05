package tasks

import "time"

type Task struct {
	ID_Task        uint
	Descricao_Task string
	Comentario     string
	Nivel          string
	PessoaID       int
	ProjetoID      int
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
