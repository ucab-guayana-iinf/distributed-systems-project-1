package rpcServer

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"proyecto1.com/main/src/count"
)

type Task int

var source = "RPC Server"

func (t *Task) IncrementCount(n int, reply *int) error {
	*reply = count.SharedCounter.Increment(n, source)
	return nil
}

func (t *Task) DecrementCount(n int, reply *int) error {
	*reply = count.SharedCounter.Decrement(n, source)
	return nil
}

func (t *Task) GetCount(n int, reply *int) error {
	*reply = count.SharedCounter.Get()
	return nil
}

func (t *Task) RestartCount(n int, reply *int) error {
	*reply = count.SharedCounter.Restart(source)
	return nil
}

func Start() {
	fmt.Println("[RPC Server]: Starting")

	task := new(Task)
	// Publish the receivers methods
	err := rpc.Register(task)
	if err != nil {
		fmt.Println("❌ Format of service Task isn't correct. ", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		fmt.Println("Listen error: ", e)
	}
	log.Printf("✅ Serving RPC server on port %d", 1234)
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		fmt.Println("❌ Error serving: ", err)
	}
}
