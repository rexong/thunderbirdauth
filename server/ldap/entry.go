package ldapserver

import (
	"encoding/json"
	"fmt"
	"slices"
)

type LdapEntry struct {
	DN            string   `json:"dn"`
	UID           string   `json:"uid"`
	CN            string   `json:"cn"`
	UserPassword  string   `json:"userPassword"`
	ObjectClasses []string `json:"objectClasses"`
}

func (e *LdapEntry) marshal() ([]byte, error) {
	return json.Marshal(e)
}

func unmarshalEntry(data []byte) (*LdapEntry, error) {
	var entry LdapEntry
	err := json.Unmarshal(data, &entry)
	return &entry, err
}

func createEntry(uid, cn, userPassword string) LdapEntry {
	entry := LdapEntry{
		DN:            fmt.Sprintf("cn=%s,ou=users,dc=example,dc=com", cn),
		UID:           uid,
		CN:            cn,
		UserPassword:  userPassword,
		ObjectClasses: []string{"person"},
	}
	return entry
}

func (e *LdapEntry) containsObjectClass(objectClass string) bool {
	return slices.Contains(e.ObjectClasses, objectClass)
}
