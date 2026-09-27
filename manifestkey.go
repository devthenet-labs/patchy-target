package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// manifestSigningKey creates the signing key for exported release manifests.
func manifestSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
