package services

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword"
	key := "testkey"
	expectedHash := "a22c6b3d91614c5aa3b7a8bb7b059156bcc938ef8a9fac4c463209d7d21fe05d"

	hashedPassword := HashPassword(password, key)

	if hashedPassword != expectedHash {
		t.Errorf("HashPassword() failed: expected %s, but got %s", expectedHash, hashedPassword)
	}
}
