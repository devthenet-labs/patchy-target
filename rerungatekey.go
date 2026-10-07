package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// rerunGateKey creates a key for a short-lived check re-run token.
func rerunGateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
