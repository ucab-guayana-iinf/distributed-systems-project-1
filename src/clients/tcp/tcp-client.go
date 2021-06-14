// package tcpClient
package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"unsafe"

	"github.com/AlecAivazis/survey/v2"
)

// To stop connection, send STOP message to server
func main() {
	var result string
	var valor string
	var s int
	var err error
	address := "localhost:2021"
	fmt.Println("[TCP Client]: Starting")
	// host, port, err := net.SplitHostPort("127.0.0.1:5432")
	// s := strings.Split("127.0.0.1:5432", ":")
	// ip, port := s[0], s[1]
	// fmt.Println(ip, port)

	// Connect to server
	c, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	// for {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print(">> ")
	// text, _ := reader.ReadString('\n')
	// Read user input
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
				fmt.Fprintf(c, "Increment "+valor+"\n")
			} else if err == nil && strings.Compare(result, "Salir") != 0 && strings.Compare(result, "Reducir Cuenta") == 0 && unsafe.Sizeof(s) <= 8 && s != 0 {
				fmt.Fprintf(c, "Decrement "+valor+"\n")
			} else if strings.Compare(result, "Salir") != 0 && strings.Compare(result, "Reiniciar Cuenta") == 0 {
				fmt.Fprintf(c, "Restart\n")
			} else {
				fmt.Println("Numero invalido 🤡")
			}
		}
		s = 0
	}
	// Salida
	fmt.Fprintf(c, "Stop\n")

	// Send input to the server
	// fmt.Fprintf(c, text+"\n")

	// Read messages from the server
	// message, _ := bufio.NewReader(c).ReadString('\n')
	// fmt.Print("Server: " + message)

	// Check for exit signal
	// if strings.ToUpper(strings.TrimSpace(string(text))) == "STOP" {
	// 	fmt.Println("TCP client exiting...")
	// 	return
	// }
	// }
}
