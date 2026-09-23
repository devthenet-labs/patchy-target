package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/welcome", welcomeHandler)
}

// welcomeHandler welcomes the caller to the team named by the team parameter.
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	team := r.URL.Query().Get("team")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h2>Welcome to team %s</h2>", team)
}
