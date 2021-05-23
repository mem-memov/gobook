package main

import "fmt"

func main() {
	nats := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 1; i < 101; i++ {
			nats <- i
		}
		close(nats)
	}()

	go func() {
		for {
			i, ok := <-nats
			if !ok {
				break
			}
			squares <- i * i
		}
		close(squares)
	}()

	for {
		sq, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(sq)
	}
}
