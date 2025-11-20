package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

var ErrEmptyCredentials = errors.New("Username or Password is Empty")

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type templateData struct {
	Error string
}

func newTemplateData(err error) templateData {
	if err == nil {
		return templateData{Error: ""}
	}
	return templateData{Error: err.Error()}
}

func (a *application) createUserHandler(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles("./assets/templates/user.create.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
	}

	t.Execute(w, newTemplateData(nil))
}

func (a *application) storeUserHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/templates/user.create.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, newTemplateData(ErrEmptyCredentials))
		return
	}
	err = a.store.Users.Create(username, password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, newTemplateData(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Created!")
}
