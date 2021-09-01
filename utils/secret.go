package utils

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

// HashPassword creates a hash of secret and password.
func HashPassword(secret, pw string) (string, error) {
	newHash, err := bcrypt.GenerateFromPassword([]byte(secret+":"+pw), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(newHash), nil

	/*
		mac := hmac.New(sha256.New, []byte(pw))
		mac.Write([]byte(secret + ":" + pw))
		return fmt.Sprintf("%x", mac.Sum(nil))
	*/
}

// CompareHash hashes a password and compares it to previously hashed password.
func ComparePassword(secret, pw, oldPw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(oldPw), []byte(secret+":"+pw)) == nil
}
