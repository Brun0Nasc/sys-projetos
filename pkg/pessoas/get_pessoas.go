package pessoas

import (
	"net/http"
	"time"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type GetPessoasBody struct{
	ID_Pessoa 		uint 			`json:"id_pessoa"`
	Nome_Pessoa 	string 			`json:"nome_pessoa"`
	Funcao_Pessoa 	string 			`json:"funcao_pessoa"`
	DataContratacao time.Time		`json:"data_contratacao"`
	EquipeID 		*int 			`json:"equipe_id"`
	Equipe 			*models.Equipe 	`json:"equipe"`
	Tasks 			*[]models.Task	`json:"tasks"`
	Favoritar		string			`json:"favoritar"`
}

func (h handler) GetPessoas(c *gin.Context) {
	var pessoas []models.Pessoa //Lista que vai ser usada para consultar as pessoas com os campos definidos no struct de Pessoa
	var pe []GetPessoasBody //Lista com os campos adicionais que vai ser retornada nas requisições

	// Select da lista geral de pessoas
	if result := h.DB.Find(&pessoas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// Laço para percorrer a lista criada na consulta anterior
	for i := 0; i < len(pessoas); i++{
		var tasks []models.Task // Lista do tipo Task que vai guardar as tasks de uma pessoa
		// Consulta para armazenar as tasks da pessoa referente ao iterável atual
		if result := h.DB.Raw("select * from tasks where pessoa_id = ?", pessoas[i].ID_Pessoa).Scan(&tasks); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		
		var data time.Time // Variável do tipo Time para armazenar a data de contratação das pessoas
		// Consulta para preencher a variável data
		if result := h.DB.Raw("select data_contratacao from pessoas where id_pessoa = ?", pessoas[i].ID_Pessoa).Scan(&data); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		var eq *int // Ponteiro do tipo int que vai ser usado para verificar se a pessoa está em alguma equipe
		var equipe *models.Equipe = nil // Ponteiro do tipo equipe que tem valor inicial nulo que vai armazenar a equipe da pessoa

		// Verificação para saber se a pessoa tem um ID de Equipe associado a ela
		if pessoas[i].EquipeID != 0{
			// Caso ela esteja associada a uma equipe, a variável eq vai receber o ID de equipe em pessoa
			eq = &pessoas[i].EquipeID
			// Em seguida a equipe relacionada será buscada no banco e armazenada na variável equipe
			if result := h.DB.First(&equipe, pessoas[i].EquipeID); result.Error != nil {
				c.AbortWithError(http.StatusNotFound, result.Error)
				return
			}
		} else {
			// Caso a pessoa não tenha equipe, eq será nulo
			eq = nil
		}
		
		// Declaração de uma variável do tipo GetPessoaBody que é preenchida com as informações coletadas nas consultas
		p := &GetPessoasBody{
			ID_Pessoa: pessoas[i].ID_Pessoa,
			Nome_Pessoa: pessoas[i].Nome_Pessoa,
			Funcao_Pessoa: pessoas[i].Funcao_Pessoa,
			DataContratacao: data,
			EquipeID: eq,
			Equipe: equipe,
			Tasks: &tasks,
			Favoritar: pessoas[i].Favoritar,
		}

		// Preenchendo a lista de pessoas com todas as informações desejáveis para requisições
		pe = append(pe, *p)
	}

	// Mostrando lista preenchida
	c.JSON(http.StatusOK, &pe)
}
