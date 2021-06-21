package tcpThread

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"proyecto1.com/main/src/count"
	"proyecto1.com/main/src/utils"
)

// Connected clients
var clientCount = 0
var tag = "[TCP Thread Server]:"

const PORT = ":2020"

func handleConnection(c net.Conn, currClient string) {
	defer c.Close()
	log.Println(tag, "Client connected with IP", c.RemoteAddr().String())

	// Enviar su numero de cliente
	clientId := c.RemoteAddr().String() + "\n"
	c.Write([]byte(clientId))

	for {
		// Get messages from clients
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(tag, "Error leyendo el input de la conexion:", err)
			return
		}

		// Parse client message
		temp := strings.TrimSpace(string(netData))

		arr := strings.Split(temp, " ")
		action := arr[0]

		switch action {
		case utils.STOP:
			log.Println(tag, " Client", c.RemoteAddr().String(), "disconnected")
			return
		case utils.INCREMENT:
			num := utils.StringToInt(arr[1])
			count.Produce(action, "TCP Thread Server", num)
		case utils.DECREMENT:
			num := utils.StringToInt(arr[1])
			count.Produce(action, "TCP Thread Server", num)
		case utils.RESTART:
			count.Produce(action, "TCP Thread Server", 0)
		case utils.GET_COUNT:
			count.Produce(action, "TCP Thread Server"+c.RemoteAddr().String(), 0)
		}
	}
}

func Start() {
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
		clientCount++
		go handleConnection(c, utils.IntToString(clientCount))
	}
}
