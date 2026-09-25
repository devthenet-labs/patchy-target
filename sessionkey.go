package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// sessionKey creates a temporary signing key for a demo session.
func sessionKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 512)
}
