package main

import "testing"

func TestNormalizeWC(t *testing.T) {
	wcs := []string{
		"example.com",
		"*.exampl.com",
		"staging-*.example.com",
		"staging.*.example.com",
		"*.example-*.com",
	}

	for _, wc := range wcs {
		t.Log(normalizeWC(wc))
	}
}
