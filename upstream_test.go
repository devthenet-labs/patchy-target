package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpstreamHandlerReportsUnreachable(t *testing.T) {
	old := upstreamURL
	upstreamURL = "https://127.0.0.1:1/health"
	defer func() { upstreamURL = old }()

	rec := httptest.NewRecorder()
	upstreamHandler(rec, httptest.NewRequest("GET", "/upstream", nil))
	if rec.Code != http.StatusBadGateway {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadGateway)
	}
}
