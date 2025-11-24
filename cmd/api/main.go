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
	log.Println("Starting Server...")
	log.Println("Initialising Configuration...")
	config := configuration.Init()
	log.Println("Configuration Initialised.")

	log.Println("Initialising Database...")
	dbAddr := config.AppConfig.DbPath()
	db, err := database.NewSqlite(dbAddr)
	if err != nil {
		log.Fatalf("Unable to Secure Database Connection: %v", err)
	}
	defer db.Close()
	store := store.NewStorage(db)
	log.Println("Database Initialised.")

	log.Println("Initialising Session Manager...")
	sessionManager := http.NewSessionManager()
	log.Println("Session Manager Initialised.")

	var ldapManager *ldap.LdapManager
	if config.LdapConfig.ShouldStart() {
		log.Println("Initialising LDAP Server...")
		ldapManager, err = ldap.New(config.LdapConfig, store.Users)
		if err != nil {
			log.Fatalf("Unable to Start LDAP Server: %v", err)
		}
		defer ldapManager.Close()
		log.Println("LDAP Server Initialised.")
	}

	app := &application{
		config:         config,
		store:          store,
		sessionManager: sessionManager,
		ldapManager:    ldapManager,
	}

	mux := app.mount()

	log.Println("Server is up and listening to ", config.AppConfig.Addr())
	if err := app.run(mux); err != nil {
		log.Fatal("Unable to Start Server")
	}
}
