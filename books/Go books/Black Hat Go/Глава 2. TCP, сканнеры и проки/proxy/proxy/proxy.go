package main

import (
	"io"
	"log"
	"net"
)

func proxy(src net.Conn) {
	dst, err := net.Dial("tcp", "localhost:52007")
	if err != nil {
		log.Fatalf("cannot connect to :52007: %v", err)
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalf("cannot copy data: %v", err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

// go run proxy.go
// curl -i -X GET http://localhost:5353
// Hello from blocked server!
func main() {
	l, err := net.Listen("tcp", ":5353")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go proxy(conn)
	}
}
