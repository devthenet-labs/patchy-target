package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// agentResourceGateKey creates a key for a short-lived test token.
func agentResourceGateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
