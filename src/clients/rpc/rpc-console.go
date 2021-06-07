package rpcConsole

import (
	"strconv"
	"strings"
	"unsafe"

	//Aqui ps
	rpcClient "proyecto1.com/main/src/clients/rpc"

	"github.com/AlecAivazis/survey/v2"
	"proyecto1.com/main/src/count"
	"proyecto1.com/main/src/utils"
)

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
				rpcClient.Invoke(utils.OPERATIONS.INCREMENT, s)
			} else if err == nil && strings.Compare(result, "Salir") != 0 && strings.Compare(result, "Reducir Cuenta") == 0 && unsafe.Sizeof(s) <= 8 && s != 0 {
				count.SharedCounter.Decrement(s, "Local")
			} else {
				rpcClient.Invoke(utils.OPERATIONS.DECREMENT, s)
			}

		}
		s = 0
	}
}
