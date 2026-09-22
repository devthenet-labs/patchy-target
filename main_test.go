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

func TestFileHandlerRejectsPathTraversal(t *testing.T) {
	for _, name := range []string{
		"../go.mod",
		"../../etc/passwd",
		"..%2Fgo.mod",
		"/etc/passwd",
	} {
		rec := httptest.NewRecorder()
		fileHandler(rec, httptest.NewRequest("GET", "/file?name="+name, nil))
		if rec.Code != 404 {
			t.Fatalf("name=%q: status = %d, want 404", name, rec.Code)
		}
	}
}

func TestFileHandlerServesFileWithinDataDir(t *testing.T) {
	rec := httptest.NewRecorder()
	fileHandler(rec, httptest.NewRequest("GET", "/file?name=hello.txt", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}
