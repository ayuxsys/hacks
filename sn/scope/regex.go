package scope

import "regexp"

var (
	ipv4Regex   = regexp.MustCompile(`^[0-9.]{4}`)
	domainRegex = regexp.MustCompile(`^[\w+.]*`)
)

func IsDomain(sv string) bool {
	return domainRegex.MatchString(sv)
}

func IsIpv4(sv string) bool {
	return ipv4Regex.MatchString(sv)
}
