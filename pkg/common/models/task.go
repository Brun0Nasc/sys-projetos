package models

type Task struct {
	ID         	uint `gorm:"primary_key" json:"id"`
	Descricao    string `gorm:"type: varchar(100) not null" json:"descricao"`
	PessoaID  	int
	Pessoa		Pessoa `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pessoa"`
	ProjetoID 	int 
	Projeto		Projeto `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"projeto"`
}