// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

type Request struct {
	action string
	result chan int
}

func numberServer(increment, decrement, get chan struct{}, result chan int) {
	var i int

	for {
		select {
		case <-increment:
			i++
		case <-decrement:
			i--
		case <-get:
			result <- i
		}
	}
}

func incrementing(increment chan struct{}, done chan bool) {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000000; j++ {
		increment <- struct{}{}
	}
	done <- true
}

func decrementing(decrement chan struct{}, done chan bool) {
	//TODO: decrement i 1000000 times
	for j := 0; j < 1000000; j++ {
		decrement <- struct{}{}
	}
	done <- true
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	increment := make(chan struct{})
	decrement := make(chan struct{})
	get := make(chan struct{})
	result := make(chan int)
	done := make(chan bool)

	go numberServer(increment, decrement, get, result)
	// TODO: Spawn both functions as goroutines
	go incrementing(increment, done)
	go decrementing(decrement, done)

	<-done
	<-done
	get <- struct{}{}
	finalValue := <-result

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", finalValue)
}
