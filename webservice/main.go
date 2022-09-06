package main

import (
	"os"

	"github.com/Brun0Nasc/sys-projetos/webservice/equipes"
	"github.com/Brun0Nasc/sys-projetos/webservice/login"
	"github.com/Brun0Nasc/sys-projetos/webservice/pessoas"
	"github.com/Brun0Nasc/sys-projetos/webservice/projetos"
	"github.com/Brun0Nasc/sys-projetos/webservice/tasks"
	"github.com/Brun0Nasc/sys-projetos/webservice/user"
	"github.com/Brun0Nasc/sys-projetos/zpkg/common/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.Use(cors.Default())

	eq := r.Group("equipes", middlewares.Auth())
	pe := r.Group("pessoas", middlewares.Auth())
	pr := r.Group("projetos", middlewares.Auth())
	tk := r.Group("tasks", middlewares.Auth())
	us := r.Group("user", middlewares.Auth())
	lg := r.Group("login", middlewares.Auth())

	equipes.Router(eq)
	pessoas.Router(pe)
	projetos.Router(pr)
	tasks.Router(tk)
	user.Router(us)
	login.Router(lg)

	r.Run(":" + port)
}