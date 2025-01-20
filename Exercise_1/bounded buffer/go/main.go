package main

import (
	"fmt"
)
import "time"

func producer(s chan int) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		// TODO: push real value to buffer
		s <- i
	}

}

func consumer(s chan int) {

	time.Sleep(1 * time.Second)
	for {
		i := <-s //TODO: get real value from buffer
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {

	// TODO: make a bounded buffer
	s := make(chan int, 5)

	go consumer(s)
	go producer(s)

	select {}
}
