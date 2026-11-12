package main

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
	"golang.org/x/term"
)

// loggers
var (
	warnL = log.New(os.Stderr, "[warn] ", log.Lshortfile|log.Ltime)
	errL  = log.New(os.Stderr, "[error]", log.Lshortfile|log.Ltime)
	opL   = log.New(os.Stdout, "", 0) // output logger
)

func init() {
	if term.IsTerminal(int(os.Stderr.Fd())) {
		warnL.SetPrefix(color.YellowString("[warn] "))
		errL.SetPrefix(color.RedString("[error] "))
	}

}

func setOutput(out *os.File) {
	multiWriter := io.MultiWriter(os.Stdout, out)
	opL.SetOutput(multiWriter)
}

// helpers
func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func mustPanic(err error) {
	if err != nil {
		panic(err)
	}
}
