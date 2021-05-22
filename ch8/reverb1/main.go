package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	for input.Scan() {
		go echo(c, input.Text(), 3*time.Second)
	}

	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
