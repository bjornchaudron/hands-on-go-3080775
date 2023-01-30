// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Something bad happened")
		}
	}()

	// use os package to read the file specified as a command line argument
	filename := os.Args[1]
	bs, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("Failed to read file: %s", err))
	}

	// convert the bytes to a string
	data := string(bs)

	// initialize a map to store the stats
	stats := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(data)

	// capture the length of the words slice
	stats["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words {
		for _, r := range word {
			switch {
			case unicode.IsLetter(r):
				stats["letters"]++
			case unicode.IsNumber(r):
				stats["numbers"]++
			default:
				stats["symbols"]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(stats)
}
