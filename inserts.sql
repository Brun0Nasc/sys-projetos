insert into equipes (nome_equipe) values('Os Vingadores'), ('Quarteto Fantástico'), ('X-Men'), ('Liga da Justiça');

insert into pessoas (nome_pessoa, funcao_pessoa, equipe_id) values('Tony Stark', 'Homem de Ferro', 1), 
('Steve Rogers', 'Capitão América', 1), ('Natasha Romanoff', 'Viúva Negra', 1), 
('Wanda Maximoff', 'Feiticeira Escarlate', 1), ('Reed Richards', 'Homem Elástico', 2), 
('Sue Storm', 'Mulher Invisível', 2), ('Jhonny Storm', 'Tocha Humana', 2), ('Ben Grimm', 'Coisa', 2), 
('Charles Xavier', 'Professor X', 3), ('Logan', 'Wolverine', 3), ('Scott Summers', 'Ciclope', 3), 
('Jean Grey', 'Fênix', 3), ('Kitty', 'Lince Negra', 3), ('Klark Kent', 'Superman', 4), ('Bruce Wayne', 'Batman', 4), 
('Diana Prince', 'Mulher Maravilha', 4), ('Barry Allen', 'Flash', 4);

insert into projetos (nome_projeto, descricao_projeto, equipe_id) values('Guerra Infinita', 'Salvar universo da extinção', 1), 
('Salvar Terra', 'Salvar Terra do Galactus', 2),('Salvar Humanos', 'Salvar humanos do Magneto', 3),
('Salvar Terra da DC', 'Derrotar Darkseid', 4);