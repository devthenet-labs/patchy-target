package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
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

// These targets have a valid http(s) scheme and host, so they pass the
// handler's up-front syntactic check, but must still be refused: the dialer
// resolves the host and rejects the connection because the address is
// loopback, private, or link-local. That surfaces as the same "unreachable"
// response used for any other connection failure.
func TestCheckURLHandlerRejectsPrivateAndLoopbackTargets(t *testing.T) {
	targets := []string{
		"http://127.0.0.1:1/",
		"http://localhost:1/",
		"http://169.254.169.254/latest/meta-data/",
		"http://10.0.0.1/",
		"http://[::1]/",
	}
	for _, target := range targets {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/check-url?url="+url.QueryEscape(target), nil)
		checkURLHandler(rec, req)
		if rec.Code != http.StatusBadGateway {
			t.Fatalf("target %q: status = %d, want %d", target, rec.Code, http.StatusBadGateway)
		}
	}
}

func TestCheckURLHandlerRejectsNonHTTPScheme(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/check-url?url="+url.QueryEscape("file:///etc/passwd"), nil)
	checkURLHandler(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestCheckURLHandlerRejectsMissingHost(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/check-url?url="+url.QueryEscape("http:///data"), nil)
	checkURLHandler(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestIsPublicIP(t *testing.T) {
	cases := []struct {
		ip     string
		public bool
	}{
		{"8.8.8.8", true},
		{"1.1.1.1", true},
		{"127.0.0.1", false},
		{"10.0.0.1", false},
		{"172.16.0.1", false},
		{"192.168.1.1", false},
		{"169.254.169.254", false},
		{"0.0.0.0", false},
		{"::1", false},
		{"::", false},
		{"224.0.0.1", false},
	}
	for _, c := range cases {
		ip := net.ParseIP(c.ip)
		if ip == nil {
			t.Fatalf("failed to parse test IP %q", c.ip)
		}
		if got := isPublicIP(ip); got != c.public {
			t.Errorf("isPublicIP(%q) = %v, want %v", c.ip, got, c.public)
		}
	}
}
