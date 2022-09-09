package comentarios

import "time"

type ReqComentario struct {
	ID_Comentario *uint      `json:"id_comentario"`
	Task_ID       *uint      `json:"task_id"`
	Comentario    *string    `json:"comentario"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
