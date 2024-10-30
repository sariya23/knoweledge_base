package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from blocked server!")
	})

	err := http.ListenAndServe(":52007", nil)
	if err != nil {
		log.Fatalf("cannot start server at :52007: %v", err)
	}
}
