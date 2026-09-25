package main

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/reset", resetHandler)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	account := r.Form.Get("account")
	password := r.Form.Get("password")
	log.Printf("Resetting password for account %s to %s.\n", account, password)
	w.WriteHeader(http.StatusAccepted)
}
