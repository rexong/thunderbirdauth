package handlers

import (
	"log"
	"net/http"
)

func (u *UserHandler) AuthMiddleware(realm string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		username, password, ok := r.BasicAuth()
		log.Println(authHeader)
		log.Println("Username: ", username)
		log.Println("Password: ", password)
		log.Println("ok:", ok)
		next.ServeHTTP(w, r)
	}
}
