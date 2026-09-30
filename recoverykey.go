package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// recoverySigningKey creates the demo service's recovery signing key.
func recoverySigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 512)
}
