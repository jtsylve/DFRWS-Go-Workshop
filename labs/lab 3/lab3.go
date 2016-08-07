package main

import (
	"fmt"
	"os"
)

func main() {
	// Ensure we have exactly one command line argument
	// in addition to our file name
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		return
	}

	// Get filename argument
	filename := os.Args[1]

	// TODO: Open file

	// TODO: Parse file

	// TODO: Print results
}
