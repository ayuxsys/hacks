// responsible for creating fuzz identifiers in a domain name if not provided by the user
package permutator

import (
	"slices"
	"strings"
)

// adds an indeitifier if not present otherwise copies and returns the domains
func (p *Permutator) normalizeDomains() {
	var normalizedDomains []string
	for i := range p.domains {
		if !strings.Contains(p.domains[i], p.identifier) {
			normalizedDomains = append(normalizedDomains, mutateDomain(p.domains[i], p.identifier)...)
		} else {
			normalizedDomains = append(normalizedDomains, p.domains[i])
		}
	}
	p.domains = normalizedDomains
}

func mutateDomain(domain, identifier string) []string {
	var mutatedDomains []string

	// FUZZ.root.tld
	mutatedDomains = append(mutatedDomains, joinParts(identifier, getRootDomainWithTLD(domain)))

	parts := strings.Split(domain, ".")

	// loop over each part expect root.tld
	for i := range len(parts) - 2 {
		// mutated each part
		for _, mutated := range mutatePart(parts[i], identifier) {
			tmp := slices.Clone(parts)
			tmp[i] = mutated                        // change the specific part to the mutated one
			mutatedDomain := strings.Join(tmp, ".") // reconstruct the domain

			mutatedDomains = append(mutatedDomains, mutatedDomain)
		}
	}

	return mutatedDomains
}

func joinParts(parts ...string) string {
	return strings.Join(parts, ".")
}

// mutated domain parts
// example: input => uat-bookstore
func mutatePart(part, id string) []string {
	var o []string

	if strings.Contains(part, "-") {
		parts := strings.Split(part, "-")
		for i := range parts {
			tmp := slices.Clone(parts)
			tmp[i] = id
			o = append(o, strings.Join(tmp, "-"))
		}
	} else {
		return []string{id}
	}
	return o
}

func getRootDomainWithTLD(domain string) string {
	parts := strings.Split(domain, ".")
	return strings.Join(parts[len(parts)-2:], ".")
}
