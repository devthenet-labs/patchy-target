package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// transientSigningKey creates a key for a temporary demo session.
func transientSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
