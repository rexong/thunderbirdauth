package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Connect(dbPath string) (*sql.DB, error) {
	log.Println("Connecting to ", dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Connected to SQLite database:", dbPath)
	return db, nil
}
