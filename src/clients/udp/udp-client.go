package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"unsafe"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	var result string
	var s int
	var value string
	address := "localhost:2002"

	fmt.Println("[UDP Client]: Starting")

	//	Returns an address of a endpoint UDP
	resolveAddr, err := net.ResolveUDPAddr("udp4", address)

	// Connect to server
	c, err := net.DialUDP("udp4", nil, resolveAddr)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Read user input
	for strings.Compare(result, "Salir") != 0 {
		prompt := &survey.Select{
			Message: "Que desea realizar:",
			Options: []string{"Aumentar Cuenta", "Reducir Cuenta", "Reiniciar Cuenta", "Salir"},
		}

		survey.AskOne(prompt, &result)

		for s == 0 && strings.Compare(result, "Salir") != 0 {
			prompt := &survey.Input{
				Message: "Ingrese un valor",
			}

			survey.AskOne(prompt, &value)

			s, err = strconv.Atoi(value)

			if err == nil && strings.Compare(result, "Salir") != 0 && unsafe.Sizeof(s) <= 8 && s != 0 {
				if strings.Compare(result, "Aumentar Cuenta") == 0 {
					fmt.Fprintf(c, "Increment "+value+"\n")
				}

				if strings.Compare(result, "Reducir Cuenta") == 0 {
					fmt.Fprintf(c, "Decrement "+value+"\n")
				}

				if strings.Compare(result, "Reiniciar Cuenta") == 0 {
					fmt.Fprintf(c, "Restart\n")
				}
			} else {
				fmt.Println("Numero invalido")
			}
		}

		s = 0
	}

	defer c.Close()
}
