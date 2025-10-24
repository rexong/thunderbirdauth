package main

import (
	"fmt"
	"log"
	"net/http"
	"thunderbirdauth/handlers"
	"thunderbirdauth/server"

	_ "github.com/mattn/go-sqlite3"
)

const DB_PATH = "db/app.db"
const PORT = 8080

func main() {
	app := server.InitialiseApp(DB_PATH)
	defer app.Close()

	userhandler := &handlers.UserHandler{App: app}

	http.HandleFunc("/register", userhandler.Register)

	addr := fmt.Sprintf(":%d", PORT)
	err := http.ListenAndServe(addr, nil)
	log.Println("Server listening to port", PORT)
	if err != nil {
		log.Fatal("Server Error:", err)
	}
}
