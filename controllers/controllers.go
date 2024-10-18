package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Babiel09/Go-api/database"
	"github.com/Babiel09/Go-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Cachorros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var caoDB []models.Caes          //Armazeno em uma "var" um Array com todos os Tipos da Struct Caes
	database.DB.Find(&caoDB)         //Uso o gorm para achar a "var" que almejo.
	json.NewEncoder(w).Encode(caoDB) //Transformo ela em json
}

func PegarCao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]                   //Salvo o id do json detro da variável id
	var caoDbId models.Caes            //Salvo uma "var" contendo a struct de "Caes"
	database.DB.First(&caoDbId, id)    //Peço para ela levar em conta o id e a minha "var"
	json.NewEncoder(w).Encode(caoDbId) //Transformo minha "var" em JSON

}

// Meu primeiro método POST :)
func PostCachorros(w http.ResponseWriter, r *http.Request) {
	var caoPost models.Caes
	//O decoder transforma o "json" em algo que o Go entenda
	json.NewDecoder(r.Body).Decode(&caoPost)
	database.DB.Create(&caoPost)
	json.NewEncoder(w).Encode(caoPost)
	//Depois de postar o novo cão, eu irei transformalo em json
}

// Criando meu primeiro DELETE :)
func DeleteCachorros(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var caoDelete models.Caes

	//Deletando do meu Banco de dados:
	database.DB.Delete(&caoDelete, id)
	//Deletando no json também:
	json.NewEncoder(w).Encode(caoDelete)
}

// Criando meu primeiro PUT :)
func EditarCachorros(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var caoPut models.Caes
	database.DB.First(&caoPut, id)
	json.NewDecoder(r.Body).Decode(&caoPut)
	database.DB.Save(&caoPut)
	json.NewEncoder(w).Encode(caoPut)
}
