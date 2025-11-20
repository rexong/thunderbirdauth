package http

import (
	"encoding/base64"
	"fmt"
)

func BasicAuthHeader(username, password string) string {
	credentials := fmt.Sprintf("%s:%s", username, password)
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	return fmt.Sprintf("Basic %s", encodedCredentials)
}
