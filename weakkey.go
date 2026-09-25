package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// signingKey creates the demo service's RSA signing key.
func signingKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
