// Package main is a deliberately vulnerable sample service used to exercise
// the patchy security pipeline. Do not deploy it.
package main

import (
	"database/sql"
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

	dataDir, err := filepath.Abs("data")
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	path, err := filepath.Abs(filepath.Join(dataDir, name))
	if err != nil || (path != dataDir && !strings.HasPrefix(path, dataDir+string(os.PathSeparator))) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	content, err := os.ReadFile(path)
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

// db is the user store; main opens it when USERS_DSN is set.
var db *sql.DB

// userHandler looks up a user's email address by name.
func userHandler(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		http.Error(w, "user store unavailable", http.StatusServiceUnavailable)
		return
	}
	name := r.URL.Query().Get("name")
	var email string
	err := db.QueryRow("SELECT email FROM users WHERE name = ?", name).Scan(&email)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, email)
}

func main() {
	if dsn := os.Getenv("USERS_DSN"); dsn != "" {
		var err error
		if db, err = sql.Open(os.Getenv("USERS_DRIVER"), dsn); err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/file", fileHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/run", runHandler)
	http.HandleFunc("/go", redirectHandler)
	http.HandleFunc("/user", userHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
