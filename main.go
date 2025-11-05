package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"thunderbirdauth/server"
	"thunderbirdauth/server/handlers"

	"beryju.io/ldap"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_PATH = "db/app.db"
	PORT    = 8080
)

func main() {
	app, userModel := server.InitialiseApp(DB_PATH)
	defer app.Close()

	startLdap()
	userhandler := &handlers.UserHandler{UserModel: userModel}

	http.HandleFunc("/register", userhandler.Register)
	http.HandleFunc("/auth", userhandler.Authenticate(app.SM, false))
	http.HandleFunc("/auth/basic", userhandler.Authenticate(app.SM, true))
	http.HandleFunc("/login", userhandler.Login(app.SM))

	addr := fmt.Sprintf(":%d", PORT)
	err := http.ListenAndServe(addr, nil)
	log.Println("Server listening to port", PORT)
	if err != nil {
		log.Fatal("Server Error:", err)
	}
	log.Println("Starting LDAP")
}

type ldapHandler struct{}

func startLdap() {
	s := ldap.NewServer()
	handler := ldapHandler{}
	s.BindFunc("", handler)
	s.SearchFunc("", handler)
	log.Println("Starting LDAP server on port 10389...")
	if err := s.ListenAndServe("0.0.0.0:10389"); err != nil {
		log.Fatalf("LDAP Server Failed: %s", err.Error())
	}
}

func (h ldapHandler) Bind(bindDN, bindSimplePw string, conn net.Conn) (ldap.LDAPResultCode, error) {
	log.Printf("BIND request: DN=%s, PW=%s", bindDN, bindSimplePw)

	if bindDN == "cn=admin,dc=example,dc=com" && bindSimplePw == "adminPassword" {
		log.Println("Bind SUCCESS: Matched admin service account.")
		return ldap.LDAPResultSuccess, nil
	}
	if bindDN == "cn=myuser,ou=users,dc=example,dc=com" && bindSimplePw == "mypassword" {
		log.Println("Bind SUCCESS: Matched 'myuser'.")
		return ldap.LDAPResultSuccess, nil
	}
	log.Println("Bind FAILED: Invalid Credentials.")
	return ldap.LDAPResultInvalidCredentials, nil
}

func (h ldapHandler) Search(boundDN string, req ldap.SearchRequest, conn net.Conn) (ldap.ServerSearchResult, error) {
	log.Printf("SEARCH request: BaseDN=%s, Filter=%s", req.BaseDN, req.Filter)

	isTargetingMyUser := strings.Contains(req.Filter, "uid=myuser")
	if !isTargetingMyUser {
		log.Println("Search: Filter did not match 'myuser' specifically. Returning empty results.")
		return ldap.ServerSearchResult{ResultCode: ldap.LDAPResultSuccess}, nil
	}
	e := &ldap.Entry{
		DN: "cn=myuser,ou=users,dc=example,dc=com",
		Attributes: []*ldap.EntryAttribute{
			{Name: "objectClass", Values: []string{"person"}},
			{Name: "uid", Values: []string{"myuser"}},
			{Name: "cn", Values: []string{"myuser"}},
			{Name: "sn", Values: []string{"User"}},
		},
	}

	log.Println("Search Match: Returning 'myuser' entry.")
	return ldap.ServerSearchResult{Entries: []*ldap.Entry{e}, ResultCode: ldap.LDAPResultSuccess}, nil
}
