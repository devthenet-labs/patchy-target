package main

import (
	"encoding/json"
	"net/http"
)

// sha and built are set at build time via -ldflags "-X main.sha=... -X main.built=...".
// They default to "unknown" when not injected by a build.
var (
	sha   = "unknown"
	built = "unknown"
)

func init() {
	http.HandleFunc("/version", versionHandler)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"sha": sha, "built": built})
}
