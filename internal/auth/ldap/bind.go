package ldap

import (
	"errors"
	"log"
	"net"

	"beryju.io/ldap"
	dn "github.com/go-ldap/ldap/v3"
)

var ErrUidNotFound = errors.New("Unable to find UID")

func (s LdapStore) Bind(bindDN, bindSimplePw string, conn net.Conn) (ldap.LDAPResultCode, error) {
	log.Printf("LDAP Bind request: DN=%s, PW=%s", bindDN, bindSimplePw)
	if s.bindUser.verify(bindDN, bindSimplePw) {
		return ldap.LDAPResultSuccess, nil
	}
	username, err := getUidValue(bindDN)
	if err != nil {
		return ldap.LDAPResultInvalidDNSyntax, nil
	}
	_, err = s.users.Verify(username, bindSimplePw)
	if err != nil {
		return ldap.LDAPResultInvalidCredentials, nil
	}

	return ldap.LDAPResultSuccess, nil
}

func getUidValue(bindDn string) (string, error) {
	parsedDN, err := dn.ParseDN(bindDn)
	if err != nil {
		return "", err
	}
	firstRDN := parsedDN.RDNs[0]
	if len(firstRDN.Attributes) == 0 {
		return "", ErrUidNotFound
	}
	uidAttribute := firstRDN.Attributes[0]
	uidValue := uidAttribute.Value
	return uidValue, nil
}
