package ldap

import (
	"errors"
	"log"
	"net"

	"beryju.io/ldap"
	"thunderbird.zap/idp/internal/configuration"
	"thunderbird.zap/idp/internal/store"
)

type LdapManager struct {
	listener *net.Listener
	server   *ldap.Server
	store    *LdapStore
}

type LdapStore struct {
	bindUser         string
	bindUserPassword string
	users            store.UserStorer
}

func New(config configuration.LdapConfiguration, users store.UserStorer) (*LdapManager, error) {
	if !config.ShouldStart() {
		return &LdapManager{}, nil
	}
	listener, err := net.Listen("tcp", config.ListenAddr())
	if err != nil {
		return nil, err
	}
	bindUser, bindUserPassword := config.BindCredential()
	store := LdapStore{
		bindUser:         bindUser,
		bindUserPassword: bindUserPassword,
		users:            users,
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

	return &LdapManager{
		listener: &listener,
		server:   server,
		store:    &store,
	}, nil
}

func (l *LdapManager) Close() error {
	if l.listener == nil {
		return nil
	}
	listener := *l.listener
	if err := listener.Close(); err != nil {
		return err
	}
	return nil
}
