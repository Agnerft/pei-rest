package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Agnerft/pei-rest.git/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
