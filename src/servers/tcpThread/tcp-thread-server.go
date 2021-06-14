package tcpThread

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"proyecto1.com/main/src/count"
	"proyecto1.com/main/src/utils"
)

// Connected clients
var clientCount = 0
var tag = "[TCP Thread Server]:"

func handleConnection(c net.Conn) {
	defer c.Close()
	fmt.Println(" Client connected with IP", c.RemoteAddr().String())
	for {
		// Get messages from clients
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(tag, "Error leyendo el input de la conexion:", err)
			return
		}

		// Parse client message
		temp := strings.TrimSpace(string(netData))

		// Exit condition
		if strings.ToUpper(temp) == "STOP" {
			fmt.Println(tag, "Client disconnected")
			clientCount--
			break
		}

		arr := strings.Split(temp, " ")
		action := arr[0]

		fmt.Println(tag, "Client said", temp)
		if action == "Increment" || action == "Decrement" {
			num := utils.StringToInt(arr[1])
			count.Produce(action, "TCP Thread Server", num)
		} else if action == "Restart" {
			fmt.Println(tag, "Restart count")
			// count.Produce(action, "TCP Thread Server", num)
		} else if action == "Get" {
			fmt.Println(tag, "Get count")
		}

		// Respond to client with clientCount
		// counter := strconv.Itoa(clientCount) + "\n"
		// message := "Qlq mano, clientes conectados: " + counter
		// c.Write([]byte(string(message)))
	}
}

func Start() {
	const PORT = ":2020"
	fmt.Println(tag, "Starting")

	// Make the TCP Thread server listener
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println("Error creando el servidor TCP", err)
		return
	}
	defer l.Close()
	fmt.Println(tag, "Running in http://localhost"+PORT)

	// Accept client connections
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(tag, "Error aceptando la conexion del cliente:", err)
			return
		}
		// Serve each client with a goroutine
		go handleConnection(c)
		clientCount++
	}
}
