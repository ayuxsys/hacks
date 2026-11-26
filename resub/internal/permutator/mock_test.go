package permutator

import "strings"

var (
	mockPermutator = New(mockDomains, "FUZZ", mockWordlist)
	mockDomains    = []string{"example.FUZZ.com", "example.org", "test.io"}
	mockWordlist   = strings.NewReader("s3\ntest-uat\ncmd\nadmin\ndashboard")
)
