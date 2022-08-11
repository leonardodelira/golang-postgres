package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/leonardodelira/go-postgres/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error on decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]interface{}

	if err != nil {
		resp = map[string]interface{}{
			"Error": true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]interface{}{
			"Error": false,
			"Message": fmt.Sprintf("Todo inserido com sucesso! ID: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}