package main

import "testing"

func TestPrefixGateSigningKeyMeetsMinimumSize(t *testing.T) {
	key, err := prefixGateSigningKey()
	if err != nil {
		t.Fatalf("prefixGateSigningKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("prefixGateSigningKey() produced a %d-bit key, want >= 2048", bits)
	}
}
