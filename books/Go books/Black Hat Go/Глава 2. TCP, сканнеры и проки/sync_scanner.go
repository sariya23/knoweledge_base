package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 1025; i++ {
		addr := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("port %d closed or filtered\n", i)
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
