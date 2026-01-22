package router

import (
	"go-api-rest/internal/handler"
	"go-api-rest/internal/middleware"

	"github.com/gorilla/mux"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(personalityHandler *handler.PersonalityHandler) *mux.Router {
	r := mux.NewRouter()

	// Middlewares globais
	r.Use(middleware.Recovery)
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)
	r.Use(middleware.ContentTypeJSON)

	// Rotas da API
	r.HandleFunc("/", personalityHandler.Home).Methods("GET")

	// Rotas de personalidades
	api := r.PathPrefix("/api/personalities").Subrouter()
	api.HandleFunc("", personalityHandler.Create).Methods("POST")
	api.HandleFunc("", personalityHandler.GetAll).Methods("GET")
	api.HandleFunc("/{id:[0-9]+}", personalityHandler.GetByID).Methods("GET")
	api.HandleFunc("/{id:[0-9]+}", personalityHandler.Update).Methods("PUT")
	api.HandleFunc("/{id:[0-9]+}", personalityHandler.Delete).Methods("DELETE")

	return r
}
