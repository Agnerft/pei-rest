package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Agnerft/pei-rest.git/database"
	"github.com/Agnerft/pei-rest.git/models"
	"github.com/gorilla/mux"
)

var (
	personalidade models.Personalidade
	err           error
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
	var personalidade models.Personalidade
	database.DB.First(&personalidade, idParam)

	json.NewEncoder(w).Encode(personalidade)

}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	//json.NewEncoder(w).Encode(novaPersonalidade)
	database.DB.Create(&novaPersonalidade)

	fmt.Println("Nome: " + novaPersonalidade.Nome)
	fmt.Println("Historia: " + novaPersonalidade.Historia)

}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]

	var deletaPersonalidade models.Personalidade
	database.DB.Delete(&deletaPersonalidade, idParam)
	fmt.Println("Id deletado: ", deletaPersonalidade.ID)

}
