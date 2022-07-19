## API com Banco de Dados implementado e Front-End

Sistema que mantém projetos. Cada projeto recebe uma equipe, e cada equipe pode receber um número não delimitado de pessoas. Cada projeto tem tasks,
cada task pode ser atribuída a uma determinada pessoa que está na equipe do projeto.

## Detalhes

- Utilizando Go e Gin para desenvolvimento, e o software Insomnia para testes;
- Disponível em (HEROKU): https://sistema-aprendizes-brisanet-go.herokuapp.com/
- Front-End (React) em desenvolvimento.
<div style="display: inline_block"><br>
<img align="center" alt="bruno-Golang" height="60" width="70" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" />
<img align="center" alt="bruno-Gin" height="50" width="60" src="https://avatars.githubusercontent.com/u/15729372?s=280&v=4" />
<img align="center" alt="bruno-Postgres" height="50" width="60" src="https://wiki.postgresql.org/images/3/30/PostgreSQL_logo.3colors.120x120.png"/>
<img align="center" alt="bruno-React" height="50" width="60" src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/2300px-React-icon.svg.png"/>
</div>


## Equipe:
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
<hr>
<h4>PROJETOS</h4>

<table border=2>
<tr>
 <th>GET</th>
 <th>POST</th>
 <th>PUT</th>
 <th>DELETE</th>
</tr>
<tr>
  <td>
   <div>/projetos</div>
  </td>
  <td>
   <div>/projetos</div>
  </td>
  <td>
   <div>/projetos/:id</div>
  </td>
  <td>
   <div>/projetos/:id</div>
  </td>
 </tr>
 <tr>
  <td>
   <div>/projetos/:id</div>
  </td>
  <td></td>
  <td>
   <div>/projetos/:id/status</div>
  </td>
  <td></td>
 </tr>
</table>

<table border=1>
<th>
 DETALHES
</th>
<tr>
<td>
<p>As funções <b>GET</b> retornam os seguintes dados:</p>
<ul>
<li>ID do projeto</li>
<li>Nome do projeto</li>
<li>Status do projeto</li>
<li>Data de início</li>
<li>Data de conclusão</li>
<li>Equipe responsável pelo projeto</li>
<li>Tasks atribuídas ao projeto</li>
</ul>
</td>
</tr>

<tr>
<td>
<p>Para realizar um novo registro com a função <b>POST</b>, os seguintes dados deverão ser informados:</p>
<ul>
 <li>Nome do projeto</li>
 <li>ID da equipe que ficará responsável pelo projeto</li>
</ul>
<p>Dados como Status, Data de Inicio e Data de Conclusão são automáticos, todo projeto inicia automáticamente com o status
"Em planejamento". É possível alterar o status de desenvolvimento do projeto, mas apenas um projeto por equipe pode estar
com o status "Em desenvolvimento".</p>
</td>
</tr>

<tr>
<td>
<p>Nas funções <b>PUT</b>, dois tipos de atualização podem ser feitas:</p>
<ul>
 <li>Mudança de nome e/ou equipe do projeto</li>
 <li>Atualização do status do projeto</li>
</ul>
</td>
</tr>

<tr>
<td>
<p>Para deletar um projeto, através do <b>DELETE</b>, será necessário apenas passar o ID do projeto como parâmetro na rota.</p>
</td>
</tr>
</table>
<hr>


<h4>PESSOAS</h4>

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
   <p>Para deletar um projeto, através do <b>DELETE</b>, será necessário apenas passar o ID do projeto como parâmetro na rota.</p>
  </td>
 </tr>
</table>
<hr>
 
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
