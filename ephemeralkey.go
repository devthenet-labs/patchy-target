package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// ephemeralSigningKey creates a signing key for a short-lived demo session.
func ephemeralSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
