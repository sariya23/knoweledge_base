package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	req, err := http.NewRequest("DELETE", "https://google.com/robots.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
	var client http.Client
	reps, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(reps)

	form := url.Values{}
	form.Add("foo", "bar")
	req, err = http.NewRequest("PUT", "https://google.com/robots.txt", strings.NewReader(form.Encode()))
	resp, err := client.Do(req)
	d := []byte{}
	resp.Body.Read(d)
	fmt.Println()
	fmt.Println(string(d))
	resp.Body.Close()
}
