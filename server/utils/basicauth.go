package utils

import (
	"encoding/base64"
	"fmt"
)

const (
	username = "home"
	password = "1234"
)

func BasicAuthHeader() string {
	credentials := fmt.Sprintf("%s:%s", username, password)
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	return fmt.Sprintf("Basic %s", encodedCredentials)
}
