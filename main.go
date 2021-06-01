package main

import (
	"fmt"
	"sync"

	tcpServer "proyecto1.com/main/src/servers/tcp"
	rpcServer "proyecto1.com/main/src/servers/rpc"
	udpServer "proyecto1.com/main/src/servers/udp"
)

func start_server(wg *sync.WaitGroup, id int) {
	fmt.Printf("[Worker %v]: Started\n", id)
	defer wg.Done()

	switch id {
	case 1:
		tcpServer.Start()
	case 2:
		udpServer.Start()
	case 3:
		rpcServer.Start()
	}

	fmt.Printf("[Worker %v]: Finished\n", id)
}

func main() {
	fmt.Println("[Main]: Started")
	var wg sync.WaitGroup

	// TODO: cli que permita?
	// - Imprimir status de los servicios
	// - Conectarse al CLI local
	// - Conectarse al CLI remoto
	// - Matar todo

	for i := 1; i <= 3; i++ {
		fmt.Println("[Main]: Starting worker", i)
		wg.Add(1)
		go start_server(&wg, i)
	}

	wg.Wait()
}
