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
		log.Printf("GET %s", htmlPath)
		t, err := template.ParseFiles(htmlPath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
		}
		t.Execute(w, newTemplateData(nil))
		log.Printf("GET %s %d", htmlPath, http.StatusOK)
	}
}

func (a *application) createUserHandlerFunc() http.HandlerFunc {
	return GetPageHandlerFunc("./assets/templates/user.register.html")
}

func (a *application) storeUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("POST User Register Page")
	t, err := template.ParseFiles("./assets/templates/user.register.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		log.Printf("%d %v", http.StatusInternalServerError, err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, newTemplateData(ErrEmptyCredentials))
		log.Printf("%d %v", http.StatusBadRequest, ErrEmptyCredentials)
		return
	}
	err = a.store.Users.Create(username, password)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		t.Execute(w, newTemplateData(err))
		log.Printf("%d %v", http.StatusConflict, err)
		return
	}
	log.Printf("User %s Created", username)
	t, err = template.ParseFiles("./assets/templates/user.register.success.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		log.Printf("%d %v", http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.Execute(w, &EmptyData{})
	log.Printf("POST User Register Page %d", http.StatusOK)
}

func (a *application) loginUserHandlerFunc() http.HandlerFunc {
	return GetPageHandlerFunc("./assets/templates/user.login.html")
}

func (a *application) verifyUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("POST User Login Page")
	t, err := template.ParseFiles("./assets/templates/user.login.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		log.Printf("%d %v", http.StatusInternalServerError, err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, newTemplateData(ErrEmptyCredentials))
		log.Printf("%d %v", http.StatusInternalServerError, err)
		return
	}
	ok, err := a.store.Users.Verify(username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		log.Printf("%d %v", http.StatusInternalServerError, err)
		return
	}
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		t.Execute(w, newTemplateData(ErrInvalidCredentials))
		log.Printf("%d %v", http.StatusUnauthorized, ErrInvalidCredentials)
		return
	}
	log.Printf("User %s Verified", username)
	redirectURL := r.URL.Query().Get("redirect_url")
	if redirectURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, newTemplateData(ErrNoRedirectUrl))
		log.Printf("%d %v", http.StatusBadRequest, ErrNoRedirectUrl)
		return
	}
	log.Println("Issuing Session Token...")
	sessionToken := a.sessionManager.IssueSessionToken()
	sessionExpiry := a.sessionManager.GetSessionExpiryByToken(sessionToken)
	cookie := utils.CreateCookies(sessionToken, sessionExpiry)
	http.SetCookie(w, cookie)
	log.Println("Session Token Issued.")
	log.Printf("%d User Login, Redirecting to %s", http.StatusFound, redirectURL)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
