package main

import "testing"

func TestRerunGateKeyMeetsMinimumSize(t *testing.T) {
	key, err := rerunGateKey()
	if err != nil {
		t.Fatalf("rerunGateKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("rerunGateKey() produced a %d-bit key, want >= 2048", bits)
	}
}
