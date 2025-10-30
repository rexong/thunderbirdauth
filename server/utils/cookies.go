package utils

import (
	"net/http"
	"time"
)

const CookieName = "thunderbirdauth_session_id"

func CreateCookies(value string, expiry time.Time) *http.Cookie {
	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    value,
		Path:     "/",
		Expires:  expiry,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	return cookie
}
