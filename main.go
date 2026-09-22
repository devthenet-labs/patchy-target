// Package main is a deliberately vulnerable sample service used to exercise
// the patchy security pipeline. Do not deploy it.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	fmt.Fprintf(w, "<h1>Hello, %s</h1>", name)
}

func main() {
	http.HandleFunc("/file", fileHandler)
	http.HandleFunc("/greet", greetHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
