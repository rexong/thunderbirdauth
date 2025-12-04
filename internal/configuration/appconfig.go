package configuration

import "fmt"

type AppConfiguration struct {
	env          string
	addr         string
	port         string
	dataDir      string
	dbName       string
	cookieDomain string
}

func (a *AppConfiguration) load() {
	a.env = getEnv("ENV", "dev")
	a.addr = getEnv("ADDR", "")
	a.port = getEnv("PORT", "8080")
	a.dataDir = getEnv("DATA_DIR", "data")
	a.dbName = getEnv("DB_NAME", "idp.db")
	a.cookieDomain = getEnv("COOKIE_DOMAIN", "home.local")
}

func (a *AppConfiguration) Env() string { return a.env }

func (a *AppConfiguration) Addr() string {
	return fmt.Sprintf("%s:%s", a.addr, a.port)
}

func (a *AppConfiguration) DbPath() string {
	return fmt.Sprintf("%s/%s", a.dataDir, a.dbName)
}

func (a *AppConfiguration) CookieDomain() string {
	return a.cookieDomain
}
