## API com Banco de Dados implementado e Front-End

Sistema que mantém projetos. Cada projeto recebe uma equipe, e cada equipe pode receber um número não delimitado de pessoas. Cada projeto tem tasks,
cada task pode ser atribuída a uma determinada pessoa que está na equipe do projeto.

## Detalhes

- Utilizando Go e Gin para desenvolvimento, e o software Insomnia para testes;
- Disponível em (HEROKU): https://sistema-aprendizes-brisanet-go.herokuapp.com/
- Front-End (Angular) em desenvolvimento.
<div style="display: inline_block"><br>
<img align="center" alt="bruno-Golang" height="60" width="70" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" />
<img align="center" alt="bruno-Gin" height="50" width="60" src="https://avatars.githubusercontent.com/u/15729372?s=280&v=4" />
<img align="center" alt="bruno-Postgres" height="50" width="60" src="https://wiki.postgresql.org/images/3/30/PostgreSQL_logo.3colors.120x120.png"/>
<img align="center" alt="bruno-React" height="50" width="60" src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/2300px-React-icon.svg.png"/>
</div>


## Membros:
<div>-<a href="https://github.com/Brun0Nasc"> Bruno do Nascimento:</a> Reestruturação da API; Implementação do Banco de Dados; Revisão e Reestruturação de Rotas e Funções; Criação do BDD; Definição das Consultas do BDD.</div>
<div>-<a href="https://github.com/Lucasmartinsn"> Lucas Martins:</a> Estruturação do Front-End (em desenvolvimento)</div> 
<div>-<a href="https://github.com/IaraFV"> Iara Ferreira:</a> Estruturação do Front-End (em desenvolvimento)</div>


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
| Testes 2e2 com Cypress | ❌ | 

### ROTAS

- PROJETOS:
 <div>GET:</div>
 <div>/projetos</div>
 <div>/projetos/:id</div>
 <div>/projetos/:id/tasks</div>
 <div>------------------------------------</div>
 <div> POST: </div>
 <div>/projetos</div>
 <div>------------------------------------</div>
 <div> PUT:</div>
 <div>/projetos/:id</div>
 <div>------------------------------------</div>
 <div>DELETE:</div>
 <div>/projetos/:id</div>
 <div>------------------------------------</div>
 <div>ㅤ </div>
 <div>ㅤ </div>
 
 - PESSOAS: 
 <div>GET:</div>
 <div>/pessoas</div>
 <div>/pessoas/:id</div>
 <div>/pessoas/:id/tasks</div>
 <div>------------------------------------</div>
 <div>POST:</div>
 <div>/pessoas</div>
 <div>------------------------------------</div>
 <div>PUT:</div>
 <div>/pessoas/:id</div>
 <div>------------------------------------</div>
 <div>DELETE:</div>
 <div>/pessoas/:id</div>
 <div>------------------------------------</div>
 <div>ㅤ </div>
 <div>ㅤ </div>
 
- EQUIPES: 
 <div>GET:</div>
 <div>/equipes</div>
 <div>/equipes/:id</div>
 <div>/equipes/:id/pessoas</div>
 <div>/equipes/:id/projetos</div>
 <div>------------------------------------</div>
 <div>POST:</div>
 <div>/equipes</div>
 <div>------------------------------------</div>
 <div>PUT:</div>
 <div>/equipes/:id</div>
 <div>------------------------------------</div>
 <div>DELETE:</div>
 <div>/equipes/:id</div>
 <div>------------------------------------</div>
 <div>ㅤ </div>
 <div>ㅤ </div>
 
- TASKS: 
 <div>GET:</div>
 <div>/tasks</div>
 <div>/tasks/:id</div>
 <div>------------------------------------</div>
 <div>POST:</div>
 <div>/tasks</div>
 <div>------------------------------------</div>
 <div>PUT:</div>
 <div>/tasks/:id</div>
 <div>------------------------------------</div>
 <div>DELETE:</div>
 <div>/taks/:id</div>
 <div>------------------------------------</div>
