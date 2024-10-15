package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Produtos struct {
	Nome, Descricao, Preco string
	Quantidade             int
}

var tpm = template.Must(template.ParseGlob("templates/*.html")) //Eu queor tudo que seja .html

func main() {
	fmt.Println("Escutando a porta 8000")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) { //Esse carinha precisa de receber os prâmetros para ler e para escrever
	produtos := []Produtos{
		{"Nootebook", "Muito bom", "R$ 789,00", 2},
	}
	tpm.ExecuteTemplate(w, "Index", produtos)
}
