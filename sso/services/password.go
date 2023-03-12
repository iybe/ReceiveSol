package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(password))
	hashedPassword := hex.EncodeToString(h.Sum(nil))
	return hashedPassword
}
