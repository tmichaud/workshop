package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

//from net package's Listen function
//type Listener interface {
//	// Accept waits for and returns the next conenction to the listener
//	Accept() (c Conn, err error)
//
//	// close closes the listener
//	// Any blocked Accept operations will be unblocked and return errors
//	Close() error
//
//	// Addr returns the listener's network adddress.
//	Addr() Addr
//}

func server() {
	fmt.Println("Server starting")
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("net.Listen returned error: ", err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("Error acccepting a connection:", err)
			continue
		}
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println("Error decoding message:", err)
	} else {
		fmt.Println("Server Received: ", msg)
	}
}

func client() {
	fmt.Println("Starting client")
	// Connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("Error in client connecting to 127.0.0.1:9999", err)
		return
	}

	// send the message
	msg := "Hello World from a go client"
	fmt.Println("Sending from client: ", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println("Error in client sending message:", err)
	}

	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
