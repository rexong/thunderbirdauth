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
	log.Println("Registering user...")
	if r.Method != http.MethodPost {
		log.Printf("Method %s not allowed for path %s", r.Method, r.URL.Path)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RegisterRequest
	log.Println("Decoding JSON Request Body...")
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Invalid JSON")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	log.Println("JSON Request Body Decoded")

	log.Println("Checking if Credentials are missing...")
	if requestBody.Username == "" || requestBody.Password == "" {
		log.Printf("Missing Username or Password")
		http.Error(w, "Missing Username Or Password", http.StatusBadRequest)
		return
	}
	log.Println("Credentials Provided")

	userModel := u.UserModel
	user, err := userModel.Create(&requestBody.UserCredential)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered"))
	log.Println("User ", user.Username, " registered")
}
