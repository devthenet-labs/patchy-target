package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/echo", echoHandler)
}

// echoHandler echoes the msg parameter back as an HTML paragraph.
func echoHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<p>%s</p>", msg)
}
