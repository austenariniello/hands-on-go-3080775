// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if panicErr := recover(); panicErr != nil {
			fmt.Println("recovered from panic:", panicErr)
		}
	}()

	// use os package to read the file specified as a command line argument
	fileData, err := os.ReadFile(os.Args[1])
	checkErr(err)

	// convert the bytes to a string
	fileStr := string(fileData)

	// initialize a map to store the counts
	fileCount := map[string]int{
		"letters": 0,
		"symbols": 0,
		"numbers": 0,
		"words":   0,
	}

	// use the standard library utility package that can help us split the string into words
	fileLst := strings.Fields(fileStr)

	// capture the length of the words slice
	fileCount["words"] = len(fileLst)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range fileLst {
		for _, char := range word {
			switch {
			case unicode.IsLetter(char):
				fileCount["letters"]++
			case unicode.IsNumber(char):
				fileCount["numbers"]++
			case unicode.IsSymbol(char), unicode.IsPunct(char):
				fileCount["symbols"]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(fileCount)
}
