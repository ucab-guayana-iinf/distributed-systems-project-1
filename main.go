package main

import (
	"fmt"
	"runtime"
	"sync"

	countService "proyecto1.com/main/src/count"
	rpcServer "proyecto1.com/main/src/servers/rpc"
	tcpProcess "proyecto1.com/main/src/servers/tcpProcess"
	tcpThread "proyecto1.com/main/src/servers/tcpThread"
	udpServer "proyecto1.com/main/src/servers/udp"
)

func start_server(wg *sync.WaitGroup, id int) {
	// fmt.Printf("[Worker %v]: Started\n", id)
	defer wg.Done()

	switch id {
	case 1:
		tcpThread.Start()
	case 2:
		udpServer.Start()
	case 3:
		rpcServer.Start()
	case 4:
		if runtime.GOOS == "windows" {
			fmt.Println("Saltando el TCP con procesos en windows")
		} else {
			tcpProcess.Start()
		}
	case 5:
		countService.ProcessMessages()
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

	for i := 1; i <= 4; i++ {
		fmt.Println("[Main]: Starting worker", i)
		wg.Add(1)
		go start_server(&wg, i)
	}

	wg.Wait()
}
