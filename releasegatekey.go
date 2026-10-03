package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// releaseGateSigningKey creates a key for a short-lived handoff token.
func releaseGateSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
