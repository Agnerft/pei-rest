package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Agnerft/pei-rest.git/database"
	"github.com/Agnerft/pei-rest.git/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.ID) == idParam {
			json.NewEncoder(w).Encode(personalidade)
		}
	}

}
