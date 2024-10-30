package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalf("cannot copt data: %s", err)
	}
	fmt.Println("qwe")
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:5353")
	if err != nil {
		panic(err)
	}
	log.Printf("listen: %s\n", l.Addr().String())
	for {
		conn, err := l.Accept()
		log.Println("recieved data")
		if err != nil {
			log.Fatalf("cannot recieve data: %v\n", err)
		}
		go echo(conn)
	}
}
