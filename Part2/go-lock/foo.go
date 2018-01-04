// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
	"sync"
)

var i = 0
// TODO: create a global Mutex

func incrementing() {
	for j := 0; j<1000000; j++ {
		//TODO: sync access to variable i
		i++;
	}
}

func decrementing() {
	for j := 0; j<1000000; j++ {
		//TODO: sync access to variable i
		i--;
	}
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
	                                    // Try doing the exercise both with and without it!

    // TODO: Spawn both functions as goroutines
	
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(100*time.Millisecond)
    Println("The magic number is:", i)
}
