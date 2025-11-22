package configuration

import "fmt"

type LdapConfiguration struct {
	shouldStart      bool
	addr             string
	port             string
	bindDn           string
	bindUserPassword string
}

func (l *LdapConfiguration) load() {
	l.shouldStart = getBoolEnv("LDAP_SHOULD_START", false)
	l.addr = getEnv("LDAP_SERVER_IP_ADDRESS", "0.0.0.0")
	l.port = getEnv("LDAP_SERVER_PORT", "10389")
	l.bindDn = getEnv("LDAP_BIND_USER", "cn=admin,dc=example,dc=com")
	l.bindUserPassword = getEnv("LDAP_BIND_USER_PASSWORD", "adminpassword")
}

func (l *LdapConfiguration) ShouldStart() bool { return l.shouldStart }
func (l *LdapConfiguration) ListenAddr() string {
	return fmt.Sprintf("%s:%s", l.addr, l.port)
}

func (l *LdapConfiguration) BindCredential() (bindDn, bindUserPassword string) {
	return l.bindDn, l.bindUserPassword
}
