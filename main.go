package main

import (
	"fmt"

	"github.com/Agnerft/pei-rest.git/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
