package main

import (
	"encoding/json"
	"net/http"
)

func init() {
	http.HandleFunc("/ready", readyHandler)
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"ready": true})
}
