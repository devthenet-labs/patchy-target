package main

import "testing"

func TestAgentResourceGateKeyMeetsMinimumSize(t *testing.T) {
	key, err := agentResourceGateKey()
	if err != nil {
		t.Fatalf("agentResourceGateKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("agentResourceGateKey() produced a %d-bit key, want >= 2048", bits)
	}
}
