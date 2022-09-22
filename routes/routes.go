package routes

import (
	"log"
	"net/http"

	"github.com/Agnerft/pei-rest.git/controller"
)

func HandleRequest() {
	http.HandleFunc("/", controller.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
