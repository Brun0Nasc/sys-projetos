package equipes

import "time"

type Equipe struct {
	ID_Equipe   *uint
	Nome_Equipe *string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
