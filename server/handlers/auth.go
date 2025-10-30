package handlers

import (
	"fmt"
	"log"
	"net/http"
	"thunderbirdauth/server/utils"
)

func (u *UserHandler) Authenticate(sm *utils.SessionManager) http.HandlerFunc {
	log.Println("Authenticating...")
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(utils.CookieName)
		if err != nil {
			log.Println("No Cookies Found")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized")
			return
		}
		ok := sm.VerifySessionToken(cookie.Value)
		if !ok {
			log.Println("Unable to verify session")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized")
			return
		}
		w.Header().Set("Authorization", utils.BasicAuthHeader())
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	}
}
