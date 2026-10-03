package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// previewGateSigningKey creates a key for a short-lived demo token.
func previewGateSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
