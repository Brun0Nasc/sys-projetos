package login

type ReqMakeLogin struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
