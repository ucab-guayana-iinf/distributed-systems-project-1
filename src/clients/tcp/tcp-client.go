package tcpClient

import (
	"errors"
	"fmt"
	"net"
)

func InitTCPClientConnection() (net.Conn, error) {
	address := "localhost:2020"
	c, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error connecting to localhost:2020")
	}
	return c, nil
}

func InitTCPProcessClientConnection() (net.Conn, error) {
	address := "localhost:2021"
	c, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error connecting to localhost:2020")
	}
	return c, nil
}

func InvokeTCPClientCall(client net.Conn, operation string, num int) {
	fmt.Fprintf(client, "%v %v\n", operation, num)
}
