package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "blog.irrelevant.ninja:80")
	if err != nil {
		log.Fatalln("Unable to connect to unreachable host")
	}
	defer dst.Close()

	// Run in goroutine to prevent blocking
	go func() {
		// Copy source output to destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Copy destination output back to source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Listen on port 80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
