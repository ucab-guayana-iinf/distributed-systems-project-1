package main

import (
	"proyecto1.com/main/count"
)

func main() {
	count.Print()
	
	count.Increment(1)
	count.Increment(1)
	count.Increment(1)
	
	count.Print()

	count.Decrement(2)
	
	count.Print()
}
