package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown")
	tick := time.Tick(1 * time.Second)
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		<- tick
	}
	fmt.Println("Launch!")
}
