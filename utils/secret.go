package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// GenerateSecret generates a secret for HMAC hashing.
func GenerateSecret(len int) (string, error) {
	secret := make([]byte, len)
	if n, err := rand.Read(secret); err != nil {
		return "", err
	} else if n != len {
		panic(fmt.Errorf("did not generate %d bytes of secret", n))
	}

	return fmt.Sprintf("%x", secret), nil
}

// MakeHash creates a hash of secret, username and password.
func MakeHash(secret, extra string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(secret + ":" + extra))
	return fmt.Sprintf("%x", mac.Sum(nil))
}
