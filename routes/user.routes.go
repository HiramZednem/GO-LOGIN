package routes

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}
