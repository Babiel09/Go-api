package main

import (
	"fmt"

	"github.com/Babiel09/Go-api/routes"
)

func main() {
	fmt.Println("API localizada em: http://localhost:8000")
	routes.HandleRequest()
}
