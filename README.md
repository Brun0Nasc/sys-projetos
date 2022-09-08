![Badge em Desenvolvimento](http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge)
<br>
## SISTEMA GERENCIADOR DE PROJETOS - BACKEND

Sistema que mantém projetos. Cada projeto recebe uma equipe, e cada equipe pode receber um número não delimitado de pessoas. Cada projeto tem tasks,
cada task pode ser atribuída a uma determinada pessoa que está na equipe do projeto. <br><br>
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Heroku](https://img.shields.io/badge/heroku-%23430098.svg?style=for-the-badge&logo=heroku&logoColor=white)
![cypress](https://img.shields.io/badge/-cypress-%23E5E5E5?style=for-the-badge&logo=cypress&logoColor=058a5e)

 Backend desenvolvido por `Bruno do Nascimento de Brito`

## Índice 

* [Detalhes](#detalhes)
* [Equipe](#equipe)
* [Andamento do Projeto](#andamento-do-projeto)
* [Rotas](#rotas)
* [Equipes](#equipes)
* [Pessoas](#pessoas)
* [Projetos](#projetos)
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
| /equipes/:id/ | | | |
| /equipes/:id/projetos | | | |

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

`/equipes/1/projetos/` Retorna um *GET* dos projetos da equipe especificada:
```json
[
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
]
```

`/equipes/` Para dar *POST* em uma nova equipe, a estrutura será a seguinte:
```json
{
	"nome_equipe":"Equipe 1"
}
```

`/equipes/1/` Para atualizar equipes através do *PUT*, o único dado mutável será o nome da equipe:
```json
{
	"nome_equipe":"Equipe 1"
}
```

`equipes/1/` Para excluir uma equipe através do *DELETE* só será necessário passar o id da equipe na rota, a resposta será a seguinte:
```json 
{
		"message":"Cadastro deletado com sucesso"
} 
```

<hr/>

### PESSOAS
| GET | POST | PUT | DELETE |
|-----|------|-----|--------|
| /pessoas/ | /pessoas/ | /pessoas/:id/ | /pessoas/:id/ |
| /pessoas/:id/ | | /pessoas/:id/favoritar/ | |

`/pessoas/` A função de *GET* geral retorna os dados básicos de pessoas:
```json
{
	"id_pessoa": 1,
	"nome_pessoa": "Tony Stark",
	"funcao_pessoa": "Homem de Ferro",
	"equipe_id": 1,
	"favoritar": 0,
	"data_contratacao": "2022-09-06T17:16:14.553395Z",
	"updated_at": "2022-09-06T17:16:14.553395Z"
}
```

`/pessoas/5/` O *GET* específico retorna, além dos dados básicos de pessoas, as tasks da pessoa especificada:
```json
{
	"id_pessoa": 5,
	"nome_pessoa": "Reed Richards",
	"funcao_pessoa": "Homem Elástico",
	"equipe_id": 2,
	"tasks": [
		{
			"id_task": 1,
			"descricao_task": "Precisamos de bravos heróis, para auxiliar na luta contra o mal.",
			"comentario": "",
			"nivel": "dificil",
			"pessoa_id": 5,
			"projeto_id": 2,
			"status": "Em desenvolvimento",
			"created_at": "2022-09-08T10:52:10.952152Z",
			"updated_at": "2022-09-08T11:00:49.724382Z"
		},
		{
			"id_task": 2,
			"descricao_task": "encontrar  heróis ",
			"comentario": "",
			"nivel": "dificil",
			"pessoa_id": 5,
			"projeto_id": 2,
			"status": "Concluído",
			"created_at": "2022-09-08T10:53:35.741875Z",
			"updated_at": "2022-09-08T11:00:33.07067Z"
		},
		{
			"id_task": 3,
			"descricao_task": "matar o inimigo",
			"comentario": "",
			"nivel": "dificil",
			"pessoa_id": 5,
			"projeto_id": 2,
			"status": "Em planejamento",
			"created_at": "2022-09-08T10:56:44.057123Z",
			"updated_at": "2022-09-08T10:56:44.057123Z"
		}
	],
	"favoritar": 0,
	"data_contratacao": "2022-09-06T17:16:14.553395Z",
	"updated_at": "2022-09-06T17:16:14.553395Z"
}
```

`pessoas/` Para dar *POST* e adicionar uma nova pessoa, a seguinte estrutura é necessária:
```json
{
	"nome_pessoa":"Bruce Banner",
	"funcao_pessoa":"Hulk",
	"equipe_id":1
}
```

`pessoas/21/` Para atualizar uma pessoa através do *PUT*, os dados serão os mesmos que o *POST*:
```json
{
	"nome_pessoa":"Bruce Banner",
	"funcao_pessoa":"Professor Hulk",
	"equipe_id":1
}
```

`pessoas/1/favoritar/` Para favoritar uma pessoa, nenhum dado precisará ser passado, apenas a rota deve ser acessada com o método *PUT*, se a pessoa estiver favoritada, o câmbio entre favoritado e desfavoritado é automático.
```json
{
	"message": "Pessoa favoritada!"
}
```

`pessoas/1/` Para deletar uma pessoa através do método *DELETE*, apenas o id será necessário na rota, a resposta, caso dê certo, é a seguinte:
```json
{
	"message": "Cadastro deletado com sucesso!"
}
```
<hr/>

### PROJETOS

```
https://sistema-aprendizes-brisanet-go.herokuapp.com/projetos
```

| GET | POST | PUT | DELETE |
|-----|------|-----|--------|
| /projetos/ | /projetos/ | /projetos/:id/ | /projetos/:id/ |
| /projetos/:id/ | | /projetos/:id/status/ | |

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

### TASKS
```
https://sistema-aprendizes-brisanet-go.herokuapp.com/tasks
```
A parte de tasks funciona com alguns detalhes adicionais:
- Se uma task for cadastrada em um projeto que ja está com status "Concluído", o status do projeto volta automáticamente para "Em desenvolvimento";
- Se uma task já concluída for atualizada para status de "Fazendo" ou "A fazer" em um projeto já Concluído, o projeto muda automáticamente para "Em desenvolvimento"
- Uma task só pode ser atribuída a uma pessoa que está na equipe responsável pelo projeto que a task está sendo cadastrada. 

| GET | POST | PUT | DELETE |
|-----|------|-----|--------|
| /tasks/ | /tasks/ | /tasks/:id/ | /tasks/:id/ |
| /tasks/:id/ | | /tasks/:id/status/ | |

`/tasks/` `/tasks/1/` As funções *GET* retornam as seguintes informações de tasks:
```json
{
	"id_task": 1,
	"descricao_task": "Precisamos de bravos heróis, para auxiliar na luta contra o mal.",
	"comentario": "",
	"nivel": "dificil",
	"pessoa_id": 5,
	"projeto_id": 2,
	"status": "Em desenvolvimento",
	"created_at": "2022-09-08T10:52:10.952152Z",
	"updated_at": "2022-09-08T11:00:49.724382Z"
}
```

`/tasks/` `/tasks/:id/` Para dar *POST* ou *PUT* em tasks, os seguintes dados serão necessários:
```json
{
	"descricao_task":"Task teste",
	"comentario":"É um teste",
	"nivel":"facil",
	"pessoa_id":2,
	"projeto_id":1
}
```

`/tasks/5/status/` Para atualizar o status da task, através do *PUT* deverá ser passado o seguinte JSON:
```json
{
	"status":"Fazendo"
}
```

`/tasks/5/` Para deletar uma task, através do *DELETE* só será necessário acessar a rota com o id da task que deverá ser excluída:
```json
{
	"message": "Cadastro deletado com sucesso"
}
```
