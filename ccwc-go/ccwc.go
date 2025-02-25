package main

import (
	"flag"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func byteCount(filename string) int64 {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	check(err)
	size := stat.Size()
	return size

}

func main() {
	counterFlag := flag.Bool("c", false, "count the number of characters")
	wordFlag := flag.Bool("w", false, "count the number of words")
	lineFlag := flag.Bool("l", false, "count the number of lines")
	flag.Parse()
	if !*counterFlag && !*wordFlag && !*lineFlag {
		fmt.Println("Please specify at least one flag")
		return
	}
	tail := flag.Args()
	if len(tail) == 0 {
		fmt.Println("Please specify a file")
		return
	}
	filename := tail[0]

	if *counterFlag {
		result := byteCount(filename)
		fmt.Printf("%d %s", result, filename)
	}

}
