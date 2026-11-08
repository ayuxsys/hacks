package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var filterSuffixPattern = []string{
	"_bbp",
	"_vdp",
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handle := scanner.Text()
		for _, suffix := range filterSuffixPattern {
			if new, found := strings.CutSuffix(handle, suffix); found {
				handle = new
				break
			}
		}
		fmt.Println(handle)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
