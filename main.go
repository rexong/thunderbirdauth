package main

import (
	"log"
	"thunderbirdauth/db"

	_ "github.com/mattn/go-sqlite3"
)

const DB_PATH = "db/app.db"

func main() {
	db, err := db.Initialise(DB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Database setup complete.")

}
