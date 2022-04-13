package main

import (
	"fmt"
	"net/http"
)

func firstGreeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Welcome! Although this site is still lacking in content, I'll add as many as I can in the future!")
}

func main() {
	http.HandleFunc("/", firstGreeting)
	http.ListenAndServe(":80", nil)
}
