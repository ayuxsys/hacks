package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func runOrgDetectors(org string) {
	for _, fp := range fingerPrints {
		url := strings.Replace(fp.Pattern, "{org}", org, 1)
		resp, err := http.Get(url)
		if err != nil {
			errL.Printf("'http.Get': %v", err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			data := must(json.Marshal(NewOrg(org, fp.Cloud)))
			opL.Println(string(data))
		}

	}

	// switch resp.StatusCode {
	// case 200:
	// 	if fp.Cloud == Github || fp.Cloud == Gitlab {
	// 		ct = fp.Cloud
	// 		data := must(json.Marshal(NewOrg(org, ct)))
	// 		opL.Println(string(data))
	// 	}
	// }
}
