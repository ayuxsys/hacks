package scope

import (
	"encoding/json"
	"regexp"
	"slices"
	"sn/log"
	"sn/utils/must"
	"strings"
)

// shurg: tried tlds from iana was mapping tons of domain as
var (
	tlds = []string{
		"com",
		"in",
		"org",
	}
)

// regex
var (
	domainRegex   = regexp.MustCompile(`^[A-Za-z0-9*-]+(\.[A-Za-z0-9*-]+)+$`)
	wildcardRegex = regexp.MustCompile(`^.*\*.*$`) // use after detecting the domain
)

func detectDomainType(v string) bool {
	if !domainRegex.MatchString(v) {
		return false
	}

	if strings.Contains(v, "*") {
		log.Output.Println(string(must.Eval(json.Marshal(New(v, WildCard)))))
		return true
	}

	dParts := strings.Split(v, ".")
	tld := dParts[len(dParts)-1]

	switch tld {
	case "apk":
		log.Output.Println(string(must.Eval(json.Marshal(New(v, Android)))))
		return true
	case "ipa":
		log.Output.Println(string(must.Eval(json.Marshal(New(v, IOS)))))
		return true
	}

	// shrug: this might be unreliable and map domain scopes as android
	// android app package names are usually reversed form of domain name so the tld becomes the leftmost DNS label
	if isTLD(dParts[0]) {
		log.Output.Println(string(must.Eval(json.Marshal(New(v, Android)))))
		return true
	}

	log.Output.Println(string(must.Eval(json.Marshal(New(v, Domain)))))
	return true
}

func isTLD(s string) bool {
	return slices.Contains(tlds, s)
}
