package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	errL = log.New(os.Stderr, "[error] ", log.Ltime|log.Lshortfile)
	opL  = log.New(os.Stdout, "", 0)
)

func main() {
	concurrency := flag.Int("c", 5, "max concurrent requests to send")
	outputFile := flag.String("o", "", "path to the output jsonl file")
	flag.Parse()

	if *outputFile != "" {
		file, err := os.OpenFile(*outputFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		opL.SetOutput(io.MultiWriter(os.Stdout, file))
	}

	orgList := readOrg()

	var wg sync.WaitGroup
	wg.Add(len(orgList))
	sem := make(chan struct{}, *concurrency)
	for _, org := range orgList {
		sem <- struct{}{} // accquire slot
		go func(org string) {
			releaseSlot := func() { <-sem }
			defer releaseSlot()
			defer wg.Done()
			detectGithub(org)
		}(org)
	}
	wg.Wait()
}

func detectGithub(org string) {
	url := "https://github.com/" + org
	resp, err := http.Get(url)
	if err != nil {
		errL.Printf("'http.Get': %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		data := must(json.Marshal(NewOrg(org, Github)))
		opL.Println(string(data))
	}
}

func readOrg() []string {
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
