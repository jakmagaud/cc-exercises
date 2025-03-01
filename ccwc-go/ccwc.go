package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func byteCount(data []byte) int {
	return len(data)
}

func lineCount(data []byte) int {
	return bytes.Count(data, []byte("\n"))
}

func wordCount(data []byte) int {
	return len(bytes.Fields(data))
}

func multiByteCount(data []byte) int {
	return utf8.RuneCountInString(string(data))
}

func main() {
	counterFlag := flag.Bool("c", false, "count the number of characters")
	wordFlag := flag.Bool("w", false, "count the number of words")
	lineFlag := flag.Bool("l", false, "count the number of lines")
	multiByteFlag := flag.Bool("m", false, "count the number of multi-byte characters")

	flag.Parse()
	tail := flag.Args()

	var data []byte = nil
	var err error = nil
	filename := ""

	if len(tail) == 0 {
		// reading from stdin
		data, err = io.ReadAll((os.Stdin))
		check(err)
	} else {
		// reading from file
		filename = tail[0]
		data, err = os.ReadFile(filename)
		check(err)
	}

	if *counterFlag {
		result := byteCount(data)
		fmt.Printf("%d %s", result, filename)
	} else if *lineFlag {
		result := lineCount(data)
		fmt.Printf("%d %s", result, filename)
	} else if *wordFlag {
		result := wordCount(data)
		fmt.Printf("%d %s", result, filename)
	} else if *multiByteFlag {
		result := multiByteCount(data)
		fmt.Printf("%d %s", result, filename)
	} else {
		fmt.Printf("%d %d %d %s\n", lineCount(data), wordCount(data), byteCount(data), filename)
	}

}
