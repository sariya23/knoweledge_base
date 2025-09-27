package main

import (
	"fmt"
	"log"
	"net"
	"sort"
	"time"
)

func worker(ports, result chan int) {
	for p := range ports {
		addr := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.DialTimeout("tcp", addr, time.Second)
		if err != nil {
			log.Println("err")
			result <- 0
			continue
		}
		conn.Close()
		log.Println("conn close")
		result <- p
		log.Println("port handled")
	}

}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < 1025; i++ {
			log.Println("add port to chan")
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		log.Println("wait data from results chan")
		port := <-results
		log.Println("get data from results chan")
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	fmt.Println("end")
	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, v := range openPorts {
		fmt.Printf("%d open\n", v)
	}
}
