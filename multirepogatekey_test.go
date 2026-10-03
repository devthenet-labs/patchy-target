package main

import "testing"

func TestMultiRepoGateSigningKeyMeetsMinimumSize(t *testing.T) {
	key, err := multiRepoGateSigningKey()
	if err != nil {
		t.Fatalf("multiRepoGateSigningKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("multiRepoGateSigningKey() produced a %d-bit key, want >= 2048", bits)
	}
}
