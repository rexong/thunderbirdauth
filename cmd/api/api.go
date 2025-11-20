package main

import (
	"errors"
	"fmt"
	"net/http"

	auth "thunderbird.zap/idp/internal/auth/http"
	"thunderbird.zap/idp/internal/configuration"
	"thunderbird.zap/idp/internal/store"
)

type application struct {
	config         configuration.Config
	store          store.Storage
	sessionManager auth.SessionManager
}

func welcome(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, "hello world")
}

func (a *application) mount() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", a.createUserHandlerFunc())
	mux.HandleFunc("POST /users", a.storeUserHandler)
	mux.HandleFunc("GET /users/login", a.loginUserHandlerFunc())
	mux.HandleFunc("POST /users/login", a.verifyUserHandler)
	mux.HandleFunc("GET /auth", a.Authenticate(false))
	mux.HandleFunc("GET /auth/basic", a.Authenticate(true))
	mux.HandleFunc("GET /", welcome)

	return mux
}

func (a *application) run(mux http.Handler) error {
	appConfig := a.config.AppConfig
	server := &http.Server{
		Addr:    appConfig.Addr(),
		Handler: mux,
	}

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
