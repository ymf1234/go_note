package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(str string, done chan struct{}) chan string {
	c := make(chan string)

	go func() {
		i := 0
		for  {
			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c<-fmt.Sprintf("%s : message %d", str, i)
			case <-done:
				fmt.Println("break")
				return
			}


			i++
		}

	}()

	return c
}

func famIn(c1, c2 chan string) chan string {
	c :=make(chan string)
	go func() {
		for  {
			c <- <-c1
		}
	}()

	go func() {
		for  {
			c <- <-c2
		}
	}()

	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c :=make(chan string)

	go func() {
		for  {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false

	}
}

func main() {
	done := make(chan struct{})
	gen1 := msgGen("service1", done)
	gen2 := msgGen("service2", done)

	for  {
		fmt.Println(<-gen1)
		if m, ok := timeoutWait(gen2, 2* time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout")
		}
	}

	done <- struct{}{}

	time.Sleep(time.Second)
}
