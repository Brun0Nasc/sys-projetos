package main

import (
	"os"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/db"
	"github.com/Brun0Nasc/sys-projetos/pkg/equipes"
	"github.com/Brun0Nasc/sys-projetos/pkg/pessoas"
	"github.com/Brun0Nasc/sys-projetos/pkg/projetos"
	"github.com/Brun0Nasc/sys-projetos/pkg/tasks"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := os.Getenv("PORT")
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	r.Use(cors.Default())

	h := db.Init(dbUrl)

	pessoas.RegisterRoutes(r, h)
	equipes.RegisterRoutes(r, h)
	projetos.RegisterRoutes(r, h)
	tasks.RegisterRoutes(r, h)

	r.Run(":" + port)
}