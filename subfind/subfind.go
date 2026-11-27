package main

import (
	"fmt"
	"strings"
)

type Result struct {
	WildCard   string   `json:"wildcard"`
	SubDomains []string `json:"subdomains"`
}

// findSubdomain runs all the subdomain enumerator tools on the given domain
func findSubdomain(wc string) error {
	wc = trimHttpPrefix(wc)
	infoL.Println("enumerating subdomains for", wc)

	domain := normalizeWC(wc)

	var subs []string

	subs, err := subfinder([]string{domain}, strings.Split(subfinderAddonArgs, " ")...)
	if err != nil {
		errL.Println(joinErr("subfinder", err).Error())
	}

	subs = dedupeSlice(subs)
	infoL.Printf("found %d subdomains for %q", len(subs), domain)

	if recursion && len(subs) != 0 {
		infoL.Printf("running recursive search to find subdomains of %d enumerated subdomains", len(subs))
		rSubs, err := subfinder(subs, append(strings.Split(recSubfinderAddonArgs, "\n"), "-recursive")...)
		if err != nil {
			errL.Println(joinErr("subfinder recursion run", err).Error())
		}
		subs = append(subs, dedupeSlice(rSubs)...)
	}
	okL.Printf("found total %d subdomains for domain %q", len(subs), domain)

	for _, s := range subs {
		fmt.Println(s)
	}

	// infoL.Println("filtering subdomains")
	// var filteredSubs []string
	// re := regexp.MustCompile(createRegex(wc))
	// for _, s := range subs {
	// 	if re.MatchString(s) {
	// 		filteredSubs = append(filteredSubs, s)
	// 	}
	// }

	if databaseConn != nil {
		subfinder := &SubFinder{
			Domain: wc,
			Subs:   subs,
			Tag:    tag,
		}
		if err := databaseConn.Create(subfinder).Error; err != nil {
			errL.Fatal(err)
		}
		okL.Printf("pushed %d sudomain for domain %q to database", len(subs), domain)
	}

	return nil
}

// subfinder runs the subfinder by projectdiscovery to discover subdomains of the provided domains
// it returns the slice of enumerated subdomains
func subfinder(domains []string, args ...string) ([]string, error) {
	shell := shell{
		name:  "subfinder",
		stdin: strings.Join(domains, "\n"),
	}

	stdout, err := shell.execute(args...)
	if err != nil {
		return nil, joinErr("execCmd", err)
	}
	return strings.Split(stdout, "\n"), nil
}
