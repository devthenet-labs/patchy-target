package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// archiveSigningKey creates the signing key for saved export manifests.
func archiveSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
