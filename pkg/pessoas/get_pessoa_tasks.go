package pessoas

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Result struct {
	ID_Task		   uint 	`json:"id_task"`
	ID_Pessoa	   uint		`json:"id_pessoa"`
	Nome_Pessoa    string	`json:"nome_pessoa"`
	Funcao_Pessoa  string  	`json:"funcao_pessoa"`
	Nome_Equipe    string	`json:"equipe_id"`
	Nome_Projeto   string	`json:"nome_projeto"`
	Descricao_Task string	`json:"descricao_task"`
}

func (h handler) GetTaskPessoa(c *gin.Context) {
	id := c.Param("id")

	var result []Result

	if result := h.DB.Raw("SELECT tk.id_task, pe.id_pessoa, pe.nome_pessoa, pe.funcao_pessoa, eq.nome_equipe, pr.nome_projeto, tk.descricao_task FROM pessoas AS pe INNER JOIN equipes AS eq ON pe.equipe_id = eq.id_equipe INNER JOIN projetos AS pr ON pr.equipe_id = eq.id_equipe INNER JOIN tasks as tk ON tk.projeto_id = pr.id_projeto AND tk.pessoa_id = pe.id_pessoa WHERE pe.id_pessoa = ?", id).Scan(&result); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &result)
}

