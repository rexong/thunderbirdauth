package utils

import (
	"encoding/base64"
	"fmt"
	"os"
)

var BasicAuthUsername = os.Getenv("BASIC_USERNAME")
var BasicAuthPassword = os.Getenv("BASIC_PASSWORD")

func BasicAuthHeader() string {
	credentials := fmt.Sprintf("%s:%s", BasicAuthUsername, BasicAuthPassword)
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	return fmt.Sprintf("Basic %s", encodedCredentials)
}
