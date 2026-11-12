package scope

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetectIpAddr(t *testing.T) {
	ipv4Tests := []struct {
		name  string
		input string
		want  bool
	}{
		// Valid IPv4 addresses
		{name: "valid public IP", input: "1.1.1.1", want: true},
		{name: "valid google DNS", input: "8.8.8.8", want: true},
		{name: "valid localhost", input: "127.0.0.1", want: true},
		{name: "valid private IP", input: "192.168.1.1", want: true},
		{name: "valid zero address", input: "0.0.0.0", want: true},
		{name: "valid maximum address", input: "255.255.255.255", want: true},
		{name: "valid minimum octets", input: "1.1.1.1", want: true},
		{name: "valid boundary octets", input: "254.254.254.254", want: true},

		// Invalid IPv4 addresses
		{name: "empty string", input: "", want: false},
		{name: "missing octet", input: "192.168.1", want: false},
		{name: "extra octet", input: "192.168.1.1.1", want: false},
		{name: "octet greater than 255", input: "256.1.1.1", want: false},
		{name: "negative octet", input: "-1.1.1.1", want: false},
		{name: "alphabetic characters", input: "abc.def.ghi.jkl", want: false},
		{name: "mixed characters", input: "192.168.a.1", want: false},
		{name: "special characters", input: "192.168.1.@", want: false},
		{name: "leading dot", input: ".192.168.1.1", want: false},
		{name: "trailing dot", input: "192.168.1.1.", want: false},
		{name: "consecutive dots", input: "192..168.1.1", want: false},
		{name: "spaces", input: "192.168.1. 1", want: false},
		{name: "leading whitespace", input: " 192.168.1.1", want: false},
		{name: "trailing whitespace", input: "192.168.1.1 ", want: false},
		{name: "empty octets", input: "...", want: false},
		{name: "octet overflow", input: "999.999.999.999", want: false},

		// IPv6 addresses (should not be IPv4)
		{name: "IPv6 localhost", input: "::1", want: false},
		{name: "IPv6 address", input: "2001:4860:4860::8888", want: false},
	}

	for _, tt := range ipv4Tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, IsIpv4(tt.input))
		})
	}
}
