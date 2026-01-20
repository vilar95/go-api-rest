package main

import (
	"fmt"
	"go-api-rest/models"
	"go-api-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{ID: 1, Name: "Albert Einstein", History: "Físico teórico alemão, conhecido por desenvolver a teoria da relatividade."},
		{ID: 2, Name: "Marie Curie", History: "Cientista polonesa-francesa, pioneira no estudo da radioatividade."},
		{ID: 3, Name: "Isaac Newton", History: "Físico e matemático inglês, formulador das leis do movimento e da gravitação universal."},
		{ID: 4, Name: "Ada Lovelace", History: "Matemática inglesa, considerada a primeira programadora de computadores."},
		{ID: 5, Name: "Nikola Tesla", History: "Inventor e engenheiro elétrico sérvio-americano, conhecido por suas contribuições ao desenvolvimento da corrente alternada."}}

	fmt.Println("Servidor iniciado na porta 8000")
	routes.HandleRequest()
}
