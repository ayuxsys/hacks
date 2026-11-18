package scope

import (
	"encoding/json"
	"net/netip"
	"sn/log"
	"sn/utils/must"
)

func detectIpType(v string) bool {
	if _, err := netip.ParsePrefix(v); err == nil {
		log.Output.Println(
			string(must.Eval(json.Marshal(New(v, CIDR)))),
		)
		return true
	}

	if addr, err := netip.ParseAddr(v); err == nil {
		if addr.Is4() {
			log.Output.Println(
				string(must.Eval(json.Marshal(New(v, IPv4)))),
			)
			return true
		}

		if addr.Is6() {
			log.Output.Println(
				string(must.Eval(json.Marshal(New(v, IPv6)))),
			)
			return true
		}
	}
	return false
}
