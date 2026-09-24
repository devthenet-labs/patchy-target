package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
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

// checkURLClient calls caller-supplied URLs. Its dialer resolves and
// validates the destination address on every connection attempt (including
// redirects), so callers can't reach non-public network destinations.
var checkURLClient = &http.Client{
	Timeout: 5 * time.Second,
	Transport: &http.Transport{
		DialContext: safeDialContext,
	},
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return validatePublicHTTPURL(req.URL)
	},
}

// checkURLHandler reports the status of a caller-supplied service URL.
func checkURLHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")

	parsed, err := url.Parse(target)
	if err != nil || validatePublicHTTPURL(parsed) != nil {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	resp, err := checkURLClient.Get(parsed.String())
	if err != nil {
		http.Error(w, "service unreachable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	fmt.Fprintf(w, "service status: %d\n", resp.StatusCode)
}

// validatePublicHTTPURL rejects anything other than a plain http(s) request
// to a named host. It does not resolve the host itself: name resolution and
// the IP-range check happen per-connection in safeDialContext, so there is
// no gap between validating a name and actually connecting to it.
func validatePublicHTTPURL(u *url.URL) error {
	if u == nil {
		return fmt.Errorf("missing url")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("unsupported url scheme")
	}
	if u.Hostname() == "" {
		return fmt.Errorf("missing host")
	}
	return nil
}

// safeDialContext resolves the target host and only dials addresses that are
// public and routable. This blocks the SSRF class where an attacker points
// the request at loopback, private, link-local, or other reserved addresses
// (including the 169.254.169.254 cloud metadata endpoint). Resolution and
// validation happen at dial time for every connection, including redirects,
// which also closes the DNS-rebinding window between an earlier lookup and
// the actual connection.
func safeDialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	var resolver net.Resolver
	ips, err := resolver.LookupIP(ctx, "ip", host)
	if err != nil {
		return nil, err
	}

	var dialer net.Dialer
	var lastErr error
	for _, ip := range ips {
		if !isPublicIP(ip) {
			continue
		}
		conn, dialErr := dialer.DialContext(ctx, network, net.JoinHostPort(ip.String(), port))
		if dialErr == nil {
			return conn, nil
		}
		lastErr = dialErr
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("no public address found for host %q", host)
	}
	return nil, lastErr
}

// isPublicIP reports whether ip is a publicly routable address, excluding
// loopback, private, link-local, unspecified, and multicast ranges.
func isPublicIP(ip net.IP) bool {
	return !ip.IsLoopback() &&
		!ip.IsPrivate() &&
		!ip.IsLinkLocalUnicast() &&
		!ip.IsLinkLocalMulticast() &&
		!ip.IsUnspecified() &&
		!ip.IsMulticast()
}
