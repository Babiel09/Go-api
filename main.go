package main

import (
	"fmt"

	"github.com/Babiel09/Go-api/models"
	"github.com/Babiel09/Go-api/routes"
)

func main() {
	models.Cao = []models.Caes{
		{Id: 1, Nome: "Chacal", Historia: "Esse foi o meu terceiro cachorro"},
		{Id: 2, Nome: "Hiena", Historia: "Essa foi a minha primeira cadela"},
	}
	fmt.Println("API localizada em: http://localhost:8000")
	routes.HandleRequest()
}
