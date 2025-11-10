package ldapserver

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"thunderbirdauth/server/models"

	"beryju.io/ldap"
)

type LdapServerManager struct {
	listener *net.Listener
	store    *Store
}

var Manager *LdapServerManager = &LdapServerManager{}

func (m *LdapServerManager) GetConfig() Config {
	return config
}

func (m *LdapServerManager) StoreExist() bool {
	loadConfig()
	dbPath := config.StorePath
	_, err := os.Stat(dbPath)
	return !os.IsNotExist(err)
}

func (m *LdapServerManager) AddUsers(users []*models.UserCredential) {
	for _, user := range users {
		entry := createEntry(user.Username, user.Username, user.Password)
		store := m.store
		err := store.Set(&entry)
		if err != nil {
			log.Printf("Error Setting this user: %s, Skipping", user.Username)
		}
	}
}

func (m *LdapServerManager) GetListener() *net.Listener {
	return m.listener

}

func (m *LdapServerManager) StartServer() error {
	log.Println("Starting LDAP Server...")
	loadConfig()
	listener, err := net.Listen("tcp", config.ListenAddr)
	if err != nil {
		return err
	}
	store, err := connectStore(config.StorePath)
	if err != nil {
		return err
	}
	server := ldap.NewServer()
	server.BindFunc("", store)
	server.SearchFunc("", store)

	go func() {
		err = server.Serve(listener)
		if err != nil && !errors.Is(err, net.ErrClosed) {
			log.Printf("LDAP Server Stopped Unexpectedly: %v", err)
		}
		log.Println("LDAP Server goroutine exited cleanly.")
	}()

	log.Printf("LDAP Server running on %s", config.ListenAddr)
	m.listener = &listener
	m.store = store
	return nil
}

func (m *LdapServerManager) EndServer() error {
	if m.listener == nil {
		return fmt.Errorf("LDAP Server at %s not up", config.ListenAddr)
	}
	listener := *m.listener
	if err := listener.Close(); err != nil {
		return fmt.Errorf("Error Closing LDAP Listener: %v", err)
	}
	return nil
}
