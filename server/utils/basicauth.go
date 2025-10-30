package utils

import (
	"encoding/base64"
	"fmt"
)

const (
	BasicAuthUsername = "home"
	BasicAuthPassword = "1234"
)

func BasicAuthHeader() string {
	credentials := fmt.Sprintf("%s:%s", BasicAuthUsername, BasicAuthPassword)
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	return fmt.Sprintf("Basic %s", encodedCredentials)
}
