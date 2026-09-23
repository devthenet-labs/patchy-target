// Package main is a deliberately vulnerable sample service used to exercise
// the patchy security pipeline. Do not deploy it.
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

// runHandler echoes the cmd parameter back through the shell.
func runHandler(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("sh", "-c", "echo "+r.URL.Query().Get("cmd")).Output()
	if err != nil {
		http.Error(w, "command failed", http.StatusInternalServerError)
		return
	}
	w.Write(out)
}

// redirectHandler sends the caller on to the page named by next. To guard
// against open redirects (CWE-601), next must be a same-site relative path.
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	// Some browsers treat backslashes as forward slashes, so normalize
	// them before parsing to avoid scheme-relative smuggling (e.g. "/\evil.com").
	target, err := url.Parse(strings.ReplaceAll(next, "\\", "/"))
	if err != nil || target.Host != "" || target.Scheme != "" {
		http.Error(w, "invalid redirect target", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, target.String(), http.StatusFound)
}

func main() {
	http.HandleFunc("/file", fileHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/run", runHandler)
	http.HandleFunc("/go", redirectHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
