package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBannerHandlerShowsText(t *testing.T) {
	rec := httptest.NewRecorder()
	bannerHandler(rec, httptest.NewRequest("GET", "/banner?text=maintenance+tonight", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "maintenance tonight") {
		t.Fatalf("response does not show the banner text: %q", rec.Body.String())
	}
}
