package permutator

import (
	"fmt"
	"io"
	"strings"

	"github.com/ayuxsec/resub/internal/utils"
)

type Permutator struct {
	domains    []string
	identifier string
	wordlist   io.Reader
}

func New(domains []string, identifier string, wordlist io.Reader) *Permutator {
	return &Permutator{
		domains:    domains,
		identifier: identifier,
		wordlist:   wordlist,
	}
}

func (p *Permutator) PermutateDomains() error {
	wdSlice := utils.FileToSlice(p.wordlist)

	p.normalizeDomains()

	for _, d := range p.domains {
		if err := permutateDomain(d, p.identifier, wdSlice); err != nil {
			return err
		}
		//fmt.Fprintf(os.Stderr, "\rprocessed (%d/%d) domains.", counter, len(p.domains))
	}
	return nil
}

func permutateDomain(domain string, identifier string, wordlistSlice []string) error {
	for _, w := range wordlistSlice {
		if strings.Contains(domain, identifier) {
			permutated := strings.ReplaceAll(domain, identifier, w)
			fmt.Println(permutated)
		}
	}
	return nil
}
