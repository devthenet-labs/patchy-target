package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// rotationSigningKey creates the demo service's key-rotation signing key.
func rotationSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
