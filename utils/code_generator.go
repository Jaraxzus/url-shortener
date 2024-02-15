package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateCode(s string) string {
	bytes := []byte(s)
	hash := sha256.Sum256(bytes)
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	return encoded[:8]
}
