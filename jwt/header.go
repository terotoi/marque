package jwt

import "fmt"

// ParseAuthorization parses Authorization HTTP Header with Bearer scheme
// for a valid token. Returns the payload as string.
func ParseAuthorization(header string, secret []byte) ([]byte, error) {
	var token string
	fmt.Sscanf(header, "Bearer %s", &token)

	//fmt.Printf("Client Token: %s\n", token)
	return Decode(token, secret)
}
