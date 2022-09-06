package user

import "time"

type ReqUser struct {
	ID_User   uint      `json:"id_user"`
	Nome_User string    `json:"nome_user"`
	Email     string    `json:"email"`
	Senha     string    `json:"senha"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
