package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"thunderbirdauth/server/utils"
)

var ErrUsernameExists = errors.New("Username already exists")

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

	userModel := &UserModel{DB: database}
	log.Println("User Model Initialised")
	if !utils.ShouldSeed() {
		return userModel, nil
	}
	log.Println("In DEV environment, seeding user...")
	err = seedUser(userModel)
	if err != nil {
		return nil, err
	}
	log.Println("User Table Seeded")
	return userModel, nil
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

func seedUser(userModel *UserModel) error {
	user := &UserCredential{
		UserBase: UserBase{
			Username: "alice",
		},
		Password: "1234",
	}
	_, err := userModel.Create(user)
	if err != nil {
		if errors.Is(err, ErrUsernameExists) {
			return nil
		}
		return err
	}
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
			return nil, ErrUsernameExists
		}
		return nil, fmt.Errorf("Database error: %v", err)
	}
	log.Println("User Created")

	return &UserBase{
		Username: user.Username,
	}, nil
}

func (u *UserModel) Verify(user *UserCredential) (*UserBase, bool) {
	log.Println("Verifing User...")
	var existingUser UserCredential
	database := u.DB
	query := "SELECT username, password FROM users WHERE username=$1"

	log.Println(user.Username)
	err := database.QueryRow(query, user.Username).Scan(&existingUser.Username, &existingUser.Password)
	if err != nil {
		log.Println("User does not exist.")
		return nil, false
	}
	ok := utils.CheckPasswordHash(user.Password, existingUser.Password)
	if !ok {
		log.Println("Incorrect Password")
		return nil, false
	}
	return &existingUser.UserBase, true
}
