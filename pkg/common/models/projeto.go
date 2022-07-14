package models

type Projeto struct {
	ID_Projeto 		uint 	`gorm:"primary_key" json:"id_projeto"`
	Nome_Projeto 	string 	`json:"nome_projeto"`
	EquipeID 		int 	`json:"equipe_id"`
	Status			string	`json:"status"`
}