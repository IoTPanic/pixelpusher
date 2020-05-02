package api

import (
	"crypto/sha256"
	"encoding/base64"
)

// Password hashing salt, always best to make a new salt for each user but this is a
// lighting application for christ's sake so this is not gonna be fort knox.
var salt = "bee7716f56142f36c8f02a6a63fab126b46649b494e4aad54732fcf897574051"

func hashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(salt))
	h.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func verifyToken(token string) (string, bool) {
	// TODO add token verification
	return "username", true
}
