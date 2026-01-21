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
	database.DB.Order("id ASC").Find(&p)

	json.NewEncoder(w).Encode(p)

}

func GetPersonalityByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	p := models.Personality{}
	result := database.DB.First(&p, id)
	if result.Error != nil {
		http.Error(w, "Personalidade não encontrada", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	p := models.Personality{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	result := database.DB.Create(&p)
	if result.Error != nil {
		http.Error(w, "Erro ao criar personalidade", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result := database.DB.Delete(&models.Personality{}, id)
	if result.Error != nil {
		http.Error(w, "Erro ao deletar personalidade", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	p := models.Personality{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	p.ID = 0 // Garantir que o ID não seja alterado
	result := database.DB.Model(&models.Personality{}).Where("id = ?", id).Updates(p)
	if result.Error != nil {
		http.Error(w, "Erro ao atualizar personalidade", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(p)
}
