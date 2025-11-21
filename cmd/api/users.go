package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"thunderbird.zap/idp/internal/utils"
)

var ErrEmptyCredentials = errors.New("Username or Password is Empty")
var ErrNoRedirectUrl = errors.New("Redirect URL not Provided")
var ErrInvalidCredentials = errors.New("Invalid Credentials Provided")

type templateData struct {
	Error string
}

type EmptyData struct{}

func newTemplateData(err error) templateData {
	if err == nil {
		return templateData{Error: ""}
	}
	return templateData{Error: err.Error()}
}

func GetPageHandlerFunc(htmlPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		t, err := template.ParseFiles(htmlPath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
		}
		t.Execute(w, newTemplateData(nil))
	}
}

func (a *application) createUserHandlerFunc() http.HandlerFunc {
	return GetPageHandlerFunc("./assets/templates/user.register.html")
}

func (a *application) storeUserHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/templates/user.register.html")
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
		log.Println("User Conflict")
		w.WriteHeader(http.StatusConflict)
		t.Execute(w, newTemplateData(err))
		return
	}
	t, err = template.ParseFiles("./assets/templates/user.register.success.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	t.Execute(w, &EmptyData{})
}

func (a *application) loginUserHandlerFunc() http.HandlerFunc {
	return GetPageHandlerFunc("./assets/templates/user.login.html")
}

func (a *application) verifyUserHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/templates/user.login.html")
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
	ok, err := a.store.Users.Verify(username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		t.Execute(w, newTemplateData(ErrInvalidCredentials))
		return
	}
	redirectURL := r.URL.Query().Get("redirect_url")
	if redirectURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, newTemplateData(ErrNoRedirectUrl))
		return
	}
	sessionToken := a.sessionManager.IssueSessionToken()
	sessionExpiry := a.sessionManager.GetSessionExpiryByToken(sessionToken)
	cookie := utils.CreateCookies(sessionToken, sessionExpiry)
	http.SetCookie(w, cookie)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
