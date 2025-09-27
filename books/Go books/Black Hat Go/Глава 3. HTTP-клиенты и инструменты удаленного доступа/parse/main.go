package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com/robots.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	resp.Body.Close()
}
