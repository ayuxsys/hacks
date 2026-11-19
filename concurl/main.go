package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var (
	curlBinary = "curl-impersonate"
)

func main() {
	concurrency := flag.Int("c", 5, "max number of concurrent requests to send")
	outputFolder := flag.String("o", "out", "path to the output folder to save responses into")
	flag.Parse()

	// Everything after -- is passed to curl-impersonate
	impArgs = append(impArgs, flag.Args()...)

	urls := readStdin()
	log.SetFlags(0)

	if err := ensureDirExists(*outputFolder); err != nil {
		log.Fatalf("failed to create output folder: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	sem := make(chan struct{}, *concurrency)

	for _, u := range urls {
		sem <- struct{}{}
		go func(u string) {
			releaseSlot := func() { <-sem }
			defer releaseSlot()
			defer wg.Done()

			outFile := filepath.Join(*outputFolder, sha256sum(u))
			log.Println(outFile, u)

			if err := get(u, outFile); err != nil {
				log.Printf("get %s: %v", u, err)
			}
		}(u)
	}
	wg.Wait()
}

func get(url string, outputFile string) error {
	var args []string
	args = append(args, url, "-o", outputFile, "-s")
	args = append(args, impArgs...)

	cmd := exec.Command(curlBinary, args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("stderr: %q, err: %v", stderr.String(), err)
	}

	return nil
}

func sha256sum(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func readStdin() []string {
	var out []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		out = append(out, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return out
}

func ensureDirExists(path string) error {
	return os.MkdirAll(path, 0755)
}
