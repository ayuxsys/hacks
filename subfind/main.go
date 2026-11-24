package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	recursion             bool
	subfinderAddonArgs    string
	recSubfinderAddonArgs string
)

func main() {
	outputFile := flag.String("o", "", "path to the output file to write enumerated subdomains into (after filtering scopes)")
	opAllSubsFile := flag.String("oa", "", "path to the output file to write all enumerated subdomains into")
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

	if *opAllSubsFile != "" {
		defer setLogAllOP(*opAllSubsFile)()
	}

	for _, wc := range wildCards {
		if wc == "" {
			continue
		}
		domain := normalizeWC(wc)
		if err := findSubdomain(domain); err != nil {
			errL.Printf("'findSubdomain' domain=%q: %v", domain, err)
			continue
		}
	}
}

// findSubdomain runs all the subdomain enumerator tools on the given domain
func findSubdomain(wc string) error {
	wc = trimHttpPrefix(wc)
	infoL.Println("enumerating subdomains for", wc)

	re := regexp.MustCompile(createRegex(wc))
	domain := normalizeWC(wc)

	var subs []string

	subs, err := subfinder([]string{domain}, strings.Split(subfinderAddonArgs, " ")...)
	if err != nil {
		errL.Println(joinErr("subfinder", err).Error())
	}

	subs = dedupeSlice(subs)
	infoL.Printf("found %d subdomains for %q", len(subs), domain)

	if recursion && len(subs) != 0 {
		infoL.Println("running recursive search to find subdomains of enumerated subdomains")
		rSubs, err := subfinder(subs, append(strings.Split(recSubfinderAddonArgs, "\n"), "-recursive")...)
		if err != nil {
			errL.Println(joinErr("subfinder recursion run", err).Error())
		}
		subs = append(subs, dedupeSlice(rSubs)...)
	}

	for _, d := range subs {
		opAL.Println(d)
	}

	okL.Printf("found total %d subdomains for domain %q", len(subs), domain)
	infoL.Println("filtering subdomains")

	for _, s := range subs {
		if re.MatchString(s) {
			opL.Println(s)
		}
	}

	return nil
}

func joinErr(context string, err error, v ...string) error {
	if v == nil {
		return fmt.Errorf("%q: %v", context, err)
	}
	return fmt.Errorf("%q: %s %v", context, strings.Join(v, " "), err)
}
