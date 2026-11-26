package permutator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRootD(t *testing.T) {
	input := "somesubdomain.test.example.com"
	want := "example.com"
	require.Equal(t, getRootDomainWithTLD(input), want)
}

func TestNormalizeDomains(t *testing.T) {
	input := []string{"example.com", "bar.foo.example.com", "bookstore.example.com", "foo-test.example.com"}
	mockPermutator.domains = input
	mockPermutator.normalizeDomains()
	t.Log(mockPermutator.domains)
}

// func TestPermDomain(t *testing.T) {
// 	subs, err := permutateDomain("test-FUZZ.example.com", "FUZZ", []string{"test", "uat", "admin", "secret"})
// 	require.NoError(t, err)
// 	for _, s := range subs {
// 		t.Log(s)
// 	}
// }
