package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	errL = log.New(os.Stderr, "[error] ", log.Ltime|log.Lshortfile)
	opL  = log.New(os.Stdout, "", 0)
)

func main() {
	c := &cmd{}
	c.ParseFlags()

	if c.outputFile != "" {
		file, err := os.OpenFile(c.outputFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		opL.SetOutput(io.MultiWriter(os.Stdout, file))
	}

	inputSlice := readStdin()

	var wg sync.WaitGroup
	wg.Add(len(inputSlice))
	sem := make(chan struct{}, c.concurreny)
	for _, inp := range inputSlice {
		sem <- struct{}{} // accquire slot
		go func(inp string) {
			releaseSlot := func() { <-sem }
			defer releaseSlot()
			defer wg.Done()
			runOrgDetectors(inp)
		}(inp)
	}
	wg.Wait()
}

func readStdin() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var orgList []string
	for scanner.Scan() {
		org := strings.TrimSpace(scanner.Text())
		orgList = append(orgList, org)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return orgList
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
