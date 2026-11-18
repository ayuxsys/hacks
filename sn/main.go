package main

import (
	"bufio"
	"flag"
	"os"
	"sn/log"
	"sn/scope"
)

func main() {
	opJsonLPath := flag.String("o", "", "path to the output jsonl file")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)

	if *opJsonLPath != "" {
		file, err := os.OpenFile(*opJsonLPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Err.Fatalf("open file %q: %v", *opJsonLPath, err)
		}
		defer file.Close()
		log.SetOutput(file)
	}

	for scanner.Scan() {
		sValue := scanner.Text()
		scope.Detect(sValue)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
