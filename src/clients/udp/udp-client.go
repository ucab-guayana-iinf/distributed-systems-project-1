package udpClient

import (
	"errors"
	"fmt"
	"net"
)

func InitUDPClientConnection() (net.Conn, error) {
	address := "localhost:2002"

	resolveAddr, err := net.ResolveUDPAddr("udp4", address)
	c, err := net.DialUDP("udp4", nil, resolveAddr)

	if err != nil {
		fmt.Println(err)

		return nil, errors.New("error connecting to localhost:2002")
	}

	return c, nil
}

func InvokeUDPClientCall(client net.Conn, operation string, num int) {
	fmt.Fprintf(client, "%v %v\n", operation, num)
}
