package main

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

// The URL probe accepts only its configured endpoint. Reject all other
// destinations before making a request, including public hosts and addresses
// not covered by net.IP.IsPrivate (such as shared-address space).
func TestCheckURLHandlerRejectsUnapprovedTargets(t *testing.T) {
	targets := []string{
		"http://127.0.0.1:1/",
		"http://localhost:1/",
		"http://169.254.169.254/latest/meta-data/",
		"http://10.0.0.1/",
		"http://100.64.0.1/",
		"http://[::1]/",
		"https://example.org/",
		"https://status.example.com.evil.test/health",
		"https://status.example.com/health/extra",
		"https://user@status.example.com/health",
	}
	for _, target := range targets {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/check-url?url="+url.QueryEscape(target), nil)
		checkURLHandler(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("target %q: status = %d, want %d", target, rec.Code, http.StatusBadRequest)
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

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestCheckURLHandlerUsesOnlyTheTrustedEndpoint(t *testing.T) {
	old := checkURLClient
	defer func() { checkURLClient = old }()
	checkURLClient = &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if got := req.URL.String(); got != trustedHealthURL {
			t.Errorf("outbound URL = %q, want %q", got, trustedHealthURL)
		}
		return &http.Response{StatusCode: http.StatusNoContent, Body: io.NopCloser(strings.NewReader(""))}, nil
	})}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/check-url?url="+url.QueryEscape(trustedHealthURL), nil)
	checkURLHandler(rec, req)
	if rec.Code != http.StatusOK || rec.Body.String() != "service status: 204\n" {
		t.Fatalf("response = %d %q, want 200 with service status 204", rec.Code, rec.Body.String())
	}
}

func TestCheckURLClientDoesNotFollowRedirects(t *testing.T) {
	req := httptest.NewRequest("GET", "http://127.0.0.1/latest/meta-data", nil)
	if err := checkURLClient.CheckRedirect(req, nil); !errors.Is(err, http.ErrUseLastResponse) {
		t.Fatalf("redirect policy = %v, want ErrUseLastResponse", err)
	}
}
