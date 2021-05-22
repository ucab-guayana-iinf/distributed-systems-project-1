package httpServer

import (
	"fmt"
	"net/http"

	"proyecto1.com/main/src/count"
	"proyecto1.com/main/src/utils"
)

func handleGetCount(w http.ResponseWriter, req *http.Request) {
	fmt.Println("[HTTP Server]: Processing /get-count")
	count.SharedCounter.Print()
	fmt.Fprintf(w, utils.IntToString(count.SharedCounter.Get()))
}

func handleResetCount(w http.ResponseWriter, req *http.Request) {
	fmt.Println("[HTTP Server]: Processing /reset-count")
}

func handleIncrementCount(w http.ResponseWriter, req *http.Request) {
	fmt.Println("[HTTP Server]: Processing /increment-count")
	// TODO  receive value to increment from body params
	fmt.Fprintf(w, utils.IntToString(count.SharedCounter.Increment(1)))
}

func handleDecrementCount(w http.ResponseWriter, req *http.Request) {
	fmt.Println("[HTTP Server]: Processing /decrement-count")
	// TODO  receive value to increment from body params
	fmt.Fprintf(w, utils.IntToString(count.SharedCounter.Decrement(1)))
}

func Start() {
	fmt.Println("[HTTP Server]: Starting")
	var port = "2020"

	http.HandleFunc("/get-count", handleGetCount)
	http.HandleFunc("/increment-count", handleIncrementCount)
	http.HandleFunc("/decrement-count", handleDecrementCount)
	http.HandleFunc("/reset-count", handleResetCount)

	fmt.Println("[HTTP Server]: Running in http://localhost:" + port)
	http.ListenAndServe(":" + port, nil)
}
