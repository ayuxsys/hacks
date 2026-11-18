package scope

import (
	"bufio"
	"embed"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"sn/log"
	"sn/utils/must"
	"strings"
)

// tld configuration
const ianaURL = "https://data.iana.org/TLD/tlds-alpha-by-domain.txt"

var (
	//go:embed tlds.txt
	localCacheTLDs embed.FS
	tlds           = make(map[string]struct{})
	isTLDsLoaded   bool
)

// regex
var (
	domainRegex   = regexp.MustCompile(`^[A-Za-z0-9-]+(\.[A-Za-z0-9-]+)+$`)
	wildcardRegex = regexp.MustCompile(`^.*\*.*$`) // use after detecting the domain
)

func detectDomainType(v string) {
	if wildcardRegex.MatchString(v) {
		log.Output.Println(string(must.Eval(json.Marshal(New(v, Domain)))))
		return
	}
	// shrug: this might be unreliable and map domain scopes as android mut meh couldn't think of anything else
	if isTLD(strings.Split(v, ".")[0]) {
		log.Output.Println(string(must.Eval(json.Marshal(New(v, Android)))))
		return
	}
	log.Output.Println(string(must.Eval(json.Marshal(New(v, Domain)))))
}

func isTLD(s string) bool {
	if !isTLDsLoaded {
		loadTLDs()
	}
	_, exists := tlds[s]
	return exists
}

func loadTLDs() {
	loadTLD := func(body io.Reader) {
		scanner := bufio.NewScanner(body)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "#") || line == "" {
				continue
			}
			tlds[strings.ToLower(line)] = struct{}{}
		}
		must.Panic(scanner.Err())
		isTLDsLoaded = true
	}

	resp, err := http.Get(ianaURL)
	if err == nil {
		defer resp.Body.Close()
		loadTLD(resp.Body)
		log.Info.Println("loaded tlds from iana website")
	} else {
		log.Err.Printf("failed to fetch latest tld list will use local cache. err='%v'", err)
		file := must.Eval(localCacheTLDs.Open("tlds.txt"))
		defer file.Close()
		loadTLD(file)
	}
}
