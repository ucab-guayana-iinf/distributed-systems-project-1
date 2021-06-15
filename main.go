package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"proyecto1.com/main/src/count"
	countService "proyecto1.com/main/src/count"
	rpcServer "proyecto1.com/main/src/servers/rpc"
	tcpProcess "proyecto1.com/main/src/servers/tcpProcess"
	tcpThread "proyecto1.com/main/src/servers/tcpThread"
	"proyecto1.com/main/src/utils"
	// udpServer "proyecto1.com/main/src/servers/udp"
)

func start_server(wg *sync.WaitGroup, id int) {
	// fmt.Printf("[Worker %v]: Started\n", id)
	defer wg.Done()

	switch id {
	case 1:
		tcpThread.Start()
	case 2:
		// apurate miguel
		// udpServer.Start()
	case 3:
		rpcServer.Start()
	case 4:
		if runtime.GOOS == "windows" {
			fmt.Println("Saltando el TCP con procesos en windows pa que no explote")
		} else {
			tcpProcess.Start()
		}
	case 5:
		countService.ProcessMessages()
	}

	fmt.Printf("[Worker %v]: Finished\n", id)
}

func runServers() {
	var wg sync.WaitGroup
	countService.InitializeCountService()

	for i := 1; i <= 5; i++ {
		fmt.Println("[Main]: Starting worker", i)
		wg.Add(1)
		go start_server(&wg, i)
	}
	// Ejecutar Consola Local
	wg.Add(1)
	go runLocal(&wg)

	wg.Wait()
}

func main() {
	runServers()
	// var serverFlag bool

	// flag.BoolVar(&serverFlag, "server", false, "run servers instead of client")
	// flag.Parse()

	// if serverFlag {
	// 	runServers()
	// } else {
	// 	runLocal()
	// }
}

// Consola local sin el prompt de la libreria
func runLocal(wg *sync.WaitGroup) {
	// var tag = "[Local]:"
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if strings.ToUpper(strings.TrimSpace(string(text))) == "STOP" {
			fmt.Println("exiting local console...")
			return
		}

		arr := strings.Split(strings.TrimSpace(text), " ")
		action := arr[0]

		switch action {
		case "increment":
			num := utils.StringToInt(arr[1])
			count.SharedCounter.Increment(num, "Local")
		case "decrement":
			num := utils.StringToInt(arr[1])
			count.SharedCounter.Decrement(num, "Local")
		case "restart":
			count.SharedCounter.Restart("Local")
		case "get":
			count.SharedCounter.Print()
		}
	}
}
