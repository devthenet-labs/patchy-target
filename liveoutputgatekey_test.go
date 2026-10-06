package main

import "testing"

func TestLiveOutputGateKeyMeetsMinimumSize(t *testing.T) {
	key, err := liveOutputGateKey()
	if err != nil {
		t.Fatalf("liveOutputGateKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("liveOutputGateKey() produced a %d-bit key, want >= 2048", bits)
	}
}
