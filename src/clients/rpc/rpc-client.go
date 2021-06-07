package rpcClient

import (
	"log"
	"net/rpc"

	"proyecto1.com/main/src/utils"
)

func Invoke(operation int, input int) int {
	var err error

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	var reply int

	switch operation {
	case utils.OPERATIONS.GET:
		client.Call("Task.GetCount", 1, &reply)
	case utils.OPERATIONS.INCREMENT:
		client.Call("Task.IncrementCount", input, &reply)
	case utils.OPERATIONS.DECREMENT:
		client.Call("Task.DecrementCount", input, &reply)
	case utils.OPERATIONS.RESTART:
		client.Call("Task.RestartCount", 1, &reply)
	}

	return reply
}
