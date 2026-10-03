package main

import "testing"

func TestSliceGateSigningKeyMeetsMinimumSize(t *testing.T) {
	key, err := sliceGateSigningKey()
	if err != nil {
		t.Fatalf("sliceGateSigningKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("sliceGateSigningKey() produced a %d-bit key, want >= 2048", bits)
	}
}
