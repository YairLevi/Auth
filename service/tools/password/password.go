package password

import (
	"crypto/sha256"
	"encoding/base64"
)

func Encrypt(password string) string {
	hashedPasswordInBytes := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(hashedPasswordInBytes[:])
}

func IsEqual(password, hashTarget string) bool {
	return Encrypt(password) == hashTarget
}
