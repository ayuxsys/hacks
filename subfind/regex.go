package main

import (
	"regexp"
	"strings"
)

// createRegex creates a regex from a wildcard domain pattern.
func createRegex(wc string) string {
	pattern := regexp.QuoteMeta(wc)

	if !strings.Contains(wc, "*") {
		// example.com matches example.com and all its subdomains.
		return `^(?:.+\.)?` + pattern + `$`
	}

	// Wildcards must match at least one character.
	pattern = strings.ReplaceAll(pattern, `\*`, `.+`)

	return "^" + pattern + "$"
}
