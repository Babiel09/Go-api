package main

import (
	"fmt"

	"github.com/Babiel09/Go-api/models"
	"github.com/Babiel09/Go-api/routes"
)

func main() {
	models.Cao = []models.Caes{
		{"Chacal", "Esse foi o meu primeiro cachorro"},
	}
	fmt.Println("API localizada em: http://localhost:8000")
	routes.HandleRequest()
}
