package projetos

import "time"

type ReqProjeto struct {
	ID_Projeto        uint       `json:"id_projeto"`
	Nome_Projeto      string     `json:"nome_projeto"`
	Descricao_Projeto string     `json:"descricao_projeto"`
	EquipeID          int        `json:"equipe_id"`
	Status            string     `json:"status"`
	DataInicio        *time.Time `json:"data_inicio"`
	UpdatedAt         *time.Time `json:"updated_at"`
	DataConclusao     *time.Time `json:"data_conclusao"`
}