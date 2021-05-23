package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 1; ; i++ {
			naturals <- i
		}
	}()

	go func() {
		for {
			nat := <-naturals
			squares <- nat * nat
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
