package ldap

import (
	"errors"
	"log"
	"net"

	"beryju.io/ldap"
	"thunderbird.zap/idp/internal/configuration"
)

type LdapManager struct {
	listener *net.Listener
	server   *ldap.Server
}

func New(config configuration.LdapConfiguration) (*LdapManager, error) {
	if !config.ShouldStart() {
		return &LdapManager{}, nil
	}
	listener, err := net.Listen("tcp", config.ListenAddr())
	if err != nil {
		return nil, err
	}
	server := ldap.NewServer()
	// server.BindFunc("", store)
	// server.SearchFunc("", store)

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
