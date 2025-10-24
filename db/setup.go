package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Initialise(dbPath string) (*sql.DB, error) {
	log.Println("Initialising Database...")
	db, err := connectDatabase(dbPath)
	if err != nil {
		return nil, fmt.Errorf("Database Error: %v", err)
	}

	err = createTable(db)
	if err != nil {
		return nil, fmt.Errorf("Database Error: %v", err)
	}

	log.Println("Database setup complete.")
	return db, nil
}

func connectDatabase(dbPath string) (*sql.DB, error) {
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

func createTable(db *sql.DB) error {
	log.Println("Creating Table 'users' in database")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	log.Println("Table 'users' in database")

	return nil
}
