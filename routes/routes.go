package routes

import (
	"log"
	"net/http"

	"github.com/Agnerft/pei-rest.git/controller"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/personalidades", controller.TodasPersonalidades)
	r.HandleFunc("/api/personalidades/{nome}/", controller.TodasPersonalidades)
	r.HandleFunc("/", controller.Home)
	log.Fatal(http.ListenAndServe(":8000", r))
}
