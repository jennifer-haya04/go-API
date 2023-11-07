package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jennifer-haya04/go-API.git/db"
	"github.com/jennifer-haya04/go-API.git/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	//w.Write([]byte("User registred"))
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}
