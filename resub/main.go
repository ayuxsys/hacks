package main

import (
	"log"

	"github.com/ayuxsec/resub/cmd"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
