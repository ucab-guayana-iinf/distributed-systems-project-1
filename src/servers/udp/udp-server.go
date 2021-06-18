package udpServer

import (
	"bufio"
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

	//	Returns an address of a endpoint UDP
	resolveAddr, err := net.ResolveUDPAddr("udp4", PORT)

	// Create a UDP server
	c, err := net.ListenUDP("udp4", resolveAddr)

	if err != nil {
		fmt.Println("Error creando el servidor UDP", err)

		return
	}

	defer c.Close()

	fmt.Println(tag, "Running in http://localhost"+PORT)

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			fmt.Println(tag, "Error leyendo el input de la conexion:", err)
			return
		}

		// Parse client message
		temp := strings.TrimSpace(string(netData))

		arr := strings.Split(temp, " ")
		action := arr[0]

		switch action {
		case utils.INCREMENT:
			num := utils.StringToInt(arr[1])
			count.Produce(action, "UDP Server", num)
		case utils.DECREMENT:
			num := utils.StringToInt(arr[1])
			count.Produce(action, "UDP Server", num)
		case utils.RESTART:
			count.Produce(action, "UDP Server", 0)
		case utils.GET_COUNT:
			count.Produce(action, "UDP Server", 0)
		}

	}
}
