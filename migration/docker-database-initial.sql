create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES
('Chacal', 'Esse foi o meu terceiro cachorro'),
('Hiena', 'Essa foi a minha primeira cadela');