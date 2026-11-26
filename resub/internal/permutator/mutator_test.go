package permutator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMutateDomainHandleDuplicateParts(t *testing.T) {
	input := "api.api.staging-test.example.com"
	want := []string{
		"fuzz.example.com",
		"fuzz.api.staging-test.example.com",
		"api.fuzz.staging-test.example.com",
		"api.api.fuzz-test.example.com",
		"api.api.staging-fuzz.example.com",
	}
	got := mutateDomain(input, "fuzz")
	require.Equal(t, want, got)
}

func TestMutatePart(t *testing.T) {
	input := "example-uat-test"
	want := []string{"fuzz-uat-test", "example-fuzz-test", "example-uat-fuzz"}
	require.Equal(t, want, mutatePart(input, "fuzz"))
}

func TestMutateDomain(t *testing.T) {
	got := mutateDomain("staging-uat.www.example.com", "FUZZ")
	want := []string{
		"FUZZ.example.com",
		"FUZZ-uat.www.example.com",
		"staging-FUZZ.www.example.com",
		"staging-uat.FUZZ.example.com",
	}
	require.Equal(t, got, want)
}
