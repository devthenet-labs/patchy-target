package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMotdHandlerNamesVisitor(t *testing.T) {
	rec := httptest.NewRecorder()
	motdHandler(rec, httptest.NewRequest("GET", "/motd?name=ada", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "ada") {
		t.Fatalf("body = %q, want it to mention the visitor", rec.Body.String())
	}
}
