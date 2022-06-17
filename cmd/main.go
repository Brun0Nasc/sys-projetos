package main

import (
	"github.com/Brun0Nasc/sys-projetos/pkg/common/db"
	"github.com/Brun0Nasc/sys-projetos/pkg/equipes"
	"github.com/Brun0Nasc/sys-projetos/pkg/pessoas"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	pessoas.RegisterRoutes(r, h)
	equipes.RegisterRoutes(r, h)

	r.Run(port)
}