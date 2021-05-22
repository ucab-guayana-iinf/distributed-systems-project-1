package main

import (
	"fmt"
	"sync"

	httpServer "proyecto1.com/main/http"
)

func start_server(wg *sync.WaitGroup, id int) {
	fmt.Printf("[Worker %v]: Started\n", id)
	defer wg.Done()

	switch id {
	case 1:
		httpServer.Start()
	case 2:
		// TODO: start udp server
	case 3:
		// TODO: add rpc server
	}

	fmt.Printf("[Worker %v]: Finished\n", id)
}

func main() {
	fmt.Println("[Main]: Started")
	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		fmt.Println("[Main]: Starting worker", i)
		wg.Add(1)
		go start_server(&wg, i)
	}

	wg.Wait()
}
