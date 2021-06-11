// package tcpProcServer
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"proyecto1.com/main/src/count"
)

// Connected clients
var clientCount = 0
var exe string
var err error

const PORT = ":2020"

func main() {
	fmt.Println("[TCP Process Server]: Starting")

	// determine executable
	if exe, err = os.Executable(); err != nil {
		fmt.Println("Error obteniendo ejecutable")
		return
	}

	// flags
	var optChild bool

	flag.BoolVar(&optChild, "worker", false, "start as a worker process (internal only)")
	flag.Parse()

	if optChild { // we are in child process
		childMain()
	} else {
		parentMain()
	}
}

func parentMain() {
	// Make the TCP server listener
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error creando el servidor TCP", err)
		return
	}
	defer l.Close()
	fmt.Println("[TCP Process Server]: Running in http://localhost" + PORT)

	// Accept client connections
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("[TCP Process Server]: Error aceptando la conexion del cliente:", err)
			return
		}
		clientCount++

		// Get the fd copy of the TCP connection
		var f *os.File

		// Serve each client with a process
		if f, err = c.(*net.TCPConn).File(); err != nil {
			fmt.Println(err)
			log.Fatal("[TCP Process Server]: failed to obtain connection fd")
			return
		}
		defer f.Close() // After fd is passed to the child process, it can also be safely closed

		// Create a child process
		cmd := exec.Command(exe, append([]string{"-worker"}, os.Args[1:]...)...)
		cmd.Dir, _ = os.Getwd()
		cmd.Env = os.Environ()
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.ExtraFiles = []*os.File{f} // Here fd is passed to the child process
		if err = cmd.Start(); err != nil {
			log.Fatal("[TCP Process Server]: failed to start child process")
			return
		}
	}
}

func childMain() {
	// fd 0 = stdin, fd 1 = stdout, fd 2 = stderr
	// Get connection from fd 3
	var c net.Conn
	if c, err = net.FileConn(os.NewFile(3, "connection")); err != nil {
		log.Fatal("[TCP Process Server Child]: failed to obtain connection")
		return
	}
	defer c.Close()

	for {
		// Get messages from clients
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("[TCP Process Server Child]: Error leyendo el input de la conexion:", err)
			return
		}

		// Parse client message
		temp := strings.TrimSpace(string(netData))

		// Exit condition
		if strings.ToUpper(temp) == "STOP" {
			fmt.Println("[TCP Process Server Child]: Client disconnected")
			clientCount--
			break
		}

		fmt.Println("[TCP Process Server Child]: Client said", temp)
		if temp == "increment" {
			count.SharedCounter.Increment(1, "TCP Thread Server")
		} else if temp == "decrement" {
			count.SharedCounter.Decrement(1, "TCP Thread Server")
		}

		// Respond to client with clientCount
		counter := strconv.Itoa(clientCount) + "\n"
		message := "Qlq mano, clientes conectados: " + counter
		c.Write([]byte(string(message)))
	}
}
