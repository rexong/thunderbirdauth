package ldapserver

import (
	"log"

	"beryju.io/ldap"
)

func Start() {
	loadConfig()
	server := ldap.NewServer()
	store, err := NewStore(config.StorePath)
	if err != nil {
		log.Fatalf("LDAP Server Failed: %s", err.Error())
	}
	server.BindFunc("", store)
	server.SearchFunc("", store)

	log.Printf("Starting LDAP server on %s...", config.ListenAddr)
	if err := server.ListenAndServe(config.ListenAddr); err != nil {
		log.Fatalf("LDAP Server Failed: %s", err.Error())
	}
}
