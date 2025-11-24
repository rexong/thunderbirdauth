package main

import (
	"fmt"
	"log"
	"net/http"

	auth "thunderbird.zap/idp/internal/auth/http"
	"thunderbird.zap/idp/internal/utils"
)

func (a *application) Authenticate(withBasic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Authenticating.... Basic Status: %t", withBasic)
		log.Println("Fetching Cookie...")
		cookie, err := r.Cookie(utils.CookieName)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: %v", err)
			log.Printf("%d %v", http.StatusUnauthorized, err)
			return
		}
		log.Println("Verifying Session Token...")
		ok := a.sessionManager.VerifySessionToken(cookie.Value)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: %v", err)
			log.Printf("%d %v", http.StatusUnauthorized, err)
			return
		}
		if withBasic {
			basicAuthHeader := auth.BasicAuthHeader(a.config.BasicConfig.Credentials())
			log.Printf("Adding Authorization Header %s...", basicAuthHeader)
			w.Header().Set("Authorization", basicAuthHeader)
		}
		log.Printf("%d User Authenticated", http.StatusOK)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	}
}
