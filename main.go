package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/routes"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Servidor iniciado na porta http://localhost:8000")
	routes.HandleRequest()
}
