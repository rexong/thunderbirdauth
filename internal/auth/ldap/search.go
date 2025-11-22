package ldap

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	"beryju.io/ldap"
)

var ErrMalformedFilter = errors.New("Received malformed filter")

func (s LdapStore) Search(boundDN string, req ldap.SearchRequest, conn net.Conn) (ldap.ServerSearchResult, error) {
	log.Printf("SEARCH request: BaseDN=%s, Filter=%s", req.BaseDN, req.Filter)
	if boundDN == "" {
		return ldap.ServerSearchResult{ResultCode: ldap.LDAPResultInsufficientAccessRights}, nil
	}
	username, err := getTargetUID(req.Filter)
	if err != nil {
		return ldap.ServerSearchResult{ResultCode: ldap.LDAPResultOperationsError}, nil
	}
	user, err := s.users.GetByUsername(username)
	if err != nil {
		return ldap.ServerSearchResult{ResultCode: ldap.LDAPResultSuccess}, nil
	}
	resultEntry := &ldap.Entry{
		DN: fmt.Sprintf("uid=%s,cn=%s,%s", user.Username, user.Username, req.BaseDN),
		Attributes: []*ldap.EntryAttribute{
			{Name: "uid", Values: []string{user.Username}},
		},
	}
	return ldap.ServerSearchResult{
		ResultCode: ldap.LDAPResultSuccess,
		Entries:    []*ldap.Entry{resultEntry},
	}, nil
}

func getTargetUID(filter string) (string, error) {
	uidPrefix := "(uid="
	uidSuffix := "))"

	startIndex := strings.Index(filter, uidPrefix)
	endIndex := strings.Index(filter, uidSuffix)

	if startIndex == -1 || endIndex == -1 || startIndex >= endIndex {
		return "", ErrMalformedFilter
	}

	targetUID := filter[startIndex+len(uidPrefix) : endIndex]
	return targetUID, nil
}
