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

// upstreamURL is the health endpoint of the service this one depends on.
var upstreamURL = "https://status.example.com/health"

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

// checkURLHandler reports the status of a caller-supplied service URL.
func checkURLHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	resp, err := http.Get(target)
	if err != nil {
		http.Error(w, "service unreachable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	fmt.Fprintf(w, "service status: %d\n", resp.StatusCode)
}
