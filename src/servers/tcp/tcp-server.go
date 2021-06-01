// package main
package tcpServer

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Connected clients
var count = 0

func handleConnection(c net.Conn) {
	defer c.Close()
	fmt.Println("[TCP Server]: Client connected with IP", c.RemoteAddr().String())
	for {
		// Get messages from clients
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error leyendo el input de la conexion:", err)
			return
		}

		// If STOP received, close connection
		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			fmt.Println("[TCP Server]: Client disconnected")
			count--
			break
		}

		fmt.Println("[TCP Server]: Client said", temp)

		// Send client count to client
		counter := strconv.Itoa(count) + "\n"
		message := "Qlq mano, clientes conectados:" + counter
		c.Write([]byte(string(message)))
	}
}

func Start() {
	const PORT = ":2020"
	fmt.Println("[TCP Server]: Starting")

	// Make the TCP server listener
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println("Error creando el servidor TCP", err)
		return
	}
	defer l.Close()
	fmt.Println("[TCP Server]: Running in http://localhost" + PORT)

	// Accept client connections
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error aceptando la conexion del cliente:", err)
			return
		}
		// Serve each client with 1 goroutine
		go handleConnection(c)
		count++
	}
}
