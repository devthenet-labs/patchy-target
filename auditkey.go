package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// auditSigningKey creates the demo service's audit signing key.
func auditSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
