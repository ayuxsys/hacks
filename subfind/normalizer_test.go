package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegex(t *testing.T) {
	wc := "*example.com"
	re := regexp.MustCompile(createRegex(wc))

	tests := []struct {
		input string
		match bool
	}{
		{input: "foobar.staging.example.com", match: true},
		{input: "staging.example.com", match: true},
		{input: "staging.example.org", match: false},
		{input: "example.com", match: false},
	}

	for _, tt := range tests {
		require.Equal(t, tt.match, re.MatchString(tt.input))
	}
}

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
