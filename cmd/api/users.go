package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Unable to Decode JSON: %v", err)
		return
	}
	if requestBody.Username == "" || requestBody.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Missing Username or Password")
		return
	}
	err = a.store.Users.Create(requestBody.Username, requestBody.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Unable to Create User: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Created!")
}
