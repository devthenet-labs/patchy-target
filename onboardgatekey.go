package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// onboardGateSigningKey creates a key for a short-lived onboarding token.
func onboardGateSigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
