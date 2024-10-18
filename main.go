package main

import (
	"fmt"

	"github.com/Babiel09/Go-api/models"
	"github.com/Babiel09/Go-api/routes"
)

func main() {
	fmt.Println("API localizada em: http://localhost:8000")
	routes.HandleRequest()
	models.Cao = []models.Caes{
		{Id: 1, Nome: "Cao1", Historia: "Cao1"},
		{Id: 2, Nome: "Cao2", Historia: "Cao2"},
	}
}
