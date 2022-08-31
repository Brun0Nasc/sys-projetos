package postgres

import (
	"database/sql"
	"fmt"
	"net/http"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/equipes/model"
	"github.com/gin-gonic/gin"
)

type DBEquipes struct {
	DB *sql.DB
}

func (postgres *DBEquipes) NovaEquipe(req *modelData.Equipe, c *gin.Context) {
	sqlStatement := `INSERT INTO equipes
	(nome_equipe)
	VALUES($1::TEXT)`
	_, err := postgres.DB.Exec(sqlStatement, req.Nome_Equipe)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	fmt.Println("deu tudo certo")
}

func (postgres *DBEquipes) ListarEquipes(c *gin.Context) []modelApresentacao.ReqEquipe {
	sqlStatement := `SELECT * FROM equipes ORDER BY id_equipe`
	var res = []modelApresentacao.ReqEquipe{}
	var equipe = modelApresentacao.ReqEquipe{}

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	for rows.Next() {
		if err := rows.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe); err != nil {
			if err == sql.ErrNoRows {
				return nil
			} else {
				c.AbortWithError(http.StatusNotFound, err)
				return nil
			}
		}
		res = append(res, equipe)
	}
	fmt.Println("Listagem deu certo!")
	return res
}
