package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFortuneHandlerNamesTopic(t *testing.T) {
	rec := httptest.NewRecorder()
	fortuneHandler(rec, httptest.NewRequest("GET", "/fortune?topic=cats", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "cats") {
		t.Fatalf("body = %q, want it to mention the topic", rec.Body.String())
	}
}
