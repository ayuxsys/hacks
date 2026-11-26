package resub

import (
	"log"
	"os"

	"github.com/ayuxsec/resub/internal/utils/shell"
)

var (
	BaseWordListPath    = os.Getenv("HOME") + "/.config/resub/n0kovo_subdomains"
	WordlistUpStreamURL = "https://github.com/n0kovo/n0kovo_subdomains"
)

func setUpWordlist() {
	if err := shell.Exec("git", "clone", WordlistUpStreamURL, BaseWordListPath); err != nil {
		log.Fatal(err)
	}
}
