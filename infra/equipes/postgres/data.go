package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	modelPessoa "github.com/Brun0Nasc/sys-projetos/domain/pessoas/model"
	modelProjetos "github.com/Brun0Nasc/sys-projetos/domain/projetos/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/equipes/model"
)

type DBEquipes struct {
	DB *sql.DB
}

// A função adiciona uma nova equipe recebendo um corpo de requisição e retornando um corpo de resposta ou um erro
func (postgres *DBEquipes) NovaEquipe(req *modelData.Equipe) (*modelApresentacao.ReqEquipe, error) { 
	sqlStatement := `INSERT INTO equipes
	(nome_equipe)
	VALUES($1::VARCHAR(80))
	RETURNING *` // Comando sql que adiciona registro no banco e retorna registro adicionado

	var equipe = &modelApresentacao.ReqEquipe{} // Modelo que será usado para preencher e adicionar registro no banco

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Equipe) // Executando comando sql e recebendo a linha de resposta sql

	// Escaneando e testando informações recebidas após a execução do sql
	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.CreatedAt, &equipe.UpdatedAt); err != nil {
		return nil, err
	}

	// Retornando equipe adicionada e erro nulo
	return equipe, nil
}

func (postgres *DBEquipes) ListarEquipes() ([]modelApresentacao.ReqEquipe, error) {
	sqlStatement := `SELECT * FROM equipes ORDER BY id_equipe` // comando sql
	var res = []modelApresentacao.ReqEquipe{}                  // lista que vai receber resultados da consulta
	var equipe = modelApresentacao.ReqEquipe{}                 // estrutura individual que vai ser usada para preencher a lista

	rows, err := postgres.DB.Query(sqlStatement) // executando query e retornando as linhas e possíveis erros
	if err != nil {
		return nil, err // em caso de erro na consulta, a requisição será avortada e retornar um status 404 e o erro
	}
	for rows.Next() { // percorrendo linhas retornadas no sql
		if err := rows.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.CreatedAt, &equipe.UpdatedAt); err != nil { // escaneando linha por linha e gravando na estrutura de equipe
			return nil, err // se houver algum erro, a função retorna ele
		}
		res = append(res, equipe) // preenchendo lista a cada iteração
	}
	fmt.Println("Listagem deu certo!") // log que informa que essa parte geral deu certo
	return res, nil                    // retornando resposta do tipo []modelApresentacao.ReqEquipe
}

// A função ListarMembros será usada apenas pela função BuscarEquipe, e vai servir para exibir os membros da equipe buscada
func (postgres *DBEquipes) ListarMembros(id string) ([]modelPessoa.ReqPessoa, error) {
	sqlStatement := `SELECT * FROM pessoas WHERE equipe_id = $1 ORDER BY id_pessoa` // O comando vai selecionar as pessoas q forem da equipe especificada pelo id
	var pessoa = modelPessoa.ReqPessoa{} // Estrutura individual de pessoa
	var res = []modelPessoa.ReqPessoa{} // Lista que será preenchida e exibida no final

	// Executando sql e recebendo as linhas ou o erro
	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}

	// Percorrendo linhas retornadas do banco de dados e escaneando uma por uma na variável pessoa para adicionar na lista res
	for rows.Next() {
		if err := rows.Scan(&pessoa.ID_Pessoa, &pessoa.Nome_Pessoa, &pessoa.Funcao_Pessoa,
			&pessoa.EquipeID, &pessoa.Favoritar, &pessoa.DataContratacao, &pessoa.UpdatedAt); err != nil {
			return nil, err
		}
		res = append(res, pessoa)
	}

	// retornando resposta e erro nulo
	return res, nil
}

func (postgres *DBEquipes) BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error) {
	sqlStatement := `SELECT * FROM equipes WHERE id_equipe = $1` // Selecionando uma única equipe
	var equipe = &modelApresentacao.ReqEquipe{} // Modelo que será preenchido e exibido no final

	pessoas, err := postgres.ListarMembros(id) // Chamando função para listar os membros da equipe especificada

	// Testando erros da função chamada
	if err != nil {
		if err == sql.ErrNoRows {
			pessoas = nil
		} else {
			return nil, err
		}
	}

	// Preenchendo campo de pessoas com o resultado da listagem de membros da equipe
	equipe.Pessoas = &pessoas

	// executando sql e gravando resposta dentro da estrutura equipe
	row := postgres.DB.QueryRow(sqlStatement, id)
	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.CreatedAt, &equipe.UpdatedAt); err != nil {
		return nil, err
	}
	fmt.Println("Busca deu certo!")
	// retornando resposta e erro nulo
	return equipe, nil
}

func (postgres *DBEquipes) DeletarEquipe(id string) error {
	if _, err := postgres.BuscarEquipe(id); err != nil {
		return err
	} else {
		sqlStatement := `DELETE FROM equipes WHERE id_equipe = $1`

		_, err := postgres.DB.Exec(sqlStatement, id)
		if err != nil {
			return err
		}

		fmt.Println("Delete deu certo!")
		return nil
	}
}

func (postgres *DBEquipes) AtualizarEquipe(id string, req *modelData.Equipe) (*modelApresentacao.ReqEquipe, error) {
	sqlStatement := `UPDATE equipes SET nome_equipe = $1::VARCHAR(80) 
	WHERE id_equipe = $2 RETURNING *`
	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Equipe, id)

	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.CreatedAt, &equipe.UpdatedAt); err != nil {
		return nil, err
	}

	return equipe, nil
}

func (postgres *DBEquipes) ProjetosEquipe(id string) ([]modelProjetos.ReqProjeto, error) {
	sqlStatement := `SELECT * FROM projetos WHERE equipe_id = $1 ORDER BY id_projeto`
	var projeto = modelProjetos.ReqProjeto{}
	var res = []modelProjetos.ReqProjeto{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}

	for rows.Next(){
		if err := rows.Scan(&projeto.ID_Projeto, &projeto.Nome_Projeto, &projeto.Descricao_Projeto,
		&projeto.EquipeID, &projeto.Status, &projeto.DataInicio, &projeto.UpdatedAt, &projeto.DataConclusao); err != nil {
			return nil, err
		}
		res = append(res, projeto)
	}

	return res, nil
}