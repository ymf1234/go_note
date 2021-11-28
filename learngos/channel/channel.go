package main

import (
	"fmt"
	"time"
)

func creatWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c \n", id, <-c)
		}
	}()

	return c
}

func chanDemo() {
	//var c chan int // c == nil

	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = creatWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
