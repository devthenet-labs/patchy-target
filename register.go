package main

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/register", registerHandler)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	user := r.Form.Get("user")
	log.Printf("Registering new user %s.\n", user)
	w.WriteHeader(http.StatusAccepted)
}
