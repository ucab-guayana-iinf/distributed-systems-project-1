// package httpServer
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// Connected clients
var count = 0

func main() {
	const PORT = ":2020"
	fmt.Println("[TCP Server]: Starting")

	// Make the TCP server
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error creando el servidor TCP", err)
		return
	}
	defer listener.Close()

	// Accepts incoming client connections
	c, err := listener.Accept()
	if err != nil {
		fmt.Println("Error aceptando la conexion del cliente:", err)
		return
	}

	// Interaction with clients
	for {
		// Get messages from clients
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error leyendo el input de la conexion:", err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("Client: ", string(netData))

		// Send current time to clients
		t := time.Now()
		myTime := "Epale mano, son las " + t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
