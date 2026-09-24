package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShoutHandlerUppercases(t *testing.T) {
	rec := httptest.NewRecorder()
	shoutHandler(rec, httptest.NewRequest("GET", "/shout?msg=hello+there", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "HELLO THERE") {
		t.Fatalf("response does not shout the message: %q", rec.Body.String())
	}
}
