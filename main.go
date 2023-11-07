package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jennifer-haya04/go-API.git/db"
	"github.com/jennifer-haya04/go-API.git/models"
	"github.com/jennifer-haya04/go-API.git/routes"
	"github.com/rs/cors"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/registerUser", routes.RegisterUser).Methods("POST")

	handler := cors.Default().Handler(r)

	http.ListenAndServe(":3001", handler)
}
