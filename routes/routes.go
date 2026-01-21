package routes

import (
	"go-api-rest/controllers"
	"go-api-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Lida as requisições HTTP e direciona para os controladores apropriados
func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityByID).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
