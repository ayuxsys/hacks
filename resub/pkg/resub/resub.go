package resub

import (
	"os"

	"log"

	"github.com/ayuxsec/resub/internal/permutator"
	"github.com/ayuxsec/resub/internal/utils"
)

type Resub struct {
	Domains      []string
	Identifier   string
	WordListMode string
	WordListPath string
}

func New(domains []string, identifier string, wordlistMode string) *Resub {
	return &Resub{
		Domains:      domains,
		Identifier:   identifier,
		WordListMode: wordlistMode,
	}
}

func init() {
	_, err := os.Stat(BaseWordListPath)
	if os.IsNotExist(err) {
		log.Println("downloading wordlist")
		setUpWordlist()
	}
}

func (r *Resub) PermutateDomains() error {
	var wordlistPath string
	if r.WordListPath == "" {
		wordlistPath = BaseWordListPath + "/n0kovo_subdomains_" + r.WordListMode + ".txt"
	} else {
		wordlistPath = r.WordListPath
	}
	file, err := os.Open(wordlistPath)
	if err != nil {
		return utils.WrapError("os.Open", err)
	}
	defer file.Close()

	permutator := permutator.New(r.Domains, r.Identifier, file)
	return permutator.PermutateDomains()
}
