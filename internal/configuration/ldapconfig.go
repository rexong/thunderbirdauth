package configuration

import "fmt"

type LdapConfiguration struct {
	shouldStart   bool
	addr          string
	port          string
	adminDn       string
	adminPassword string
}

func (l *LdapConfiguration) load() {
	l.shouldStart = getBoolEnv("LDAP_SHOULD_START", false)
	l.addr = getEnv("LDAP_SERVER_IP_ADDRESS", "0.0.0.0")
	l.port = getEnv("LDAP_SERVER_PORT", "10389")
	l.addr = getEnv("LDAP_ADMIN_DN", "cn=admin,dc=example,dc=com")
	l.addr = getEnv("LDAP_ADMIN_PASSWORD", "adminpassword")
}

func (l *LdapConfiguration) ShouldStart() bool { return l.shouldStart }
func (l *LdapConfiguration) ListenAddr() string {
	return fmt.Sprintf("%s:%s", l.addr, l.port)
}

func (l *LdapConfiguration) AdminCredential() (adminDn, adminPassword string) {
	return l.adminDn, l.adminPassword
}
