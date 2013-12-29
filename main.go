package main

import (
	"fmt"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	frame, err := ReadFrame(conn)
	if err != nil {
		log.Print(err.Error())
	}
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:61613")
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err.Error())
			continue
		}
		go handleConn(conn)
	}
}
