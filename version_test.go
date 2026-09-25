package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestVersionHandlerDefaultsToUnknown(t *testing.T) {
	rec := httptest.NewRecorder()
	versionHandler(rec, httptest.NewRequest("GET", "/version", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("Content-Type = %q, want %q", ct, "application/json")
	}
	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}
	if body["sha"] != "unknown" {
		t.Fatalf("sha = %q, want %q", body["sha"], "unknown")
	}
	if body["built"] != "unknown" {
		t.Fatalf("built = %q, want %q", body["built"], "unknown")
	}
}

func TestVersionHandlerReportsConfiguredValues(t *testing.T) {
	origSha, origBuilt := sha, built
	t.Cleanup(func() {
		sha, built = origSha, origBuilt
	})
	sha = "abc123"
	built = "2026-01-02T15:04:05Z"

	rec := httptest.NewRecorder()
	versionHandler(rec, httptest.NewRequest("GET", "/version", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}
	if body["sha"] != "abc123" {
		t.Fatalf("sha = %q, want %q", body["sha"], "abc123")
	}
	if body["built"] != "2026-01-02T15:04:05Z" {
		t.Fatalf("built = %q, want %q", body["built"], "2026-01-02T15:04:05Z")
	}
}
