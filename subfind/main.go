package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	databaseURL           string
	recursion             bool
	subfinderAddonArgs    string
	recSubfinderAddonArgs string
	envPath               string
	tag                   string
	jsonOutputFile        string
)

func main() {
	flag.BoolVar(&recursion, "r", false, "run recursion on enumerated subdomains to find subdomains of subdomains (slow)")
	flag.StringVar(&stderrLogFile, "sl", "", "path to the log file to write stderr of executed commands into")
	flag.StringVar(&databaseURL, "u", "", "psql database url to use for connection (reads DB_URL from .env otherwise)")
	flag.StringVar(&envPath, "e", filepath.Join(must(os.UserConfigDir()), "subfind", ".env"), "path to the .env file to load configurations from")
	flag.StringVar(&tag, "t", "", "tag the data saved  (e.g. vdp/bbp)")
	flag.StringVar(&jsonOutputFile, "o", "", "fetch database for bbp and vdp subdomains then save output and exit")

	domain := flag.String("d", "", "domain to find subdomains for")

	flag.Parse()

	var wildCards []string

	toolArgs := flag.Args()
	fs := flag.NewFlagSet("tools", flag.ExitOnError)
	fs.StringVar(&subfinderAddonArgs, "sf", "", "additional args to be passed to subfinder (-t 5 -o ...)")
	fs.StringVar(&recSubfinderAddonArgs, "rsf", "", "additional args to be passed to subfinder when running on recursion (-all ...)")

	fs.Parse(toolArgs)

	initDBConnection()
	if jsonOutputFile != "" {
		if err := databaseConn.JsonOutput(jsonOutputFile); err != nil {
			errL.Fatal(err)
		}
		return
	}

	if *domain != "" {
		wildCards = append(wildCards, *domain)
	} else {
		wildCards = readStdin()
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
