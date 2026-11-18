package scope

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDomainRegex(t *testing.T) {
	tests := []struct {
		domain string
		match  bool
	}{
		{domain: "example.com", match: true},
		{domain: "*.example.com", match: true},
		{domain: "example-*.com", match: true},
		{domain: "sub.*.example.com", match: true},
		{domain: "example.*", match: true},
		// {domain: "127.0.0.1", match: false}, // todo: fix leads to true using a workaround rn by checking ip addr first
	}

	for _, tt := range tests {
		require.Equal(t, tt.match, domainRegex.MatchString(tt.domain))
	}
}
