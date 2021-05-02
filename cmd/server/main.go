package main

import (
	"fmt"
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

	connMap := make(map[int]net.Conn)

	counter := 1
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("error accept ", err)
			continue
		}
		log.Println("new conn ", counter)
		connMap[counter] = conn

		go readFromClient(conn, counter, connMap)

		counter++
	}

}

func readFromClient(conn net.Conn, connID int, connMap map[int]net.Conn) {
	for {
		data := make([]byte, 100)
		_, err := conn.Read(data)
		if err != nil {
			log.Printf("err read conn id=%v err=%v\n", connID, err)
			continue
		}

		for cID, c := range connMap {
			if cID == connID {
				continue
			}

			_, err := c.Write([]byte(fmt.Sprintf("conn id =%v , msg= %v\n", cID, string(data))))
			if err != nil {
				log.Printf("err write conn id=%v err=%v\n", cID, err)
				continue
			}
		}

	}
}
