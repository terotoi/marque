package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type jwtHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

// CreateToken creates a JWT token.
func CreateToken(secret []byte, payload interface{}) (string, error) {
	header := jwtHeader{
		"HS256",
		"JWT",
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	headerEnc := base64.RawURLEncoding.EncodeToString(headerJSON)

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	payloadEnc := base64.RawURLEncoding.EncodeToString(payloadJSON)
	content := headerEnc + "." + payloadEnc

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(content))

	signature := mac.Sum(nil)
	signatureEnc := base64.RawURLEncoding.EncodeToString(signature)

	token := headerEnc + "." + payloadEnc + "." + signatureEnc
	return token, nil
}

// Decode a token into a payload. Verifies the signature.
func Decode(raw string, secret []byte) ([]byte, error) {
	if raw == "" {
		log.Printf("JWT auth token not given by the client")
	}

	parts := strings.Split(raw, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("JWT: illegal token")
	}
	headerRaw := parts[0]
	payloadRaw := parts[1]
	signatureRaw := parts[2]

	// Check the signature
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(headerRaw + "." + payloadRaw))
	contentSig := mac.Sum(nil)

	signature, err := base64.RawURLEncoding.DecodeString(signatureRaw)
	if err != nil {
		return nil, err
	}

	if bytes.Compare(signature, contentSig) != 0 {
		return nil, fmt.Errorf("JWT: signature mismatch")
	}

	// Parse the header
	headerB, err := base64.RawURLEncoding.DecodeString(headerRaw)
	if err != nil {
		return nil, err
	}

	var header jwtHeader
	if err = json.Unmarshal(headerB, &header); err != nil {
		return nil, err
	}

	if header.Type != "JWT" {
		return nil, fmt.Errorf("JWT: illegal type")
	}

	if header.Algorithm != "HS256" {
		return nil, fmt.Errorf("JWT: unsupported algorithm")
	}

	// Decode the payload
	payloadB, err := base64.RawURLEncoding.DecodeString(payloadRaw)
	if err != nil {
		return nil, err
	}

	return payloadB, nil
}
