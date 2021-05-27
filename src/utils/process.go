package utils

import (
	"log"
	"os/exec"
	"strconv"
)

// TODO: use to validate count min/max
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

func StartProcess() {
	log.Printf("Haha business")

	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func IntToString(n int) string {
	return strconv.Itoa(n)
}