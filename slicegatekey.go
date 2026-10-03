package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// sliceGateSigningKey creates a key for a short-lived preview handoff token.
func sliceGateSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
