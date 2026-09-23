package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	http.HandleFunc("/notes", notesHandler)
}

// notesHandler returns the note named by the note parameter from the data
// directory as plain text.
func notesHandler(w http.ResponseWriter, r *http.Request) {
	note := r.URL.Query().Get("note")
	content, err := os.ReadFile(filepath.Join("data", note))
	if err != nil {
		http.Error(w, "no such note", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(content)
}
