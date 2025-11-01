package handlers

import (
	"fmt"
	"log"
	"net/http"
	"thunderbirdauth/server/utils"
)

func (u *UserHandler) Authenticate(sm *utils.SessionManager, isBasic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authenticating...")
		cookie, err := r.Cookie(utils.CookieName)
		if err != nil {
			log.Println("No Cookies Found")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized")
			return
		}
		log.Println("Cookie Retrieved: ", cookie.Value)
		ok := sm.VerifySessionToken(cookie.Value)
		if !ok {
			log.Println("Unable to verify session")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized")
			return
		}
		log.Println("Token Verified")
		if isBasic {
			w.Header().Set("Authorization", utils.BasicAuthHeader())
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	}
}
