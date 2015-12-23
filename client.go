package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprint(conn, text)

		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message form server: ", response)
	}
}
