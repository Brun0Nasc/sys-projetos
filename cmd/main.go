package main

import (
	//"os"

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
	//Arquivo de configuração que vai adicionar as informações do Banco de Dados
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	//Definição da porta em que a API vai rodar (Nesse caso vai pegar automaticamente do Heroku)
	//Pegando link do Banco de Dados do arquivo .env
	//port := os.Getenv("PORT")
	//dbUrl := os.Getenv("DB_URL")
	dbUrl := viper.Get("DB_URL").(string)

	//Atribuição de rotas à variável 'r' e definição das configurações cors 
	r := gin.Default()
	r.Use(cors.Default())

	//Variável de inicialização do Banco de Dados
	h := db.Init(dbUrl)

	//Carregamento de rotas de cada elemento da API
	pessoas.RegisterRoutes(r, h)
	equipes.RegisterRoutes(r, h)
	projetos.RegisterRoutes(r, h)
	tasks.RegisterRoutes(r, h)

	//Comando para fazer o programa rodar
	//r.Run(":" + port)
	r.Run("localhost:3030")
}