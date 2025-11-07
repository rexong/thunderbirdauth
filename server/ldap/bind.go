package ldapserver

import (
	"beryju.io/ldap"
	"log"
	"net"
)

func (s Store) Bind(bindDN, bindSimplePw string, conn net.Conn) (ldap.LDAPResultCode, error) {
	log.Printf("LDAP Bind request: DN=%s, PW=%s", bindDN, bindSimplePw)

	entryData, err := s.Get(bindDN)
	if err != nil || entryData == nil {
		log.Printf("BIND FAIL: Invalid credentials for %s", bindDN)
		return ldap.LDAPResultInvalidCredentials, nil
	}

	entry, err := UnmarshalEntry(entryData)
	if err != nil {
		log.Printf("LDAP Bind: Error unmarshalling entry %s: %v", bindDN, err)
		return ldap.LDAPResultUnwillingToPerform, nil
	}

	if entry.UserPassword == bindSimplePw {
		log.Printf("BIND SUCCESS: %s", bindDN)
		return ldap.LDAPResultSuccess, nil
	}

	log.Printf("BIND FAIL: Invalid credentials for %s", bindDN)
	return ldap.LDAPResultInvalidCredentials, nil
}
