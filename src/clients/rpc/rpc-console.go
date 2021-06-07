package main

import (
	"log"
	"net/rpc"
	"strconv"
	"strings"
	"unsafe"

	"github.com/AlecAivazis/survey/v2"
	"proyecto1.com/main/src/count"
	"proyecto1.com/main/src/utils"
)

func InvokeRpcCall(operation int, input int) int {
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


func main() {
	var result string
	var valor string
	var s int
	var err error

	for strings.Compare(result, "Salir") != 0 {
		prompt := &survey.Select{
			Message: "Que desea realizar:",
			Options: []string{"Aumentar Cuenta", "Reducir Cuenta", "Salir"},
		}
		survey.AskOne(prompt, &result, survey.WithIcons(func(icons *survey.IconSet) {
			icons.Question.Text = "🤡"
		}))

		for s == 0 && strings.Compare(result, "Salir") != 0 {

			prompt := &survey.Input{
				Message: "Ingrese un valor",
			}
			survey.AskOne(prompt, &valor, survey.WithIcons(func(icons *survey.IconSet) {
				icons.Question.Text = "🤡"
			}))
			s, err = strconv.Atoi(valor)

			if err == nil && strings.Compare(result, "Salir") != 0 && strings.Compare(result, "Aumentar Cuenta") == 0 && unsafe.Sizeof(s) <= 8 && s != 0 {
				InvokeRpcCall(utils.OPERATIONS.INCREMENT, s)
			} else if err == nil && strings.Compare(result, "Salir") != 0 && strings.Compare(result, "Reducir Cuenta") == 0 && unsafe.Sizeof(s) <= 8 && s != 0 {
				count.SharedCounter.Decrement(s, "Local")
			} else {
				InvokeRpcCall(utils.OPERATIONS.DECREMENT, s)
			}

		}
		s = 0
	}
}
