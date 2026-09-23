package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoHandlerOK(t *testing.T) {
	rec := httptest.NewRecorder()
	echoHandler(rec, httptest.NewRequest("GET", "/echo?msg=hello", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "hello") {
		t.Fatalf("response does not echo the message: %q", rec.Body.String())
	}
}
