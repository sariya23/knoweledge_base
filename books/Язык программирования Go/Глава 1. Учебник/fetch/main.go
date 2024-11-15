package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// Помним, что за Get стоит дефолтный клиент без таймаутов!
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("cannot fetch url: %s\n", url)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("cannot read body: %v\n", b)
			os.Exit(1)
		}
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				fmt.Printf("cannot close body: %v", err)
				os.Exit(1)
			}
		}()
		fmt.Printf("%s\n", b)

	}
}
