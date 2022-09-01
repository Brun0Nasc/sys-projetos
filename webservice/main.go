package main

import (
	"github.com/Brun0Nasc/sys-projetos/webservice/equipes"
	"github.com/Brun0Nasc/sys-projetos/webservice/pessoas"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eq := r.Group("equipes")
	pe := r.Group("pessoas")

	equipes.Router(eq)
	pessoas.Router(pe)

	r.Run("localhost:3030")
}