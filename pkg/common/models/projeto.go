package models

type Projeto struct {
	ID uint `gorm:"primary_key" json:"id"`
	Nome string
}