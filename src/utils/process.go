package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

// TODO: use to validate count min/max
// const MaxInt = int(MaxUint >> 1)
// const MinInt = -MaxInt - 1

func StartProcess() {
	log.Printf("Haha business")

	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func IntToString(n int) string {
	return strconv.Itoa(n)
}

func CreateProcessTest() {
	var cmd string

	if runtime.GOOS == "windows" {
		cmd = "cmd.exe"
	} else {
		cmd = "/bin/bash"
	}

	// Returns a pointer to the process
	proc := exec.Command(cmd)

	// Get process' standard input
	stdin, err := proc.StdinPipe()
	if err != nil {
		panic(err.Error())
	}
	defer stdin.Close()

	// Get process' standard output
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	// Function that reads the outputs of the process and prints them
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println("From scanner:", scanner.Text())
		}
	}()

	// Start the cmd process
	err = proc.Start()
	if err != nil {
		panic(err.Error())
	}

	// Write to the process' console and execute a command
	_, err = io.WriteString(stdin, "ping google.com\n")
	if err != nil {
		panic(err.Error())
	}

	time.Sleep(time.Second * 20)
	proc.Process.Kill()
}
