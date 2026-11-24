package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	recursion             bool
	subfinderAddonArgs    string
	recSubfinderAddonArgs string
)

func main() {
	outputFile := flag.String("o", "", "path to the jsonl file to write all enumerated subdomains into (after filtering scopes)")
	opAllSubsFile := flag.String("oa", "", "path to the jsonl file to write all enumerated subdomains into")
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
		defer setLogOp(*outputFile)()
	}

	if *opAllSubsFile != "" {
		defer setLogAllOP(*opAllSubsFile)()
	}

	for _, wc := range wildCards {
		if wc == "" {
			continue
		}
		if err := findSubdomain(wc); err != nil {
			errL.Printf("'findSubdomain' wc=%q: %v", wc, err)
			continue
		}
	}
}

func joinErr(context string, err error, v ...string) error {
	if v == nil {
		return fmt.Errorf("%q: %v", context, err)
	}
	return fmt.Errorf("%q: %s %v", context, strings.Join(v, " "), err)
}
