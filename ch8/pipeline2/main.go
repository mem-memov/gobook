package main

import "fmt"

func main() {
	nats := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			nats <- i
		}
		close(nats)
	}()

	go func() {
		for i := range nats {
			squares <- i * i
		}
		close(squares)
	}()

	for sq := range squares {
		fmt.Println(sq)
	}
}
