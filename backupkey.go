package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// backupSigningKey creates the demo service's fallback signing key.
func backupSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 512)
}
