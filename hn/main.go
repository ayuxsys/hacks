package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var filterSuffixPattern = []string{
	"_bbp",
	"-bbp",
	"_vdp",
	"_h1c",
	"-bounty",
	"_bug_bounty",
	"_public",
	"-public",
	"-private",
	"_vip",
	"_sa",
	"_ltd",
	"_inc",
	"_bv",
	"_as",
	"_com",
	"_limited",
	"_jp",
	"_tw",
	"_uk",
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handle := scanner.Text()
		for _, suffix := range filterSuffixPattern {
			if new, found := strings.CutSuffix(handle, suffix); found {
				handle = new
				break
			}
		}
		fmt.Println(handle)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
