package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	concurrency := flag.Int("c", 5, "max concurrent requests to send")
	flag.Parse()

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
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println(url)
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
