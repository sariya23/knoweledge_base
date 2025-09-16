package main

import (
	"fmt"
	"sync"
)

func asyncSum() (int, int) {
	ch1 := make(chan int, 1_000_000)
	ch2 := make(chan int, 1_000_000)

	var wg sync.WaitGroup

	var a1, a2 []int

	for i := 0; i < 1_000_000; i++ {
		a1 = append(a1, i)
		a2 = append(a2, i)
	}

	for i := 0; i < len(a1); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch1 <- a1[i]
			ch2 <- a2[i]
		}()
	}
	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()

	var sum1, sum2 int

	for v := range ch1 {
		sum1 += v
	}
	for v := range ch2 {
		sum2 += v
	}
	return sum1, sum2
}

func noAsync() (int, int) {
	var sum1, sum2 int

	for i := 0; i < 1_000_000; i++ {
		sum1 += i
	}

	for i := 0; i < 1_000_000; i++ {
		sum2 += i
	}
	return sum1, sum2
}

type Robot struct {
	hp  int
	dmg int
}

func main() {
	old := Robot{2, 5}
	new_ := Robot{2, 5}

	for i := 0; i < 20; i++ {
		old.dmg += 3
		old.hp += 1

		new_.dmg += 2
		new_.hp += 1
	}

	fmt.Println(old)
	fmt.Println(new_)
}
