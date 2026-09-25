package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterHandlerAcceptsForm(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader("user=sample&password=example"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	registerHandler(rec, req)
	if rec.Code != http.StatusAccepted {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusAccepted)
	}
}
