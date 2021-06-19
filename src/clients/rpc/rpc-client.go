package rpcConsole

import (
	"log"
	"net/rpc"

	"proyecto1.com/main/src/utils"
)

func InvokeRpcCall(operation int, input int, ip string) int {
	var err error
	var address = ip + ":1234"

	client, err := rpc.DialHTTP("tcp", address)
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
