package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

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
