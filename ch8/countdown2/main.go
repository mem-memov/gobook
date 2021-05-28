package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)

	fmt.Println("Commencing countdown")
	for i := 10; i > 0; i-- {
		select {
		case <-tick:
			fmt.Println(i)
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	fmt.Println("Launch!")
}
