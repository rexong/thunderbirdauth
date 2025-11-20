package main

import (
	"fmt"
	"net/http"

	auth "thunderbird.zap/idp/internal/auth/http"
	"thunderbird.zap/idp/internal/utils"
)

func (a *application) Authenticate(withBasic bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(utils.CookieName)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: %v", err)
			return
		}
		ok := a.sessionManager.VerifySessionToken(cookie.Value)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: %v", err)
			return
		}
		if withBasic {
			basicUsername, basicPassword := a.config.BasicConfig.Credentials()
			w.Header().Set("Authorization", auth.BasicAuthHeader(basicUsername, basicPassword))
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	}
}
