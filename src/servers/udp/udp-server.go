package udpServer

import (
	"fmt"
	"log"
	"net"
)

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	// 0 - 1: ID
	// 2: QR(1): Opcode(4)
	buf[2] |= 0x80 // Set QR bit

	pc.WriteTo(buf, addr)
}

func Start() {
	fmt.Println("[UDP Server]: Starting")
	var port = "2002"

	pc, err := net.ListenPacket("udp", ":" + port)
	fmt.Println("[UDP Server]: Running in http://localhost:" + port)
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		go serve(pc, addr, buf[:n])
	}
}
