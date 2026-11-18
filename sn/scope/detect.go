package scope

import (
	"encoding/json"
	"sn/log"
	"sn/utils/must"
	"strings"
)

func Detect(value string) {
	// detect url scopes
	if strings.HasPrefix(value, "http") {
		detectTypeWithHttpPrefix(value)
		return
	}

	// shurg: detect ip addr before domain because for ipv4 getting maped as domain b/c of my retardness :)
	if detected := detectIpType(value); detected == true {
		return
	}

	// detect domain
	if detected := detectDomainType(value); detected == true {
		return
	}

	// shrug: map rest as other
	log.Output.Println(
		string(must.Eval(json.Marshal(New(value, Other)))),
	)
}
