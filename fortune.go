package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func init() {
	http.HandleFunc("/fortune", fortuneHandler)
}

// fortuneHandler tells the caller's fortune for the topic they name.
func fortuneHandler(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	out, err := exec.Command("sh", "-c", "echo Your fortune for "+topic+": good things are coming").Output()
	if err != nil {
		http.Error(w, "no fortune today", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, string(out))
}
