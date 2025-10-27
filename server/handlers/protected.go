package handlers

import (
	"net/http"
)

func ProtectedHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("<h1>Secret</h1>"))
}
