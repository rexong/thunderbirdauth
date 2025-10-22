package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"thunderbirdauth/db"

	_ "github.com/mattn/go-sqlite3"
)

const DB_PATH = "db/app.db"
const PORT = 8080

var database *sql.DB

func main() {
	var err error
	database, err = db.Initialise(DB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	log.Println("Database setup complete.")

	http.HandleFunc("/register", registerHandler)

	addr := fmt.Sprintf(":%d", PORT)
	err = http.ListenAndServe(addr, nil)
	log.Println("Server listening to port", PORT)
	if err != nil {
		log.Fatal("Server Error:", err)
	}
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Method %s not allowed for path %s", r.Method, r.URL.Path)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Invalid JSON")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	if requestBody.Username == "" || requestBody.Password == "" {
		log.Printf("Missing Username or Password")
		http.Error(w, "Missing Username Or Password", http.StatusBadRequest)
	}

	//insert user
	_, err = database.Exec("INSERT INTO users (username, password) VALUES (?, ?)", requestBody.Username, requestBody.Password)
	if err != nil {
		log.Printf("Username already exists")
		http.Error(w, "Username already exists", http.StatusConflict)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered"))
}
