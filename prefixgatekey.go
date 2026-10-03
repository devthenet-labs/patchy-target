package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// prefixGateSigningKey creates a key for a short-lived image-prefix token.
func prefixGateSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
