package models

type Caes struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

// Criando um slice para armazenar esses conteúdos
var Cao []Caes
