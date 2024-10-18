package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Babiel09/Go-api/database"
	"github.com/Babiel09/Go-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Cachorros(w http.ResponseWriter, r *http.Request) {
	var caoDB []models.Caes          //Armazeno em uma "var" um Array com todos os Tipos da Struct Caes
	database.DB.Find(&caoDB)         //Uso o gorm para achar a "var" que almejo.
	json.NewEncoder(w).Encode(caoDB) //Transformo ela em json
}

func PegarCao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]                  //Salvo o id do json detro da variável id
	for _, caes := range models.Cao { //Dou uma range em todos os cães no slice de ca~es
		if strconv.Itoa(caes.Id) == id { //Transformo em string para fazer o "check"
			json.NewEncoder(w).Encode(caes) //Transformo em json o determinado cão com o id passado
		}
	}

}
