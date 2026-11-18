package main

import "flag"

type cmd struct {
	concurreny int
	outputFile string
	code       codeDetectors
}

type codeDetectors struct {
	golang bool
}

func (c *cmd) ParseFlags() {
	flag.IntVar(&c.concurreny, "c", 5, "max number of concurrent requests to send")
	flag.StringVar(&c.outputFile, "o", "", "path to the jsonl outputfile")
	flag.BoolVar(&c.code.golang, "g", false, "find golang source codes in the provided repos or org")
	flag.Parse()
}
