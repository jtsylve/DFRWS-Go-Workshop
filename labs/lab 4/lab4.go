package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/jtsylve/DFRWS2015/file/recycle"
)

// process processes a single $I file
func process(r io.Reader, filename string) {
	// Parse file
	md, err := recycle.DecodeI(r)
	if err == recycle.ErrInvalid {
		log.Printf("%s is not a valid $Recycle.Bin $I file.\n", filename)
		return
	}
	if err != nil {
		log.Printf("Can not parse file %s: %v\n", filename, err)
		return
	}

	// Print results
	fmt.Printf(
		"File Name: %s\nFile Size: %d\nDeleted on %s\n",
		md.Name,
		md.Size,
		md.Deleted,
	)
}

// processZip batch processes all files contained in a zip file
func processZip(filename string) {
	zr, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatalf("Can not open zip file %s: %v\n", filename, err)
	}
	defer zr.Close()

	// Create a channel to buffer the files.
	fileChan := make(chan *zip.File, runtime.NumCPU())

	var wg sync.WaitGroup

	// Let's create one go routine per processor
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for f := range fileChan {
				r, err := f.Open()
				if err != nil {
					log.Printf("Can not open file %s: %v\n", f.Name, err)
				}
				defer r.Close()

				process(r, f.Name)
			}

		}()
	}

	for _, f := range zr.File {
		// Add file to channel
		fileChan <- f
	}

	close(fileChan)

	wg.Wait()
}

func main() {
	// Since we're using multi-processing we need to make sure that go
	// will use more than one processor.  Starting in Go 1.5, this is
	// set to runtime.NumCPU() by default.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Ensure we have exactly one command line argument
	// in addition to our file name
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		return
	}

	filename := os.Args[1]

	// If we're a zip file, batch process
	if strings.HasSuffix(filename, ".zip") {
		processZip(filename)
		return
	}

	// Otherwise, we can just process the single file

	// Open file for reading
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can not open file %s: %v\n", filename, err)
	}
	defer f.Close()

	process(f, filename)
}
