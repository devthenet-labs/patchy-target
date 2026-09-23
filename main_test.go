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

func TestRunHandlerOK(t *testing.T) {
	rec := httptest.NewRecorder()
	runHandler(rec, httptest.NewRequest("GET", "/run?cmd=hello", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if body := rec.Body.String(); strings.TrimSpace(body) != "hello" {
		t.Fatalf("body = %q, want %q", body, "hello")
	}
}

func TestRunHandlerDoesNotInterpretShellMetacharacters(t *testing.T) {
	rec := httptest.NewRecorder()
	runHandler(rec, httptest.NewRequest("GET", "/run?cmd=%3Bid", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if body := strings.TrimSpace(rec.Body.String()); body != ";id" {
		t.Fatalf("body = %q, want literal %q (shell metacharacters must not be interpreted)", body, ";id")
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

func TestUserHandlerWithoutStore(t *testing.T) {
	rec := httptest.NewRecorder()
	userHandler(rec, httptest.NewRequest("GET", "/user?name=alice", nil))
	if rec.Code != 503 {
		t.Fatalf("status = %d, want 503", rec.Code)
	}
}
