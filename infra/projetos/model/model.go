package projetos

import "time"

type Projeto struct {
	ID_Projeto        uint
	Nome_Projeto      string
	Descricao_Projeto string
	EquipeID          int
	Status            string
	DataInicio        *time.Time
	UpdatedAt         *time.Time
	DataConclusao     *time.Time
}
