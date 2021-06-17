package menu

import (
	"fmt"
	"strconv"
	"unsafe"

	rpcConsole "proyecto1.com/main/src/clients/rpc"
	utils "proyecto1.com/main/src/utils"

	"github.com/AlecAivazis/survey/v2"
)

const TCP_THREAD_CLI = "TCP Hilos"
const TCP_PROCESS_CLI = "TCP Procesos"
const UDP_CLI = "UDP"
const RPC_CLI = "RPC"
const EXIT = "Salir"
const BACK = "Volver"
const INCREMENT = "Incrementar cuenta"
const DECREMENT = "Decrementar cuenta"
const RESTART = "Reiniciar cuenta"
const GET_COUNT = "Imprimir cuenta"

var MAIN_MENU_OPTIONS = []string{TCP_THREAD_CLI, TCP_PROCESS_CLI, UDP_CLI, RPC_CLI, EXIT}
var OPERATIONS_OPTIONS = []string{INCREMENT, DECREMENT, RESTART, GET_COUNT, BACK}

func askForNumber() int {
	var v int
	var value string
	var output int
	var err error

	for value == "" {
		value_prompt := &survey.Input{
			Message: "Ingrese un valor",
		}
		survey.AskOne(value_prompt, &value, survey.WithIcons(func(icons *survey.IconSet) {
			icons.Question.Text = "➡️ "
		}))
		v, err = strconv.Atoi(value)
	
		if err == nil && unsafe.Sizeof(v) <= 8 && v != 0 {
			output = v
		} else {
			fmt.Println("Numero invalido ❌")
		}
	}

	return output
}

func printCount(count int) {
	fmt.Println("\nLa cuenta ahora es de:", count)
}

func Start() {
	var in_main_prompt bool = true
	var in_operation_prompt bool = false
	var result string
	var operation string

	for in_main_prompt == true {
		prompt := &survey.Select{
			Message: "¿A qué consola se desea conectar?",
			Options: MAIN_MENU_OPTIONS,
		}
		survey.AskOne(prompt, &result, survey.WithIcons(func(icons *survey.IconSet) {
			icons.Question.Text = "➡️ "
		}))
		
		if result == EXIT {
			in_main_prompt = false
		} else {
			in_operation_prompt = true
		}

		for in_operation_prompt == true {
			operation_prompt := &survey.Select{
				Message: "¿Qué operacion desea realizar?",
				Options: OPERATIONS_OPTIONS,
			}
			survey.AskOne(operation_prompt, &operation, survey.WithIcons(func(icons *survey.IconSet) {
				icons.Question.Text = "➡️ "
			}))

			switch operation {
				case INCREMENT:
					num := askForNumber()
					fmt.Printf("%v", num)
					// TODO: increment using the selected client
					switch result {
						case TCP_THREAD_CLI:
							// TODO: increment with tcp thread
						case TCP_PROCESS_CLI:
							// TODO: increment with tcp thread
						case UDP_CLI:
							// TODO: increment with udp
						case RPC_CLI:
							rpcConsole.InvokeRpcCall(utils.OPERATIONS.INCREMENT, num)
							printCount(rpcConsole.InvokeRpcCall(utils.OPERATIONS.GET, 1))
					}
				case DECREMENT:
					num := askForNumber()
					fmt.Printf("%v", num)
					switch result {
						case TCP_THREAD_CLI:
							// TODO: decrement with tcp thread
						case TCP_PROCESS_CLI:
							// TODO: decrement with tcp thread
						case UDP_CLI:
							// TODO: decrement with udp
						case RPC_CLI:
							rpcConsole.InvokeRpcCall(utils.OPERATIONS.DECREMENT, num)
					}
				case RESTART:
					switch result {
						case TCP_THREAD_CLI:
							// TODO: restart with tcp thread
						case TCP_PROCESS_CLI:
							// TODO: restart with tcp thread
						case UDP_CLI:
							// TODO: restart with udp
						case RPC_CLI:
							rpcConsole.InvokeRpcCall(utils.OPERATIONS.RESTART, 1)
					}
				case GET_COUNT:
					switch result {
						case TCP_THREAD_CLI:
							// TODO: get count with tcp thread
						case TCP_PROCESS_CLI:
							// TODO: get count with tcp thread
						case UDP_CLI:
							// TODO: get count with udp
						case RPC_CLI:
							printCount(rpcConsole.InvokeRpcCall(utils.OPERATIONS.GET, 1))
					}
				case BACK:
					in_operation_prompt = false
			}
		}

	}
}
