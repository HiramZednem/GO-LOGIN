package main

import (
	"net/http"

	"github.com/HiramZednem/GO-LOGIN/db"
	"github.com/HiramZednem/GO-LOGIN/models"
	"github.com/HiramZednem/GO-LOGIN/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/create/user", routes.CreateUser).Methods("POST")
	r.HandleFunc("/get/user/{id}", routes.GetUser).Methods("GET")
	r.HandleFunc("/login", routes.Login).Methods("POST")

	http.ListenAndServe(":3000", r)
}