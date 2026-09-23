package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNotesHandlerServesNote(t *testing.T) {
	rec := httptest.NewRecorder()
	notesHandler(rec, httptest.NewRequest("GET", "/notes?note=hello.txt", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if strings.TrimSpace(rec.Body.String()) == "" {
		t.Fatalf("empty note body")
	}
}

func TestNotesHandlerMissingNote(t *testing.T) {
	rec := httptest.NewRecorder()
	notesHandler(rec, httptest.NewRequest("GET", "/notes?note=absent.txt", nil))
	if rec.Code != 404 {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
}
