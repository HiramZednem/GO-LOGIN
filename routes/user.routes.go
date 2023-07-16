package routes

import (
	"encoding/json"
	"net/http"

	"github.com/HiramZednem/GO-LOGIN/db"
	"github.com/HiramZednem/GO-LOGIN/models"
	"github.com/HiramZednem/GO-LOGIN/models/request"
	"github.com/HiramZednem/GO-LOGIN/models/response"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest request.LoginRequest
	json.NewDecoder(r.Body).Decode(&loginRequest)

	var user models.User
	db.DB.Where("email = ?", loginRequest.Email).First(&user)

	response := response.LoginResponse {
		Success: false,
		UserID:  0,
	}

	if user.ID == 0 || user.Password != loginRequest.Password {
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Success = true
	response.UserID = user.ID
	json.NewEncoder(w).Encode(response)
}

