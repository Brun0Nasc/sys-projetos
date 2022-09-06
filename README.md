![Badge em Desenvolvimento](http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge)
<br>
## SISTEMA GERENCIADOR DE PROJETOS - BACKEND

Sistema que mantém projetos. Cada projeto recebe uma equipe, e cada equipe pode receber um número não delimitado de pessoas. Cada projeto tem tasks,
cada task pode ser atribuída a uma determinada pessoa que está na equipe do projeto. <br>
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)
![Heroku](https://img.shields.io/badge/heroku-%23430098.svg?style=for-the-badge&logo=heroku&logoColor=white)
![cypress](https://img.shields.io/badge/-cypress-%23E5E5E5?style=for-the-badge&logo=cypress&logoColor=058a5e)

## Índice 

* [Detalhes](#detalhes)
* [Equipe](#equipe)
* [Andamento do Projeto](#andamento-do-projeto)
* [Rotas](#rotas)
* [Projetos](#projetos)
* [Pessoas](#pessoas)
* [Equipes](#equipes)
* [Tasks](#tasks)

## Detalhes

- Utilizando linguagem `Go` com `Gin` para desenvolvimento, e o software `Insomnia` para testes;
- Disponível no `Heroku`: https://sistema-aprendizes-brisanet-go.herokuapp.com/
- Front-End em `React`(Em desenvolvimento)
- Detalhes para login: 
```json
{
  "email":"admin",
  "senha":"admin123"
}
```

## Equipe:
- <a href="https://github.com/Brun0Nasc"> Bruno do Nascimento</a>: `Reestruturação da API` `Implementação do Banco de Dados` `Revisão e Reestruturação de Rotas e Funções` `Criação do BD` `Definição das Consultas do BD.`
- <a href="https://github.com/Lucasmartinsn"> Lucas Martins:</a> `Criação do Front-End (em desenvolvimento)`
- <a href="https://github.com/IaraFV"> Iara Ferreira:</a> `Criação do Front-End (em desenvolvimento)`


## Andamento do projeto

| Funcionalidade        | Estado |
| ------------- |:-------------:|
| Manter equipe      | ✔️ |
| Manter projeto      | ✔️ |
| Associar equipe a projeto | ✔️ | 
| Criar tarefa no projeto | ✔️ | 
| Atribuir tarefa | ✔️ | 
| Manter dados no Banco de Dados | ✔️ | 
| Front-End | ⌛ |
| Testes e2e com Cypress | ❌ | 

## ROTAS

### EQUIPES
```
https://sistema-aprendizes-brisanet-go.herokuapp.com/equipes
```

| GET | POST | PUT | DELETE |
|-----|------|-----|--------|
| /equipes | /equipes | /equipes/:id | /equipes/:id |
| /equipes/:id/ | | |

#### DETALHES
`/equipes/` A função de *GET* geral retorna os dados básicos de equipe:
```json
	{
		"id_equipe": 1,
		"nome_equipe": "Os Vingadores",
		"created_at": "2022-09-06T17:15:42.71278Z",
		"updated_at": "2022-09-06T17:15:42.71278Z"
	}
```

`/equipes/1/` A função específica de *GET* retorna junto aos dados de equipe, os membros dela:
```json
{
	"id_equipe": 1,
	"nome_equipe": "Os Vingadores",
	"pessoas": [
		{
			"id_pessoa": 4,
			"nome_pessoa": "Wanda Maximoff",
			"funcao_pessoa": "Feiticeira Escarlate",
			"equipe_id": 1,
			"favoritar": 0,
			"data_contratacao": "2022-09-06T17:16:14.553395Z",
			"updated_at": "2022-09-06T17:16:14.553395Z"
		},
		{
			"id_pessoa": 3,
			"nome_pessoa": "Natasha Romanoff",
			"funcao_pessoa": "Viúva Negra",
			"equipe_id": 1,
			"favoritar": 0,
			"data_contratacao": "2022-09-06T17:16:14.553395Z",
			"updated_at": "2022-09-06T17:16:14.553395Z"
		},
		{
			"id_pessoa": 2,
			"nome_pessoa": "Steve Rogers",
			"funcao_pessoa": "Capitão América",
			"equipe_id": 1,
			"favoritar": 0,
			"data_contratacao": "2022-09-06T17:16:14.553395Z",
			"updated_at": "2022-09-06T17:16:14.553395Z"
		},
		{
			"id_pessoa": 1,
			"nome_pessoa": "Tony Stark",
			"funcao_pessoa": "Homem de Ferro",
			"equipe_id": 1,
			"favoritar": 0,
			"data_contratacao": "2022-09-06T17:16:14.553395Z",
			"updated_at": "2022-09-06T17:16:14.553395Z"
		}
	],
	"created_at": "2022-09-06T17:15:42.71278Z",
	"updated_at": "2022-09-06T17:15:42.71278Z"
}
```

`/equipes/` Para dar *POST* em uma nova equipe, a estrutura será a seguinte:
```json
{
	"nome_equipe":"Equipe 1",
}
```

`/equipes/1/` Para atualizar equipes através do *PUT*, o único dado mutável será o nome da equipe:
```json
{
	"nome_equipe":"Equipe 1",
}
```

`equipes/1/` Para excluir uma equipe através do *DELETE* só será necessário passar o id da equipe na rota, a resposta será a seguinte:
```json 
{
		"message":"Cadastro deletado com sucesso"
} 
```

<hr/>

### PROJETOS

```
https://sistema-aprendizes-brisanet-go.herokuapp.com/projetos
```

| GET | POST | PUT | DELETE |
|-----|------|-----|--------|
| /projetos | /projetos | /projetos/:id | /projetos/:id |
| /projetos/:id | | /projetos/:id/status | |

#### DETALHES
`/projetos/` `/projetos/1/` As funções *GET* retornam os seguintes dados: 
```json 
{
		"id_projeto": 1,
		"nome_projeto": "Guerra Infinita",
		"descricao_projeto": "Salvar universo da extinção",
		"equipe_id": 1,
		"status": "Em planejamento",
		"data_inicio": "2022-09-06T17:20:17.605881Z",
		"updated_at": "2022-09-06T17:20:17.605881Z",
		"data_conclusao": null
} 
```

`/projetos/` Para dar *POST* em um novo projeto, a estrutura será a seguinte:
```json 
{
		"nome_projeto": "Guerra Infinita",
		"descricao_projeto": "Salvar universo da extinção",
		"equipe_id": 1
} 
```

`/projetos/1/` Para dar *PUT* nas informações gerais de um projeto, a estrutura será a mesma que a de *POST*:
```json 
{
		"nome_projeto": "Guerra Infinita",
		"descricao_projeto": "Salvar universo da extinção",
		"equipe_id": 1
} 
```

`/projetos/1/status/` Para dar *PUT* atualizando o status do projeto, a única informação necessária será a de status:
```json 
{
		"status": "Concluído"
} 
```

`/projetos/1/` Para deletar através do *DELETE*, apenas o id do projeto que será deletado precisa ser passado, se for deletado corretamente, a resposta será a seguinte:
```json 
{
		"message":"Cadastro deletado com sucesso"
} 
```
<hr>

<h4>PESSOAS</h4>

```
https://sistema-aprendizes-brisanet-go.herokuapp.com/pessoas
```

<table border=2>
<tr>
 <th>GET</th>
 <th>POST</th>
 <th>PUT</th>
 <th>DELETE</th>
</tr>

<tr>
 <td>/pessoas</td>
 <td>/pessoas</td>
 <td>/pessoas/:id</td>
 <td>/pessoas/:id</td>
</tr>

<tr>
 <td>/pessoas/:id</td>
 <td>/pessoas/:id/tasks</td>
 <td></td>
 <td></td>
</tr>

<tr>
 <td>/pessoas/:id/tasks</td>
 <td></td>
 <td></td>
 <td></td>
</tr>
</table>

<table border=1>
 <tr>
  <th>
   DETALHES 
  </th>
 </tr>
 <tr>
  <td>
   <p>As funções <b>GET</b> retornam as seguintes informações</p>
   <ul>
    <li>ID da pessoa</li>
    <li>Nome da pessoa</li>
    <li>Data de contratação</li>
    <li>Informações da equipe em que ela está</li>
   </ul>
  </td>
 </tr>
 <tr>
  <td>
   <p>As funções <b>POST</b> permitem realizar as seguintes ações:</p>
   <ul>
    <li>Registrar uma nova pessoa no Banco de Dados</li>
    <li>Registrar novas tasks para uma pessoa</li>
   </ul>
   <p>Para cadastrar uma nova pessoa, os seguintes dados são necessários:</p>
   <ul>
    <li>Nome da pessoa</li>
    <li>Função da pessoa</li>
    <li>Equipe que ela fará parte (pode ser registrada sem equipe)</li>
   </ul>
   <p>Para atribuir uma nova task a uma pessoa, será necessário passar apenas a descrição da Task.</p>
   <p>- Informações como ID do Projeto e Status da task serão automáticas. Para que uma task seja atribuída a uma pessoa, essa pessoa precisa estar em uma equipe, e, por sua vez, esta equipe precisa estar atrelada a um projeto que esteja em desenvolvimento.</p>
  </td>
 </tr>
 <tr>
  <td>
   <p>A função <b>PUT</b> permite alterar os seguintes dados:</p>
   <ul>
    <li>Nome da pessoa</li>
    <li>Função da pessoa</li>
    <li>ID da equipe (podendo ser nulo)</li>
   </ul>
  </td>
 </tr>
 <tr>
  <td>
   <p>Para deletar uma pessoa, através do <b>DELETE</b>, será necessário apenas passar o ID da pessoa como parâmetro na rota.</p>
  </td>
 </tr>
</table>

<hr>

<h4>TASKS</h4>

```
https://sistema-aprendizes-brisanet-go.herokuapp.com/tasks
```

<table border=2>
<tr>
 <th>GET</th>
 <th>POST</th>
 <th>PUT</th>
 <th>DELETE</th>
</tr>

<tr>
 <td>/tasks</td>
 <td>/tasks</td>
 <td>/tasks/:id</td>
 <td>/tasks/:id</td>
<tr>

<tr>
 <td>/tasks/:id</td>
 <td></td>
 <td>/tasks/:id/status</td>
 <td></td>
<tr>
</table>
 
 <table border = 1>
  <tr>
    <th>DETALHES</th>
  </tr>

  <tr>
    <td>
      <p>As funções <b>GET</b> retornam as seguintes informações:</p>
      <ul>
        <li>ID da task</li>
        <li>Descrição</li>
        <li>Status</li>
        <li>ID da pessoa responsável pela task</li>
        <li>Nome da pessoa responsável</li>
        <li>ID do projeto que a task está associada</li>
        <li>Nome do projeto</li>
      </ul>
    </td>
  </tr>
  <tr>
    <td>
     <p>Para realizar um novo registro com a função <b>POST</b>, os seguintes dados devem ser informados:</p>
     <ul>
       <li>Descrição da Task</li>
       <li>ID da pessoa que ficará responsável</li>
       <li>ID do projeto associado</li>
     </ul>
     <p>Para atribuir uma task a uma pessoa, essa pessoa precisa estar na equipe que está responsável pela pelo projeto que a task será atribuída, e esse projeto precisa estar com o status "Em desenvolvimento".</p>
    </td>
  </tr>
  <tr>
    <td>
      <p>A função <b>PUT</b> apenas com ID poderá ser usada para alterar informações de Descrição, Pessoa e Projeto.</p>
      <p>O segundo PUT, relacionado ao status, só altera o Status da Task.</p>
    </td>
  </tr>
  <tr>
    <td>
      <p>Para deletar uma Task, através do <b>DELETE</b>, será necessário apenas passar o ID da Task como parâmetro na rota.</p>
    </td>
 </tr>
</table>
