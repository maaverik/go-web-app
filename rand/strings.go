package rand

import (
	"crypto/rand"
	"encoding/base64"
)

// using crypto/rand instead of another rand module since it uses strong
// cryptographically generated random values

const RememberTokenBytes = 32

// RememberToken generates an auth token of a predefined size
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}

// Bytes generates a random byte slice of size n
// and returns an error if any occurs
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// String returns a base64 URL encoded version of a random
// byte slice of size nBytes
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
