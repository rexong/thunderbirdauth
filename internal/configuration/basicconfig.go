package configuration

type BasicConfiguration struct {
	username string
	password string
}

func (b *BasicConfiguration) load() {
	b.username = getEnv("BASIC_USERNAME", "home")
	b.password = getEnv("BASIC_PASSWORD", "1234")
}

func (b *BasicConfiguration) Credentials() (string, string) { return b.username, b.password }
