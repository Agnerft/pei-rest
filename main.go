package main

import (
	"fmt"

	"github.com/Agnerft/pei-rest.git/models"
	"github.com/Agnerft/pei-rest.git/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Nome 1", Historia: "História 1"},
		{ID: 2, Nome: "Nome 2", Historia: "História 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
