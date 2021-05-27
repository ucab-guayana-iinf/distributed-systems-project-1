// package httpServer

// import (
// 	"fmt"
// 	"net/http"

// 	"proyecto1.com/main/src/count"
// 	"proyecto1.com/main/src/utils"
// )

// func Start() {
// 	fmt.Println("[HTTP Server]: Starting")
// 	var port = "2020"

// 	http.HandleFunc("/get-count", handleGetCount)
// 	http.HandleFunc("/increment-count", handleIncrementCount)
// 	http.HandleFunc("/decrement-count", handleDecrementCount)
// 	http.HandleFunc("/reset-count", handleResetCount)

// 	fmt.Println("[HTTP Server]: Running in http://localhost:" + port)
// 	http.ListenAndServe(":" + port, nil)
// }
