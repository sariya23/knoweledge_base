package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	res := make(chan int, 1)

	for i := 0; i < 11; i++ {
		go work(ctx, i, res)
	}

	foundBy := <-res
	fmt.Println("result found by", foundBy)
	cancel()
	time.Sleep(time.Second)
}

func work(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println("worker done")
		out <- workerNum
	}
}
