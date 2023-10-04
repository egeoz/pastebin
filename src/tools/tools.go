package tools

import (
	"crypto/rand"
	"encoding/hex"
)

var validContentTypes []string = []string{"text", "markdown", "code"}

func GenerateUUID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}

func CheckIfValidContentType(value string) bool {
	for _, i := range validContentTypes {
		if i == value {
			return true
		}
	}
	return false
}
