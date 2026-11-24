package main

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	infoL = log.New(os.Stderr, color.CyanString("[info] "), 0)
	warnL = log.New(os.Stderr, color.YellowString("[warn] "), 0)
	errL  = log.New(os.Stderr, color.RedString("[error]"), 0)
	okL   = log.New(os.Stderr, color.GreenString("[ok] "), 0)
	opL   = log.New(os.Stdout, "", 0)
	opAL  = log.New(io.Discard, "", 0)
)

func setLogOp(file *os.File) {
	opL.SetOutput(io.MultiWriter(file, os.Stdout))
}

func setLogAllOP(path string) func() {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	opAL.SetOutput(file)
	return func() { defer file.Close() }
}
