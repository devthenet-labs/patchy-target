package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// dashboardGateSigningKey creates a key for a short-lived dashboard session token.
func dashboardGateSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
