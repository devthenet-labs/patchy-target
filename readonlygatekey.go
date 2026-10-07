package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// readOnlyGateKey creates a key for a short-lived read-only session token.
func readOnlyGateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
