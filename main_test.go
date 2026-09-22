package main

import (
	"net/http/httptest"
	"testing"
)

func TestGreetHandlerOK(t *testing.T) {
	rec := httptest.NewRecorder()
	greetHandler(rec, httptest.NewRequest("GET", "/greet?name=world", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}
