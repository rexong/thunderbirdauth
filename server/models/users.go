package models

import (
	"database/sql"
	"fmt"
	"log"
	"thunderbirdauth/server/utils"
)

type UserModel struct {
	DB *sql.DB
}

type UserBase struct {
	Username string `json:"username"`
}

type UserCredential struct {
	UserBase
	Password string
}

func InitialiseUserModel(database *sql.DB) (*UserModel, error) {
	log.Println("Initialising User Model...")

	err := createTable(database)
	if err != nil {
		return nil, err
	}

	log.Println("User Model Initialised")
	return &UserModel{DB: database}, nil
}

func createTable(database *sql.DB) error {
	log.Println("Creating Table 'users' in database")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`
	_, err := database.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	log.Println("Table 'users' in database")

	return nil

}

func (u *UserModel) Create(user *UserCredential) (*UserBase, error) {
	log.Println("Creating user...")
	hashpassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	database := u.DB
	_, err = database.Exec(
		"INSERT INTO users (username, password) VALUES (?, ?)",
		user.Username,
		hashpassword,
	)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" {
			return nil, fmt.Errorf("Username already exists")
		}
		return nil, fmt.Errorf("Database error: %v", err)
	}
	log.Println("User Created")

	return &UserBase{
		Username: user.Username,
	}, nil
}
