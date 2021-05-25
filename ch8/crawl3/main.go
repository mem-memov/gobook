package main

import (
	"./links"
	"fmt"
	"log"
	"os"
)

func main() {
	workList := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		workList <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					workList <- foundLinks
				}()
			}
		}()
	}

	seenLinks := make(map[string]bool)

	for list := range workList {
		for _, url := range list {
			if !seenLinks[url] {
				seenLinks[url] = true
				fmt.Println(url)
				unseenLinks <- url
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<- tokens
	if err != nil {
		log.Print(err)
	}
	return list
}