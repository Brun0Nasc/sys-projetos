package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	modelData "github.com/Brun0Nasc/sys-projetos/infra/equipes/model"
)

type DBEquipes struct {
	DB *sql.DB
}

func (postgres *DBEquipes) NovaEquipe(req *modelData.Equipe) (*modelApresentacao.ReqEquipe, error) {
	sqlStatement := `INSERT INTO equipes
	(nome_equipe)
	VALUES($1::VARCHAR(80))
	RETURNING *`

	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Equipe)
	
	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.CreatedAt, &equipe.UpdatedAt); err != nil {
		return nil, err
	}

	return equipe, nil
}

func (postgres *DBEquipes) ListarEquipes() ([]modelApresentacao.ReqEquipe, error) {
	sqlStatement := `SELECT * FROM equipes ORDER BY id_equipe` // comando sql
	var res = []modelApresentacao.ReqEquipe{} // lista que vai receber resultados da consulta
	var equipe = modelApresentacao.ReqEquipe{} // estrutura individual que vai ser usada para preencher a lista

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
	return res, nil // retornando resposta do tipo []modelApresentacao.ReqEquipe
}

func (postgres *DBEquipes) BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error){
	sqlStatement := `SELECT * FROM equipes WHERE id_equipe = $1`
	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, id)
	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.CreatedAt, &equipe.UpdatedAt); err != nil {
		return nil, err
	}
	fmt.Println("Busca deu certo!")
	return equipe, nil
}

func (postgres *DBEquipes) DeletarEquipe(id string) error {
	sqlStatement := `DELETE FROM equipes WHERE id_equipe = $1`

	_, err := postgres.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	fmt.Println("Delete deu certo!")
	return nil
}

func (postgres *DBEquipes) AtualizarEquipe(id string, req *modelData.Equipe) (*modelApresentacao.ReqEquipe, error ){
	sqlStatement := `UPDATE equipes SET nome_equipe = $1::VARCHAR(80) 
	WHERE id_equipe = $2 RETURNING *`
	var equipe = &modelApresentacao.ReqEquipe{} 

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Equipe, id)

	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.CreatedAt, &equipe.UpdatedAt); err != nil {
		return nil, err
	}
	
	return equipe, nil
}