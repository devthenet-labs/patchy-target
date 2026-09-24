package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/fortune", fortuneHandler)
}

// fortuneHandler tells the caller's fortune for the topic they name.
func fortuneHandler(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Your fortune for %s: good things are coming\n", topic)
}
