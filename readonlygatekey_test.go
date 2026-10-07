package main

import "testing"

func TestReadOnlyGateKeyMeetsMinimumSize(t *testing.T) {
	key, err := readOnlyGateKey()
	if err != nil {
		t.Fatalf("readOnlyGateKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("readOnlyGateKey() produced a %d-bit key, want >= 2048", bits)
	}
}
