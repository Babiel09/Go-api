package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Babiel09/Go-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Cachorros(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Cao)
}
