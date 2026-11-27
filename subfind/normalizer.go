package main

import (
	"strings"
)

func normalizeWC(wc string) string {
	if !strings.Contains(wc, "*") {
		return wc
	}
	if strings.Count(wc, "*") == 1 {
		return normalizeSingleAsterisk(wc)
	}
	return "" // skip other subdomains for now i guess
}

func normalizeSingleAsterisk(w string) string {
	// *.example.com => example.com
	if strings.HasPrefix(w, "*.") {
		return strings.Replace(w, "*.", "", 1)
	}
	// *example.com => example.com
	if strings.HasPrefix(w, "*") {
		return strings.Replace(w, "*", "", 1)
	}

	// staging-uat.*.example.com => example.com (do the filtering later)
	pos := strings.Index(w, "*")
	return strings.Join(strings.Split(w[pos+1:], ".")[1:], ".")
}

func cleanWC(s string) string {
	s, _ = strings.CutSuffix(s, "/")
	return trimHttpPrefix(s)
}

// trimHttpPrefix cuts 'http://' or 'https://' prefix from a string if present
// otherwise the same string is returned back
func trimHttpPrefix(s string) string {
	if st, found := strings.CutPrefix(s, "http://"); found {
		return st
	}
	if st, found := strings.CutPrefix(s, "https://"); found {
		return st
	}
	return s
}
