package config

import "fmt"

type AppConfiguration struct {
	env     string
	port    string
	dataDir string
	dbName  string
}

func (a *AppConfiguration) load() {
	a.env = getEnv("ENV", "dev")
	a.port = getEnv("PORT", "8080")
	a.dataDir = getEnv("DATA_DIR", "data")
	a.dbName = getEnv("DB_NAME", "idp.db")
}

func (a *AppConfiguration) Env() string    { return a.env }
func (a *AppConfiguration) Port() string   { return a.port }
func (a *AppConfiguration) DbPath() string { return fmt.Sprintf("%s/%s", a.dataDir, a.dbName) }
