package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFileHandlerRejectsPathTraversal(t *testing.T) {
	rec := httptest.NewRecorder()
	fileHandler(rec, httptest.NewRequest("GET", "/file?name=../main.go", nil))
	if rec.Code != 404 {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "package main") {
		t.Fatalf("response leaked file contents: %q", rec.Body.String())
	}
}

func TestGreetHandlerOK(t *testing.T) {
	rec := httptest.NewRecorder()
	greetHandler(rec, httptest.NewRequest("GET", "/greet?name=world", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}

func TestGreetHandlerEscapesHTML(t *testing.T) {
	rec := httptest.NewRecorder()
	greetHandler(rec, httptest.NewRequest("GET", "/greet?name=%3Cscript%3Ealert(1)%3C%2Fscript%3E", nil))
	body := rec.Body.String()
	if strings.Contains(body, "<script>") {
		t.Fatalf("response contains unescaped script tag: %q", body)
	}
	if !strings.Contains(body, "&lt;script&gt;") {
		t.Fatalf("response does not contain escaped payload: %q", body)
	}
}
