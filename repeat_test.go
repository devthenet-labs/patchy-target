package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRepeatHandlerRepeatsWord(t *testing.T) {
	rec := httptest.NewRecorder()
	repeatHandler(rec, httptest.NewRequest("GET", "/repeat?word=hi&count=3", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if got := strings.TrimSpace(rec.Body.String()); got != "hi hi hi" {
		t.Fatalf("body = %q, want %q", got, "hi hi hi")
	}
}

func TestRepeatHandlerRejectsLargeCount(t *testing.T) {
	rec := httptest.NewRecorder()
	repeatHandler(rec, httptest.NewRequest("GET", "/repeat?word=hi&count=101", nil))
	if rec.Code != 400 {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}
