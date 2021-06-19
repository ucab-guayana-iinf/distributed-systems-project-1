package udpServer

import (
	"fmt"
	"net"
	"strings"

	"proyecto1.com/main/src/count"
	"proyecto1.com/main/src/utils"
)

func Start() {
	const PORT = ":2002"
	var tag = "[UDP Server]:"

	fmt.Println(tag, "Starting")

	// Returns an address of a endpoint UDP
	resolveAddr, err := net.ResolveUDPAddr("udp4", PORT)

	// Create a UDP server
	c, err := net.ListenUDP("udp4", resolveAddr)

	if err != nil {
		fmt.Println("Error creando el servidor UDP", err)
		return
	}

	defer c.Close()

	fmt.Println(tag, "Running in http://localhost"+PORT)

	buffer := make([]byte, 1024)

	for {
		n, addr, err := c.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println(tag, "Error leyendo el input de la conexion:", err)
			return
		}

		clientDataUdp := strings.Trim(string(buffer[:n]), "\n")
		arr := strings.Split(clientDataUdp, " ")
		action := arr[0]
		num := utils.StringToInt(arr[1])

		switch action {
		case utils.INCREMENT:
			count.Produce(action, "UDP Server", num)
		case utils.DECREMENT:
			count.Produce(action, "UDP Server", num)
		case utils.RESTART:
			count.Produce(action, "UDP Server", 0)
		case utils.GET_COUNT:
			fmt.Println("UDP Client:", addr)
			count.Produce(action, "UDP Server", 0)
		}
	}
}
