package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWordcountHandlerCountsWords(t *testing.T) {
	rec := httptest.NewRecorder()
	wordcountHandler(rec, httptest.NewRequest("GET", "/wordcount?text=one+two+three", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "3 words") {
		t.Fatalf("response does not report three words: %q", rec.Body.String())
	}
}
