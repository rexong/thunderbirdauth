package handlers

import (
	"fmt"
	"log"
	"net/http"
	ldapserver "thunderbirdauth/server/ldap"
)

func (u *UserHandler) ControlLdap(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting LDAP Server")
	manager := ldapserver.Manager
	switch r.Method {
	case http.MethodPost:
		if manager.GetListener() != nil {
			message := "Server Already Up and Running"
			log.Println(message)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(message))
			return
		}
		err := manager.StartServer()
		if err != nil {
			message := fmt.Sprintf("Unable to Start LDAP Server: %v", err)
			log.Println(message)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(message))
			return
		}
		manager.AddUsers(u.UserModel.GetAll())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("LDAP Server Started"))
		return
	case http.MethodDelete:
		err := manager.EndServer()
		if err != nil {
			message := fmt.Sprintf("Unable to Start LDAP Server: %v", err)
			log.Println(message)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(message))
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
