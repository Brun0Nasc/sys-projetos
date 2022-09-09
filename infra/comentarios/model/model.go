package comentarios

import "time"

type Comentario struct {
	ID_Comentario *uint
	Task_ID       *uint
	Comentario    *string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
