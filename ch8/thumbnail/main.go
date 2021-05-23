package main

import (
	"./thumbnail"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	var fileNames []string = []string{
		"/home/u/repos/gobook/ch8/thumbnail/flowers.jpg",
		"/home/u/repos/gobook/ch8/thumbnail/cat.jpg",
	}
	//makeThumbnails5(fileNames)

	names := make(chan string)

	go func() {
		for _, f := range fileNames {
			names <- f
		}
		close(names)
	}()

	size := makeThumbnails6(names)
	fmt.Println(size)
}

func makeThumbnails(fileNames []string) {
	for _, f := range fileNames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(fileNames []string) {
	for _, f := range fileNames {
		go thumbnail.ImageFile(f)
	}
}

func makeThumbnails3(fileNames []string) {
	ch := make(chan struct{})

	for _, f := range fileNames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	for range fileNames {
		<- ch
	}
}

func makeThumbnails5(fileNames []string) (thumbNames []string, err error) {
	type item struct {
		thumbName string
		err error
	}

	ch := make(chan item, len(fileNames))

	for _, f := range fileNames {
		go func(f string) {
			var it item
			it.thumbName, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range fileNames {
		it := <- ch
		if it.err != nil {
			return nil, it.err
		}
		thumbNames = append(thumbNames, it.thumbName)
	}

	return thumbNames, nil
}

func makeThumbnails6(fileNames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range fileNames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumbName, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			stat, _ := os.Stat(thumbName)
			sizes <- stat.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}
