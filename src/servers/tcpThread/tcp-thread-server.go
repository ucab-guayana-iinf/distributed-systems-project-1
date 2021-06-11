// package tcpThreadServer
package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"proyecto1.com/main/src/count"
)

// Connected clients
var clientCount = 0

func handleConnection(c net.Conn) {
	defer c.Close()
	fmt.Println("[TCP Server]: Client connected with IP", c.RemoteAddr().String())
	for {
		// Get messages from clients
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("[TCP Server]: Error leyendo el input de la conexion:", err)
			return
		}

		// Parse client message
		temp := strings.TrimSpace(string(netData))

		// Exit condition
		if strings.ToUpper(temp) == "STOP" {
			fmt.Println("[TCP Server]: Client disconnected")
			clientCount--
			break
		}

		fmt.Println("[TCP Server]: Client said", temp)
		if temp == "increment" {
			count.SharedCounter.Increment(1, "TCP Thread Server")
		} else if temp == "decrement" {
			count.SharedCounter.Decrement(1, "TCP Thread Server")
		}

		// Respond to client with clientCount
		counter := strconv.Itoa(clientCount) + "\n"
		message := "Qlq mano, clientes conectados: " + counter
		c.Write([]byte(string(message)))
	}
}

// func Start() {
func main() {
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
		// Serve each client with a goroutine
		go handleConnection(c)
		clientCount++
	}
}
