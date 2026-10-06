package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// liveOutputGateKey creates a key for a short-lived live-output session token.
func liveOutputGateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
