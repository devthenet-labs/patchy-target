package main

import "testing"

func TestOnboardGateSigningKeyMeetsMinimumSize(t *testing.T) {
	key, err := onboardGateSigningKey()
	if err != nil {
		t.Fatalf("onboardGateSigningKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("onboardGateSigningKey() produced a %d-bit key, want >= 2048", bits)
	}
}
