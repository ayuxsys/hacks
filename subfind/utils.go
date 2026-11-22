package main

import (
	"bufio"
	"os"
)

// dedupeSlice removes duplicate entires from a slice and returns the new unique slice
func dedupeSlice(slice []string) []string {
	var newSlice []string
	seen := make(map[string]struct{}, len(slice))

	for _, s := range slice {
		if _, found := seen[s]; found {
			continue
		}
		if s != "" {
			newSlice = append(newSlice, s)
		}
		seen[s] = struct{}{}
	}

	return newSlice
}

// readStdin reads stdin and returns the slice of string seperated by newline
func readStdin() []string {
	var s []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return s
}
