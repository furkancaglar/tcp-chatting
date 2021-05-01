package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Println("error listen ", err)
		return
	}

	data := make([]byte, 0, 100)

	data = []byte("hello server!")

	_, err = conn.Write(data)
	if err != nil {
		log.Println("error listen ", err)
		return
	}

}
