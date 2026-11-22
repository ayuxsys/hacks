package main

import "strings"

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
