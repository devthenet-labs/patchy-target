package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// bundleSigningKey creates the signing key for exported diagnostic bundles.
func bundleSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 512)
}
