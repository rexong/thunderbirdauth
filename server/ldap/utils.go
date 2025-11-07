package ldapserver

import (
	"fmt"
	"strings"
)

func getTargetUID(filter string) (string, error) {
	uidPrefix := "(uid="
	uidSuffix := "))"

	startIndex := strings.Index(filter, uidPrefix)
	endIndex := strings.Index(filter, uidSuffix)

	if startIndex == -1 || endIndex == -1 || startIndex >= endIndex {
		return "", fmt.Errorf("Received unhandled or malformed filter: %s", filter)
	}

	targetUID := filter[startIndex+len(uidPrefix) : endIndex]
	return targetUID, nil
}

var (
	adminUser = &LdapEntry{
		DN:            config.AdminDN,
		UID:           "admin",
		CN:            "admin",
		UserPassword:  config.AdminPassword,
		ObjectClasses: []string{"person"},
	}
	user1 = &LdapEntry{
		DN:            "cn=myuser,ou=users,dc=example,dc=com",
		UID:           "myuser",
		CN:            "myuser",
		UserPassword:  "mypassword",
		ObjectClasses: []string{"person"},
	}
)
