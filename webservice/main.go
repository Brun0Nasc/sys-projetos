package main

import (
	"github.com/Brun0Nasc/sys-projetos/webservice/equipes"
	"github.com/Brun0Nasc/sys-projetos/webservice/pessoas"
	"github.com/Brun0Nasc/sys-projetos/webservice/projetos"
	"github.com/Brun0Nasc/sys-projetos/webservice/tasks"
	"github.com/Brun0Nasc/sys-projetos/webservice/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eq := r.Group("equipes")
	pe := r.Group("pessoas")
	pr := r.Group("projetos")
	tk := r.Group("tasks")
	us := r.Group("user")

	equipes.Router(eq)
	pessoas.Router(pe)
	projetos.Router(pr)
	tasks.Router(tk)
	user.Router(us)

	r.Run("localhost:3030")
}