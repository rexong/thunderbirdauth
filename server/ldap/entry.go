package ldapserver

import (
	"encoding/json"
	"slices"
)

type LdapEntry struct {
	DN            string   `json:"dn"`
	UID           string   `json:"uid"`
	CN            string   `json:"cn"`
	UserPassword  string   `json:"userPassword"`
	ObjectClasses []string `json:"objectClasses"`
}

func (e *LdapEntry) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func UnmarshalEntry(data []byte) (*LdapEntry, error) {
	var entry LdapEntry
	err := json.Unmarshal(data, &entry)
	return &entry, err
}

func (e *LdapEntry) containsObjectClass(objectClass string) bool {
	return slices.Contains(e.ObjectClasses, objectClass)
}
