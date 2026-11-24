package main

import (
	"regexp"
	"strings"
)

// createRegex creates a regex from a wildcard string
func createRegex(wc string) string {
	pattern := regexp.QuoteMeta(wc)
	pattern = strings.ReplaceAll(pattern, `\*`, `.*`)
	return "^" + pattern + "$"
}
