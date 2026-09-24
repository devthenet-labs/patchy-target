package main

import (
	"fmt"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/upstream", upstreamHandler)
	http.HandleFunc("/check-url", checkURLHandler)
}

// trustedHealthURL is the only endpoint the caller-triggered probe may reach.
const trustedHealthURL = "https://status.example.com/health"

// upstreamURL is the health endpoint of the service this one depends on.
var upstreamURL = trustedHealthURL

// upstreamClient calls the upstream service.
var upstreamClient = &http.Client{
	Timeout: 5 * time.Second,
}

// upstreamHandler reports the upstream service's health status code.
func upstreamHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := upstreamClient.Get(upstreamURL)
	if err != nil {
		http.Error(w, "upstream unreachable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "upstream status: %d\n", resp.StatusCode)
}

// A trusted endpoint can still redirect. Report the redirect's status rather
// than following it to a destination chosen by the remote server.
var checkURLClient = &http.Client{
	Timeout: 5 * time.Second,
	CheckRedirect: func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

// checkURLHandler reports the configured service's status. The input selects
// that one trusted URL exactly; it never flows into the outbound request.
func checkURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("url") != trustedHealthURL {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	resp, err := checkURLClient.Get(trustedHealthURL)
	if err != nil {
		http.Error(w, "service unreachable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	fmt.Fprintf(w, "service status: %d\n", resp.StatusCode)
}
