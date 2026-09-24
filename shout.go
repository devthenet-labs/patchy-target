package main

import (
	"fmt"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/shout", shoutHandler)
}

// shoutHandler echoes the msg parameter back in upper case.
func shoutHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>%s!</h1>", strings.ToUpper(msg))
}
