package handlers

import (
	"log"
	"net/http"
	"thunderbirdauth/server/models"
)

func (u *UserHandler) AuthMiddleware(realm string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Unpacking Credentials...")
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		log.Println("Credentials Unpacked, Verifying User...")
		user := models.UserCredential{
			UserBase: models.UserBase{Username: username},
			Password: password,
		}

		_, ok = u.UserModel.Verify(&user)
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
