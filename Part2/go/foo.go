// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

func number_server(add_number <-chan int, control <-chan int, number chan<- int) {
	var i = 0

	for ;; {
		select {
		case num := <- add_number:
			i += num
		case cmd := <- control:
			switch cmd {
			case GetNumber:
				// TODO: Send updated number
			case Exit:
				// TODO: End this goroutine gracefully
			}
		}
	}
}

func incrementing(add_number chan<-int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- 1
	}
	//TODO: signal that the goroutine is finished
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- -1
	}
	//TODO: signal that the goroutine is finished
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the required channels
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.
	add_number := make(chan int)
	number := make(chan int)
	control := make(chan int)

	// TODO: Spawn the required goroutines

	// TODO: block on finished from both "worker" goroutines

	control<-GetNumber
	Println("The magic number is:", <- number)
	control<-Exit
}
