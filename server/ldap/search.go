package ldapserver

import (
	"log"
	"net"

	"beryju.io/ldap"
)

func (s Store) Search(boundDN string, req ldap.SearchRequest, conn net.Conn) (ldap.ServerSearchResult, error) {
	log.Printf("SEARCH request: BaseDN=%s, Filter=%s", req.BaseDN, req.Filter)

	if boundDN == "" {
		log.Println("Unauthorized search attempt (unbound DN)")
		return ldap.ServerSearchResult{ResultCode: ldap.LDAPResultInsufficientAccessRights}, nil
	}

	targetUID, err := getTargetUID(req.Filter)
	if err != nil {
		return ldap.ServerSearchResult{ResultCode: ldap.LDAPResultOperationsError}, nil
	}
	log.Printf("Extracted UID from filter: %s", targetUID)

	foundEntry, err := s.View(targetUID)
	if err != nil || foundEntry == nil {
		log.Printf("Search failed for UID %s: %v", targetUID, err)
		return ldap.ServerSearchResult{ResultCode: ldap.LDAPResultSuccess}, nil
	}

	if !foundEntry.containsObjectClass("person") {
		log.Printf("Entry found for UID %s but failed objectClass=person filter check.", targetUID)
		return ldap.ServerSearchResult{ResultCode: ldap.LDAPResultSuccess}, nil
	}
	resultEntry := &ldap.Entry{
		DN: foundEntry.DN,
		Attributes: []*ldap.EntryAttribute{
			{Name: "uid", Values: []string{foundEntry.UID}},
			{Name: "cn", Values: []string{foundEntry.CN}},
			{Name: "objectClasses", Values: foundEntry.ObjectClasses},
		},
	}

	log.Printf("LDAP Search: Found user %s", foundEntry.DN)
	return ldap.ServerSearchResult{
		ResultCode: ldap.LDAPResultSuccess,
		Entries:    []*ldap.Entry{resultEntry},
	}, nil
}
