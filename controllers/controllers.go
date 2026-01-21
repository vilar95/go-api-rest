package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-rest/database"
	"go-api-rest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo à API REST em Go!")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	p := []models.Personality{}
	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)

}

func GetPersonalityByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	p := models.Personality{}
	result := database.DB.First(&p, id)
	if result.Error == nil {
		json.NewEncoder(w).Encode(p)
		return
	}
	http.Error(w, "Personalidade não encontrada", http.StatusNotFound)

}
