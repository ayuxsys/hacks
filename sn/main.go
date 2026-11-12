package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"net/url"
	"os"
	"sn/scope"
	"strings"
)

func main() {
	opJsonLPath := flag.String("o", "scopes.jsonl", "path to the output jsonl file")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)

	file, err := os.OpenFile(*opJsonLPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		errL.Fatalf("open file %q: %v", *opJsonLPath, err)
	}
	defer file.Close()
	setOutput(file)

	for scanner.Scan() {
		sValue := scanner.Text()
		var sType scope.Type

		if strings.HasPrefix(sValue, "http") {
			parsed, err := url.Parse(sValue)
			if err != nil {
				warnL.Printf("parse scope value with prefix 'http' %q as url: %v", sValue, err)
				continue
			}
			sType = scope.DetectTypeWithHttpPrefix(parsed)
			if sType == scope.WildCard {
				sValue = parsed.Host
			}
		}

		opL.Println(
			string(must(json.Marshal(scope.New(sValue, sType)))),
		)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
