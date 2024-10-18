package routes

import (
	"log"
	"net/http"

	"github.com/Babiel09/Go-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/caes", controllers.Cachorros).Methods("Get") //Definindo o método que será usado
	r.HandleFunc("/caes/{id}", controllers.PegarCao).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
