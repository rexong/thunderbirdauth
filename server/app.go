package server

import (
	"database/sql"
	"log"
	"thunderbirdauth/db"
	"thunderbirdauth/server/models"
	"thunderbirdauth/server/utils"
)

type App struct {
	DB *sql.DB
	SM *utils.SessionManager
}

func InitialiseApp(database_path string) (*App, *models.UserModel) {
	log.Println("Initialising App...")
	database, err := db.Connect(database_path)
	if err != nil {
		log.Fatal(err)
	}

	userModel, err := models.InitialiseUserModel(database)
	if err != nil {
		log.Fatal(err)
	}

	sessionManager := utils.InitialiseSessionManager()

	log.Println("App Initialised")
	app := &App{
		DB: database,
		SM: sessionManager,
	}
	return app, userModel
}

func (a *App) Close() {
	a.DB.Close()
}
