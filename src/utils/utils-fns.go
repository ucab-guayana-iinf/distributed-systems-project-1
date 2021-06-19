package utils

import (
	"strconv"
)

type OperationsStruct struct {
	GET       int
	INCREMENT int
	DECREMENT int
	RESTART   int
}

var OPERATIONS = OperationsStruct{1, 2, 3, 4}

const STOP = "STOP"
const INCREMENT = "Incrementar"
const DECREMENT = "Decrementar"
const RESTART = "Reiniciar"
const GET_COUNT = "Imprimir"

func IntToString(n int) string {
	return strconv.Itoa(n)
}

func StringToInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}
