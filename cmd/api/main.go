package main

import (
	"log"

	"thunderbird.zap/idp/internal/auth/http"
	"thunderbird.zap/idp/internal/auth/ldap"
	"thunderbird.zap/idp/internal/configuration"
	"thunderbird.zap/idp/internal/database"
	"thunderbird.zap/idp/internal/store"
)

func main() {
	config := configuration.Init()

	dbAddr := config.AppConfig.DbPath()
	db, err := database.NewSqlite(dbAddr)
	if err != nil {
		log.Fatalf("Unable to Secure Database Connection: %v", err)
	}
	defer db.Close()

	ldapManager, err := ldap.New(config.LdapConfig)
	if err != nil {
		log.Fatalf("Unable to Start LDAP Server: %v", err)
	}
	defer ldapManager.Close()

	store := store.NewStorage(db)
	sessionManager := http.NewSessionManager()
	app := &application{
		config:         config,
		store:          store,
		sessionManager: sessionManager,
		ldapManager:    ldapManager,
	}
	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Fatal("Unable to Start Server")
	}
}
