package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second * 5}
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch, &c)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Println("end")
}

func fetch(url string, ch chan<- string, c *http.Client) {
	fmt.Printf("fetch url: %v\n", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		ch <- err.Error()
		return
	}
	resp, err := c.Do(req)
	if err != nil {
		ch <- err.Error()
		return
	}
	_, err = io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- err.Error()
		return
	}
	ch <- fmt.Sprintf("fetched url: %s\n", url)
	return
}
