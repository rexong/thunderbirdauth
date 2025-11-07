package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"thunderbirdauth/server"
	"thunderbirdauth/server/handlers"
	"thunderbirdauth/server/ldap"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_PATH = "db/app.db"
	PORT    = 8080
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println(os.Getenv("BASIC_USERNAME"))
	app, userModel := server.InitialiseApp(DB_PATH)
	defer app.Close()

	ldapserver.Start()
	userhandler := &handlers.UserHandler{UserModel: userModel}

	http.HandleFunc("/register", userhandler.Register)
	http.HandleFunc("/auth", userhandler.Authenticate(app.SM, false))
	http.HandleFunc("/auth/basic", userhandler.Authenticate(app.SM, true))
	http.HandleFunc("/login", userhandler.Login(app.SM))

	addr := fmt.Sprintf(":%d", PORT)
	err = http.ListenAndServe(addr, nil)
	log.Println("Server listening to port", PORT)
	if err != nil {
		log.Fatal("Server Error:", err)
	}
	log.Println("Starting LDAP")
}
