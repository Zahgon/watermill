package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", http.DefaultServeMux)
}
