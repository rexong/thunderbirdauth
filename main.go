package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"thunderbirdauth/server"
	"thunderbirdauth/server/handlers"
	ldapserver "thunderbirdauth/server/ldap"
	"thunderbirdauth/server/utils"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB_PATH = utils.GetEnv("DB_PATH", "db/app.db")
	PORT    = utils.GetEnv("PORT", "8080")
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println(os.Getenv("BASIC_USERNAME"))
	app, userModel := server.InitialiseApp(DB_PATH)
	defer app.Close()

	userhandler := &handlers.UserHandler{UserModel: userModel}

	// Check for LDAP Store
	if ldapserver.Manager.StoreExist() {
		log.Println("LDAP Store Exist")
		err := ldapserver.Manager.StartServer()
		if err != nil {
			log.Printf("Unable to Start LDAP Server: %v", err)
			return
		}
	}

	http.HandleFunc("/register", userhandler.Register)
	http.HandleFunc("/auth", userhandler.Authenticate(app.SM, false))
	http.HandleFunc("/auth/basic", userhandler.Authenticate(app.SM, true))
	http.HandleFunc("/login", userhandler.Login(app.SM))
	http.HandleFunc("/ldap", userhandler.ControlLdap)

	addr := fmt.Sprintf(":%s", PORT)
	log.Println("Server listening to port", PORT)
	if err = http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Server Error:", err)
	}
}
