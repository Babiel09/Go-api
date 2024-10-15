package models

type Caes struct {
	Nome     string `json: "nome"`
	Historia string `json: "historia` //Dentro do json na página irá aparecer o "nome": "Nome do cão" e "historia": "Historia do cão"
}

// Criando um slice para armazenar esses conteúdos
var Cao = []Caes{}
