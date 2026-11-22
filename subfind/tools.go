package main

import "strings"

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
