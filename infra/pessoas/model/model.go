package pessoas

import "time"

type Pessoa struct {
	ID_Pessoa       *uint
	Nome_Pessoa     *string
	Funcao_Pessoa   *string
	EquipeID        *uint
	Favoritar       *int
	DataContratacao *time.Time
	UpdatedAt       *time.Time
}
