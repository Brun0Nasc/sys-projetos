package user

import "time"

type User struct {
	ID_User   uint     
	Nome_User string   
	Email     string   
	Senha     string   
	CreatedAt time.Time
	UpdatedAt time.Time
}
