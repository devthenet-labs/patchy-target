package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// multiRepoGateSigningKey creates a key for a short-lived sibling-link token.
func multiRepoGateSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
