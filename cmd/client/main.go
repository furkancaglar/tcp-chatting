package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Println("error dial ", err)
		return
	}
	defer conn.Close()

	log.Println("connection successful")

	go readFromClient(conn)

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("Enter text: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error read from console ", err)
			return
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Println("err write ", err)
			continue
		}

	}

}

func readFromClient(conn net.Conn) error {
	for {
		data := make([]byte, 100)
		_, err := conn.Read(data)
		if err != nil {
			return err
		}

		fmt.Println("\n\nmessage from client \n\n", string(data))
	}
}
