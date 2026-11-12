package scope

import "encoding/json"

type Scope struct {
	Value string `json:"scope"`
	Type  Type   `json:"type"`
}

func New(value string, stype Type) Scope {
	return Scope{Value: value, Type: stype}
}

type Type int

const (
	WildCard Type = iota + 1
	Domain
	URL
	IPv4
	CIDR
	// Mobile Applications
	IOS
	Android
	// Source Codes
	Git
	NPM
	PyPi      // rare
	GoPackage // rare

	Other
)

func (t Type) string() string {
	switch t {
	case WildCard:
		return "wildcard"
	case URL:
		return "url"
	case Domain:
		return "domain"
	case IPv4:
		return "ipv4"
	case CIDR:
		return "cidr"
	case Git:
		return "git"
	case IOS:
		return "ios"
	case NPM:
		return "npm"
	case PyPi:
		return "pypi"
	case GoPackage:
		return "gopkg"
	case Android:
		return "android"
	case Other:
		return "other"
	default:
		return "unknown"
	}
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.string())
}
