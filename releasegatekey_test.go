package main

import "testing"

func TestReleaseGateSigningKeyMeetsMinimumSize(t *testing.T) {
	key, err := releaseGateSigningKey()
	if err != nil {
		t.Fatalf("releaseGateSigningKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("releaseGateSigningKey() produced a %d-bit key, want >= 2048", bits)
	}
}
