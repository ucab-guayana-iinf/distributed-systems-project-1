// package tcpClient
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// To stop connection, send STOP message to server
func main() {
	address := "localhost:2020"
	fmt.Println("[TCP Client]: Starting")

	c, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// Read user input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		// Send input to the server
		fmt.Fprintf(c, text+"\n")

		// Read messages from the server
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("Server: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
