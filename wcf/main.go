package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	wc string
)

func main() {
	flag.StringVar(&wc, "w", "", "input wildcard to filter")
	flag.Parse()
	subs := readStdin()

	re := regexp.MustCompile(createRegex(wc))
	for _, s := range subs {
		if re.MatchString(s) {
			fmt.Println(s)
		}
	}
}

// createRegex creates a regex from a wildcard string
func createRegex(wc string) string {
	pattern := regexp.QuoteMeta(wc)
	pattern = strings.ReplaceAll(pattern, `\*`, `.*`)
	return "^" + pattern + "$"
}

func readStdin() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var res []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			res = append(res, line)
		}
	}

	mustPanic(scanner.Err())
	return res
}

func mustPanic(err error) {
	if err != nil {
		panic(err)
	}
}
