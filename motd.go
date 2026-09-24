package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func init() {
	http.HandleFunc("/motd", motdHandler)
}

// motdHandler greets the named visitor with the message of the day.
func motdHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	message := fmt.Sprintf("Message of the day for %s: ship small changes", name)
	out, err := exec.Command("echo", message).Output()
	if err != nil {
		http.Error(w, "no message today", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, string(out))
}
