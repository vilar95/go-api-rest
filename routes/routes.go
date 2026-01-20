package routes

import (
	"go-api-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Lida as requisições HTTP e direciona para os controladores apropriados
func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetPersonalities)
	log.Fatal(http.ListenAndServe(":8000", r))
}
