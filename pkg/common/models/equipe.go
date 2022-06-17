package models

type Equipe struct {
    ID          uint   `gorm:"primary_key" json:"id"`
    Nome       string `json:"nome"`
}