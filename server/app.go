package server

import (
	"database/sql"
	"log"
	"thunderbirdauth/db"
)

type App struct {
	DB *sql.DB
}

type AppInterface interface {
	GetDB() *sql.DB
}

func InitialiseApp(database_path string) *App {
	log.Println("Initialising App...")
	database, err := db.Initialise(database_path)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("App Initialised")
	return &App{DB: database}
}

func (a *App) GetDB() *sql.DB {
	return a.DB
}

func (a *App) Close() {
	a.DB.Close()
}
