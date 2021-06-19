package menu

import (
	"fmt"
	"net"
	"strconv"
	"unsafe"

	rpcConsole "proyecto1.com/main/src/clients/rpc"
	tcpClient "proyecto1.com/main/src/clients/tcp"
	udpClient "proyecto1.com/main/src/clients/udp"
	utils "proyecto1.com/main/src/utils"

	"github.com/AlecAivazis/survey/v2"
	"github.com/adjust/rmq/v3"
)

const TCP_THREAD_CLI = "TCP Hilos"
const TCP_PROCESS_CLI = "TCP Procesos"
const UDP_CLI = "UDP"
const RPC_CLI = "RPC"
const EXIT = "Salir"
const BACK = "Volver"

var MAIN_MENU_OPTIONS = []string{TCP_THREAD_CLI, TCP_PROCESS_CLI, UDP_CLI, RPC_CLI, EXIT}
var OPERATIONS_OPTIONS = []string{utils.INCREMENT, utils.DECREMENT, utils.RESTART, utils.GET_COUNT, BACK}

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
	var udp_client net.Conn
	var tcp_client net.Conn
	var err error = nil
	var queue rmq.Queue
	var udpQueue rmq.Queue

	for in_main_prompt == true {
		prompt := &survey.Select{
			Message: "¿A qué consola se desea conectar?",
			Options: MAIN_MENU_OPTIONS,
		}
		survey.AskOne(prompt, &result, survey.WithIcons(func(icons *survey.IconSet) {
			icons.Question.Text = "➡️ "
		}))

		switch result {
		case EXIT:
			in_main_prompt = false
		case TCP_THREAD_CLI:
			tcp_client, err = tcpClient.InitTCPClientConnection()
			queue = tcpClient.ProcessTCPResponses(tcp_client.LocalAddr().String())
			in_operation_prompt = true
		case TCP_PROCESS_CLI:
			tcp_client, err = tcpClient.InitTCPProcessClientConnection()
			in_operation_prompt = true
		case UDP_CLI:
			udp_client, err = udpClient.InitUDPClientConnection()
			udpQueue = udpClient.ProcessUDPResponses()
			in_operation_prompt = true
		default:
			in_operation_prompt = true
		}

		if err != nil {
			fmt.Println(err)
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
			case utils.INCREMENT:
				num := askForNumber()
				switch result {
				case TCP_THREAD_CLI:
					tcpClient.InvokeTCPClientCall(tcp_client, utils.INCREMENT, num)
				case TCP_PROCESS_CLI:
					tcpClient.InvokeTCPClientCall(tcp_client, utils.INCREMENT, num)
				case UDP_CLI:
					udpClient.InvokeUDPClientCall(udp_client, utils.INCREMENT, num)
				case RPC_CLI:
					rpcConsole.InvokeRpcCall(utils.OPERATIONS.INCREMENT, num)
					printCount(rpcConsole.InvokeRpcCall(utils.OPERATIONS.GET, 1))
				}
			case utils.DECREMENT:
				num := askForNumber()
				switch result {
				case TCP_THREAD_CLI:
					tcpClient.InvokeTCPClientCall(tcp_client, utils.DECREMENT, num)
				case TCP_PROCESS_CLI:
					tcpClient.InvokeTCPClientCall(tcp_client, utils.DECREMENT, num)
				case UDP_CLI:
					udpClient.InvokeUDPClientCall(udp_client, utils.DECREMENT, num)
				case RPC_CLI:
					rpcConsole.InvokeRpcCall(utils.OPERATIONS.DECREMENT, num)
				}
			case utils.RESTART:
				switch result {
				case TCP_THREAD_CLI:
					tcpClient.InvokeTCPClientCall(tcp_client, utils.RESTART, 0)
				case TCP_PROCESS_CLI:
					tcpClient.InvokeTCPClientCall(tcp_client, utils.RESTART, 0)
				case UDP_CLI:
					udpClient.InvokeUDPClientCall(udp_client, utils.RESTART, 0)
				case RPC_CLI:
					rpcConsole.InvokeRpcCall(utils.OPERATIONS.RESTART, 1)
				}
			case utils.GET_COUNT:
				switch result {
				case TCP_THREAD_CLI:
					tcpClient.InvokeTCPClientCall(tcp_client, utils.GET_COUNT, 0)
				case TCP_PROCESS_CLI:
					tcpClient.InvokeTCPClientCall(tcp_client, utils.GET_COUNT, 0)
				case UDP_CLI:
					udpClient.InvokeUDPClientCall(udp_client, utils.GET_COUNT, 0)
				case RPC_CLI:
					printCount(rpcConsole.InvokeRpcCall(utils.OPERATIONS.GET, 1))
				}
			case BACK:
				switch result {
				case TCP_THREAD_CLI:
					queue.StopConsuming()
					tcpClient.InvokeTCPClientCall(tcp_client, utils.STOP, 0)
				case TCP_PROCESS_CLI:
					tcpClient.InvokeTCPClientCall(tcp_client, utils.STOP, 0)
				case UDP_CLI:
					udpQueue.StopConsuming()
					udpClient.InvokeUDPClientCall(udp_client, utils.STOP, 0)
				}
				in_operation_prompt = false
			}
		}

	}
}
