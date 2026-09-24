package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func init() {
	http.HandleFunc("/repeat", repeatHandler)
}

// repeatHandler echoes a word back the requested number of times, at most
// 100.
func repeatHandler(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	n, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		http.Error(w, "bad count", http.StatusBadRequest)
		return
	}
	if n < 0 || n > 100 {
		http.Error(w, "count out of range", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, strings.TrimSpace(strings.Repeat(word+" ", n)))
}
