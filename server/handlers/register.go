package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"thunderbirdauth/server/models"
)

type RegisterRequest struct {
	models.UserCredential
}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Method %s not allowed for path %s", r.Method, r.URL.Path)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Invalid JSON")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if requestBody.Username == "" || requestBody.Password == "" {
		log.Printf("Missing Username or Password")
		http.Error(w, "Missing Username Or Password", http.StatusBadRequest)
		return
	}

	//insert user
	userModel := u.UserModel
	_, err = userModel.Create(&requestBody.UserCredential)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered"))
}
