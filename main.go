package main

import (
	"fmt"

	"github.com/Babiel09/Go-api/database"
	"github.com/Babiel09/Go-api/models"
	"github.com/Babiel09/Go-api/routes"
)

func main() {
	models.Cao = []models.Caes{
		{Id: 1, Nome: "Cao1", Historia: "Cao1"},
		{Id: 2, Nome: "Cao2", Historia: "Cao2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("API localizada em: http://localhost:8000")
	routes.HandleRequest()
}
