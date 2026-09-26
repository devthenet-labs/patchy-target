package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// exportSigningKey creates the signing key for exported audit bundles.
func exportSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
