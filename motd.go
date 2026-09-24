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
	out, err := exec.Command("sh", "-c", "echo Message of the day for "+name+": ship small changes").Output()
	if err != nil {
		http.Error(w, "no message today", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, string(out))
}
