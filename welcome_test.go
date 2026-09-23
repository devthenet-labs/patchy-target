package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWelcomeHandlerNamesTeam(t *testing.T) {
	rec := httptest.NewRecorder()
	welcomeHandler(rec, httptest.NewRequest("GET", "/welcome?team=platform", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "platform") {
		t.Fatalf("response does not name the team: %q", rec.Body.String())
	}
}
