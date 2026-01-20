package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-rest/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo à API REST em Go!")
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)

}

func GetPersonalityByID(w http.ResponseWriter, r *http.Request) {
	// Implementar lógica para obter personalidade por ID
}

