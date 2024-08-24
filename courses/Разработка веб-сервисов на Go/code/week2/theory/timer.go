package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.AfterFunc(1*time.Second, func() { fmt.Println("hello") })
	time.Sleep(time.Second * 2)
	timer.Stop()
}
