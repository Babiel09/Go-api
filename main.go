package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpm = template.Must(template.ParseGlob("templates/*.html")) //Eu queor tudo que seja .html

func main() {
	fmt.Println("Escutando a porta 8000")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) { //Esse carinha precisa de receber os prâmetros para ler e para escrever
	tpm.ExecuteTemplate(w, "Index", nil)
}
