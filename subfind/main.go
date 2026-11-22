package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	recursion             bool
	subfinderAddonArgs    string
	recSubfinderAddonArgs string
)

func main() {
	outputFile := flag.String("o", "", "path to the output file")
	recursion = *flag.Bool("r", false, "run recursion on enumerated subdomains to find subdomains of subdomains (slow)")
	domain := flag.String("d", "", "domain to find subdomains for")
	flag.StringVar(&stderrLogFile, "sl", "", "path to the log file to write stderr of executed commands into")

	flag.Parse()

	var wildCards []string

	toolArgs := flag.Args()
	fs := flag.NewFlagSet("tools", flag.ExitOnError)
	fs.StringVar(&subfinderAddonArgs, "sf", "", "additional args to be passed to subfinder (-t 5 -o ...)")
	fs.StringVar(&recSubfinderAddonArgs, "rsf", "", "additional args to be passed to subfinder when running on recursion (-all ...)")

	fs.Parse(toolArgs)

	if *domain != "" {
		wildCards = append(wildCards, *domain)
	} else {
		wildCards = readStdin()
	}

	if *outputFile != "" {
		file, err := os.OpenFile(*outputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(joinErr("os.OpenFile", err))
		}
		defer file.Close()
		setLogOp(file)
	}

	for _, wc := range wildCards {
		if wc == "" {
			continue
		}
		domain := normalizeWC(wc)
		subs, err := findSubdomain(domain)
		if err != nil {
			errL.Printf("'findSubdomain' domain=%q: %v", domain, err)
			continue
		}
		uniqSubs := dedupeSlice(subs)
		for _, d := range uniqSubs {
			opL.Println(d)
		}
		okL.Printf("found %d subdomains for domain %q", len(uniqSubs), domain)
	}
}

// findSubdomain runs all the subdomain enumerator tools on the given domain
// and returns the enumerated subdomains
func findSubdomain(domain string) ([]string, error) {
	infoL.Println("enumerating subdomains for", domain)
	var subs []string
	subs, err := subfinder([]string{domain}, strings.Split(subfinderAddonArgs, " ")...)
	if err != nil {
		errL.Println(joinErr("subfinder", err).Error())
	}

	uniqSubs := dedupeSlice(subs)
	infoL.Printf("found %d subdomains for %q", len(uniqSubs), domain)

	if recursion {
		infoL.Println("running recursive search to find subdomains of enumerated subdomains")
		rSubs, err := subfinder(uniqSubs, append(strings.Split(recSubfinderAddonArgs, "\n"), "-recursive")...)
		if err != nil {
			errL.Println(joinErr("subfinder", err).Error())
		}
		uniqSubs = append(uniqSubs, dedupeSlice(rSubs)...)
	}

	return uniqSubs, nil
}

func joinErr(context string, err error, v ...string) error {
	if v == nil {
		return fmt.Errorf("%q: %v", context, err)
	}
	return fmt.Errorf("%q: %s %v", context, strings.Join(v, " "), err)
}
