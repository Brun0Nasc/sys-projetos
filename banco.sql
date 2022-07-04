create table equipes(
    id_equipe serial primary key not null,
    nome_equipe varchar(50)
)

create table pessoas(
    id_pessoa serial primary key not null,
    nome_pessoa varchar(50),
    funcao_pessoa varchar(50),
    equipe_id int,
    data_contratacao date default current_date,
    foreign key (equipe_id) references equipes(id_equipe) on delete set null on update cascade
)

create table projetos(
    id_projeto serial primary key not null, 
    nome_projeto varchar(50) not null,
    equipe_id int,
    status varchar(20) default 'Em planejamento',
    data_inicio date default current_date,
    foreign key (equipe_id) references equipes(id_equipe) on delete set null on update cascade
)

create table tasks(
    id_task serial primary key not null,
    descricao_task varchar(50),
    pessoa_id int,
    projeto_id int,
    status varchar(20) default 'A Fazer',
    foreign key (pessoa_id) references pessoas(id_pessoa) on delete set null on update cascade,
    foreign key (projeto_id) references projetos(id_projeto) on delete set null on update cascade
)