package routes

import (
	"log"
	"net/http"

	"github.com/Babiel09/Go-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/caes", controllers.Cachorros)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
