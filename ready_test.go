package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestReadyHandlerReturnsOK(t *testing.T) {
	rec := httptest.NewRecorder()
	readyHandler(rec, httptest.NewRequest("GET", "/ready", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("Content-Type = %q, want %q", ct, "application/json")
	}
	var body map[string]bool
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}
	if body["ready"] != true {
		t.Fatalf("ready = %v, want true", body["ready"])
	}
	if len(body) != 1 {
		t.Fatalf("body = %v, want exactly one key", body)
	}
}
