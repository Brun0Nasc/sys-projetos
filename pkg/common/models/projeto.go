package models

type Projeto struct {
	ID 			uint 	`gorm:"primary_key" json:"id"`
	Nome 		string 	`gorm:"type: varchar(30) not null" json:"nome"`
	EquipeID 	int 	`json:"equipeId"`
	Equipe 		Equipe 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"equipe"`
}