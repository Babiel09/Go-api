package routes

import (
	"log"
	"net/http"

	"github.com/Babiel09/Go-api/controllers"
	"github.com/Babiel09/Go-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/caes", controllers.Cachorros).Methods("Get")
	r.HandleFunc("/caes", controllers.PostCachorros).Methods("Post")
	r.HandleFunc("/caes/{id}", controllers.EditarCachorros).Methods("Put")
	r.HandleFunc("/caes/{id}", controllers.DeleteCachorros).Methods("Delete")
	r.HandleFunc("/caes/{id}", controllers.PegarCao).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
