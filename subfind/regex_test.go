package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckRegex(t *testing.T) {
	tests := []struct {
		Domain         string
		ShouldMatch    []string
		ShouldNotMatch []string
	}{
		{
			Domain: "*.example.com",
			ShouldMatch: []string{
				"test.example.com",
				"foobar.example.com",
				"baz.foobar.example.com",
				"api.example.com",
			},
			ShouldNotMatch: []string{
				"example.com",
				"test.example.org",
				"notexample.com",
			},
		},
		{
			Domain: "staging-*.example.com",
			ShouldMatch: []string{
				"staging-foobar.example.com",
				"staging-api.example.com",
			},
			ShouldNotMatch: []string{
				"test.example.com",
				"staging.example.com",
				"staging-foobar.example.org",
				"example.com",
			},
		},
		{
			Domain: "*.staging-*.example.com",
			ShouldMatch: []string{
				"staging-foobar.example.com",
				"staging-api.example.com",
				"test.staging-foobar.example.com",
				"foobar.staging-test.example.com",
			},
			ShouldNotMatch: []string{
				"test.example.com",
				"staging.example.com",
				"staging-foobar.example.org",
				"example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Domain, func(t *testing.T) {
			re, err := regexp.Compile(createRegex(tt.Domain))
			require.NoError(t, err, "invalid regex for domain %s", tt.Domain)

			for _, sub := range tt.ShouldMatch {
				require.True(t, re.MatchString(sub),
					"expected %q to match domain pattern %q", sub, tt.Domain)
			}

			for _, sub := range tt.ShouldNotMatch {
				require.False(t, re.MatchString(sub),
					"expected %q NOT to match domain pattern %q", sub, tt.Domain)
			}
		})
	}
}
