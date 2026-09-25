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
	password := r.Form.Get("password")
	log.Printf("Registering new user %s with password %s.\n", user, password)
	w.WriteHeader(http.StatusAccepted)
}
