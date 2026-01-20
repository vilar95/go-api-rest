package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-rest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo à API REST em Go!")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)

}

func GetPersonalityByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Faz uma iteração sobre a lista de personalidades para encontrar a personalidade com o ID correspondente
	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.ID) == id {
			// Retorna a personalidade encontrada em formato JSON
			json.NewEncoder(w).Encode(personality)
			return
		}
	}
	http.Error(w, "Personalidade não encontrada", http.StatusNotFound)

}
