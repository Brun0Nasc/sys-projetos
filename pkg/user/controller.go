package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	//Registro das rotas de equipes
	
	h := &handler{
		DB: db,
	}

	//Rotas Get, Post, Put e Delete, referentes a Equipes
	routes := r.Group("/user")
	routes.POST("/", h.AddUser)
}

