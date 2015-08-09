package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jtsylve/DFRWS2015/file/recycle"
)

func main() {
	// Ensure we have exactly one command line argument
	// in addition to our file name
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		return
	}

	// Open file for reading
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can not open file %s: %v\n", filename, err)
	}
	defer f.Close()

	// Parse file
	md, err := recycle.DecodeI(f)
	if err == recycle.ErrInvalid {
		log.Fatalf("%s is not a valid $Recycle.Bin $I file.\n", filename)
	}
	if err != nil {
		log.Fatalf("Can not parse file %s: %v\n", filename, err)
	}

	// Print results
	fmt.Printf(
		"File Name: %s\nFile Size: %d\nDeleted on %s\n",
		md.Name,
		md.Size,
		md.Deleted,
	)
}
