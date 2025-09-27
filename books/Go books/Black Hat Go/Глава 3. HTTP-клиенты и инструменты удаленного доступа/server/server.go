package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			io.WriteString(w, "hello from /")

			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	})
	mux.HandleFunc("/q", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			io.WriteString(w, "hello from /q")
			return
		}
		return
	})

	err := http.ListenAndServe(":5353", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
