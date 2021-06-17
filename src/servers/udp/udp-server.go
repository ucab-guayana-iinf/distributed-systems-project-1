package udpServer

import (
	"fmt"
	"net"
	"bufio"
	"strings"
)

func Start() {
	const PORT = ":2002"

	tag := "[UDP Server]:"

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
	
	}
}
