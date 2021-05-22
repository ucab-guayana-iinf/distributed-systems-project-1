package count

import (
	"log"

	"proyecto1.com/main/utils"
)

// Private

var count_value = 0
var blocked = false

func block() {
	blocked = true
}

func unblock() {
	blocked = false
}

func waitAvailability() {
	// do nothing if count blocked 
	// i.e: count is being accessed
	for (blocked == true) {} 
}

func get() int {
	return count_value
}

// Public

func Restart() int {
	defer unblock()
	waitAvailability()
	block()
	count_value = 0
	return get()
}

func Increment(n int) int {
	defer unblock()
	waitAvailability()
	block()
	count_value += n
	return get()
}

func Decrement(n int) int {
	defer unblock()
	waitAvailability()
	block()
	count_value -= n
	return get()
}

func Print() {
	log.Printf("La cuenta es %v", utils.IntToString(get()))
}