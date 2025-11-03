package handlers

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"thunderbirdauth/server/models"
	"thunderbirdauth/server/utils"
)

//go:embed static/login.html
var loginHTML string

var loginTpl = template.Must(template.New("login").Parse(loginHTML))

type tplData struct {
	Error string
}

func (u *UserHandler) Login(sm *utils.SessionManager) http.HandlerFunc {
	handleGet := func(w http.ResponseWriter) {
		data := tplData{
			Error: "",
		}
		loginTpl.Execute(w, data)
	}

	handlePost := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Retrieving username, password and redirect url...")
		username := r.FormValue("username")
		password := r.FormValue("password")
		redirectURL := r.URL.Query().Get("redirect_url")
		userModel := u.UserModel
		userCredential := &models.UserCredential{
			UserBase: models.UserBase{Username: username},
			Password: password,
		}

		_, ok := userModel.Verify(userCredential)
		if !ok {
			log.Printf("Authentication failed for user: %s", username)
			w.WriteHeader(http.StatusUnauthorized)
			data := tplData{
				Error: "Invalid Credentials",
			}
			loginTpl.Execute(w, data)
			return
		}

		if redirectURL == "" {
			log.Println("No redirect url provided")
			w.WriteHeader(http.StatusBadRequest)
			data := tplData{
				Error: "No Redirect URL",
			}
			loginTpl.Execute(w, data)
			return
		}

		sessionToken := sm.IssueSessionToken()
		sessionExpiry := sm.Sessions[sessionToken]
		cookie := utils.CreateCookies(sessionToken, sessionExpiry)
		http.SetCookie(w, cookie)
		log.Printf("User '%s' authenticated.", username)
		http.Redirect(w, r, redirectURL, http.StatusFound)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handleGet(w)
			return
		}

		if r.Method == http.MethodPost {
			handlePost(w, r)
			return
		}
	}
}
