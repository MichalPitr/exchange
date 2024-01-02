package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func handleConnection(c net.Conn) {
	// Parse

	// Add to queue
}

func main() {
	var port int
	flag.IntVar(&port, "port", 2020, "server port")
	flag.Parse()
	if port < 1024 {
		log.Println("Using reserved port.")
		os.Exit(1)
	}

	address := fmt.Sprintf("localhost:%d", port)
	ln, err := net.Listen("tcp", address)
	defer ln.Close()

	if err != nil {
		log.Printf("Error creating a server: %v", err)
		os.Exit(1)
	}
	log.Printf("Starting server on %s\n", address)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		handleConnection(conn)
	}
}
