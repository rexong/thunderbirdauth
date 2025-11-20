package main

import (
	"log"

	"thunderbird.zap/idp/internal/configuration"
	"thunderbird.zap/idp/internal/database"
	"thunderbird.zap/idp/internal/store"
)

func main() {
	config := configuration.Init()
	addr := config.AppConfig.DbPath()
	db, err := database.NewSqlite(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := store.NewStorage(db)
	seed(store)
}

func seed(store store.Storage) {
	username := "home"
	password := "1234"
	err := store.Users.Create(username, password)
	if err != nil {
		log.Printf("Error: Unable to Create User: %v", err)
	}
	log.Println("Users Successfully Seeded")
}
