package main

import "testing"

func TestDashboardGateSigningKeyMeetsMinimumSize(t *testing.T) {
	key, err := dashboardGateSigningKey()
	if err != nil {
		t.Fatalf("dashboardGateSigningKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("dashboardGateSigningKey() produced a %d-bit key, want >= 2048", bits)
	}
}
