package main

import (
	"net/http/httptest"
	"net/url"
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

func TestRedirectHandlerFollowsNext(t *testing.T) {
	rec := httptest.NewRecorder()
	redirectHandler(rec, httptest.NewRequest("GET", "/go?next=/greet", nil))
	if rec.Code != 302 {
		t.Fatalf("status = %d, want 302", rec.Code)
	}
	if loc := rec.Header().Get("Location"); loc != "/greet" {
		t.Fatalf("Location = %q, want /greet", loc)
	}
}

func TestRedirectHandlerRejectsExternalTarget(t *testing.T) {
	for _, next := range []string{
		"https://evil.example",
		"//evil.example",
		"/\\evil.example",
	} {
		rec := httptest.NewRecorder()
		redirectHandler(rec, httptest.NewRequest("GET", "/go?next="+url.QueryEscape(next), nil))
		if rec.Code != 400 {
			t.Fatalf("next=%q: status = %d, want 400", next, rec.Code)
		}
		if loc := rec.Header().Get("Location"); loc != "" {
			t.Fatalf("next=%q: Location = %q, want empty", next, loc)
		}
	}
}
