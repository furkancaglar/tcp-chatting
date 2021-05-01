package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println("error listen ", err)
		return
	}
	defer ln.Close()
	log.Println("server started listening on port 8000")

	conn, err := ln.Accept()
	if err != nil {
		log.Println("error listen ", err)
		return
	}

	data := make([]byte, 100)
	_, err = conn.Read(data)
	if err != nil {
		log.Println("error listen ", err)
		return
	}

	log.Println(string(data)) // it converts the []byte to string and prints them
	// log.Println(data)  // logs as []byte (not human readable)

}
