package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	database := u.App.GetDB()
	_, err = database.Exec("INSERT INTO users (username, password) VALUES (?, ?)", requestBody.Username, requestBody.Password)
	if err != nil {
		log.Printf("Username already exists")
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered"))
}
