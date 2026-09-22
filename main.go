// Package main is a deliberately vulnerable sample service used to exercise
// the patchy security pipeline. Do not deploy it.
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// fileHandler serves a file from the data directory by name.
func fileHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	content, err := os.ReadFile(filepath.Join("data", name))
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Write(content)
}

// greetHandler greets the caller by name.
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Hello, %s</h1>", html.EscapeString(name))
}

// runHandler echoes the cmd parameter back via the echo binary.
func runHandler(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("echo", r.URL.Query().Get("cmd")).Output()
	if err != nil {
		http.Error(w, "command failed", http.StatusInternalServerError)
		return
	}
	w.Write(out)
}

func main() {
	http.HandleFunc("/file", fileHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/run", runHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
