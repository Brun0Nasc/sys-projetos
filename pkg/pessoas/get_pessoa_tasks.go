package pessoas

import (
	"net/http"

	//"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	//"github.com/Brun0Nasc/sys-projetos/pkg/common/db"
	"github.com/gin-gonic/gin"
)

type result struct {
	ID		 	uint	`json:"id"`
	Nome 	 	string	`json:"nome"`
	Funcao	 	string `json:"funcao"`
	EquipeID 	int	`json:"id_equipe"`
	Projeto	 	string	`json:"projeto.nome"`
	Descricao	string	`json:"descricao"`
}

func (h handler) GetTaskPessoa(c *gin.Context) {
	id := c.Param("id")

	var result []result

	if result := h.DB.Raw("SELECT pe.id, pe.nome, pe.funcao, pe.equipe_id, pr.nome, tk.descricao FROM pessoas AS pe INNER JOIN equipes AS eq ON pe.equipe_id = eq.id INNER JOIN projetos AS pr ON pr.equipe_id = eq.id INNER JOIN tasks as tk ON tk.projeto_id = pr.id AND tk.pessoa_id = pe.id WHERE pe.id = ?", id).Scan(&result); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &result)
}

