package routes

import (
	"log"
	"net/http"

	"github.com/Agnerft/pei-rest.git/controller"
	"github.com/Agnerft/pei-rest.git/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/api/personalidades", controller.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controller.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controller.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controller.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controller.EditarUmaPersonalidade).Methods("Put")
	r.HandleFunc("/", controller.Home)
	log.Fatal(http.ListenAndServe(":8000", r))
}
