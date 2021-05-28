package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	roots := flag.Args()

	fileSizes := make(chan int64)

	go func () {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nFiles, nBytes int64

	for size := range fileSizes {
		nFiles++
		nBytes += size
	}

	printDiskUsage(nFiles, nBytes)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo{
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error %v", err)
		return nil
	}
	return entries
}

func printDiskUsage(nFiles, nBytes int64) {
	fmt.Printf("\n%d files %.1f GB\n", nFiles, float64(nBytes)/1e9)
}
