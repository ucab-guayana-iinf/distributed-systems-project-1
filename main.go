package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	menu "proyecto1.com/main/src/clients/main"
	"proyecto1.com/main/src/count"
	countService "proyecto1.com/main/src/count"
	rpcServer "proyecto1.com/main/src/servers/rpc"
	tcpProcess "proyecto1.com/main/src/servers/tcpProcess"
	tcpThread "proyecto1.com/main/src/servers/tcpThread"
	"proyecto1.com/main/src/utils"
	// udpServer "proyecto1.com/main/src/servers/udp"
)

func start_server(wg *sync.WaitGroup, id int) {
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
}

func runServers() {
	var wg sync.WaitGroup
	countService.InitializeCountService()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go start_server(&wg, i)
	}
	// Ejecutar Consola Local
	wg.Add(1)
	runLocal()

	wg.Wait()
}


func printHelp() {
	fmt.Println("|-------------------------------------|")
	fmt.Println("| Instrucciones Consola Local         |")
	fmt.Println("|-------------------------------------|")
	fmt.Println("| Comando | Descripción               |")
	fmt.Println("|-------------------------------------|")
	fmt.Println("| inc x   | Incrementa la cuenta en x |")
	fmt.Println("| dec x   | Incrementa la cuenta en x |")
	fmt.Println("| restart | Reinicia la cuenta a 0    |")
	fmt.Println("| count   | Imprime la cuenta actual  |")
	fmt.Println("|-------------------------------------|")
}

// Consola local sin el prompt de la libreria
func runLocal() {
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
		case "inc":
			num := utils.StringToInt(arr[1])
			count.SharedCounter.Increment(num, "Local")
		case "dec":
			num := utils.StringToInt(arr[1])
			count.SharedCounter.Decrement(num, "Local")
		case "restart":
			count.SharedCounter.Restart("Local")
		case "count":
			count.SharedCounter.Print()
		default:
			fmt.Println("Instrucción invalida")
			printHelp()
		}
	}
}

func main() {
	var serverFlag bool

	flag.BoolVar(&serverFlag, "server", false, "run servers instead of client")
	flag.Parse()

	if serverFlag {
		runServers()
	} else {
		menu.Start()
	}
}
