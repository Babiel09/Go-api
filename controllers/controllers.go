package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Produtos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Página de Produtos")
}
