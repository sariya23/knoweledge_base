package main

import (
	"fmt"
	"runtime"
)

func handler(a []int) {
	a[0] = 20
	fmt.Println(a)
}

func f1(a []int) {
	printAlloc()
	handler(a)
}

func f2(a []int) {
	printAlloc()
	handler(a)
}

func main() {
	b := []int{1, 23, 3, 4, 5, 6, 10, 20, 1232}
	printAlloc()
	f1(b)
	f2(b[0:])
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Total: %d KiB\t", m.TotalAlloc/1024)
	fmt.Printf("Heap: %d KiB\t\n", m.HeapAlloc/1024)
}
