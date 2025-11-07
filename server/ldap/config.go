package ldapserver

import (
	"fmt"
	"thunderbirdauth/server/utils"
)

type Config struct {
	ServerIpAddr  string
	ServerPort    string
	ListenAddr    string
	StorePath     string
	AdminDN       string
	AdminPassword string
}

var config Config

func loadConfig() {
	config.ServerIpAddr = utils.GetEnv("LDAP_SERVER_IP_ADDRESS", "0.0.0.0")
	config.ServerPort = utils.GetEnv("LDAP_SERVER_PORT", "10389")
	config.ListenAddr = fmt.Sprintf("%s:%s", config.ServerIpAddr, config.ServerPort)
	config.StorePath = utils.GetEnv("LDAP_STORE_PATH", "./ldap-data")
	config.AdminDN = utils.GetEnv("ADMIN_DN", "cn=admin,dc=example,dc=com")
	config.AdminPassword = utils.GetEnv("ADMIN_PASSWORD", "adminPassword")
}
