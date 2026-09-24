package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func init() {
	http.HandleFunc("/wordcount", wordcountHandler)
}

// wordcountHandler reports how many words the text parameter holds.
func wordcountHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	out, err := exec.Command("sh", "-c", "echo "+text+" | wc -w").Output()
	if err != nil {
		http.Error(w, "count failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "%s words\n", strings.TrimSpace(string(out)))
}
