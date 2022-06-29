package models

type Task struct {
	ID_Task         uint 	`json:"id_task"`
	Descricao_Task  string  `json:"descricao_task"`
	PessoaID  		int		`json:"pessoa_id"`
	ProjetoID 		int 	`json:"projeto_id"`
}