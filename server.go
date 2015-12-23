package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("tcp Listen error:", err)
	}

	for {

		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error on listener Accept: ", err)
		}

		go func(c net.Conn) {
			defer c.Close()

			for {
				m, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					break
				}
				fmt.Print("Message received: ", m)
				fmt.Fprint(c, "Response -- ", m)
			}
		}(conn)
	}

	fmt.Println("server init")
}
