package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/banner", bannerHandler)
}

// bannerHandler renders the text parameter as the page's announcement banner.
func bannerHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<div class=\"banner\">%s</div>", text)
}
