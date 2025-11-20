package main

import (
	"errors"
	"fmt"
	"net/http"

	"thunderbird.zap/idp/internal/configuration"
	"thunderbird.zap/idp/internal/store"
)

type application struct {
	config configuration.Config
	store  store.Storage
}

func welcome(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, "hello world")
}

func (a *application) mount() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", a.createUserHandler)
	mux.HandleFunc("POST /users", a.storeUserHandler)
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
